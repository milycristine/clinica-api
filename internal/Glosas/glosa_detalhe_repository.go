package glosa

import (
    "clinica-api/database"
    "clinica-api/internal/models"
    "database/sql"
    "fmt"
    "strings"
)

type GlosaDetalheRepository interface {
    CriarGlosa(g *models.GlosaDetalhe) error
    EditarGlosa(g *models.GlosaDetalhe) error
    ListarGlosas() ([]models.GlosaDetalhe, error)
    BuscarPorId(id int) (*models.GlosaDetalhe, error)
    ExisteGlosa(g *models.GlosaDetalhe) (bool, error)
}

type glosaDetalheRepository struct {
    db *sql.DB
}

func NovoGlosaDetalheRepository(conn *database.SQLStr) GlosaDetalheRepository {
    return &glosaDetalheRepository{
        db: conn.DB(),
    }
}

func (r *glosaDetalheRepository) CriarGlosa(g *models.GlosaDetalhe) error {
    tx, err := r.db.Begin()
    if err != nil {
        return err
    }

    query := `
        INSERT INTO GlosasDetalhes 
        (GlosasMensalId, UnidadeId, Guia, DataOcorrencia, PacienteId, Mo, NomeContratado, 
         DenteRegiao, MotivoGlosa, StatusRecurso)
        OUTPUT INSERTED.GlosaDetalheId
        VALUES (@MensalId, @UnidadeId, @Guia, @DataOcorrencia, @PacienteId, @Mo, @NomeContratado,
                @DenteRegiao, @MotivoGlosa, @StatusRecurso)
    `

    var detalheId int
    err = tx.QueryRow(
        query,
        sql.Named("MensalId", g.GlosasMensalId),
        sql.Named("UnidadeId", g.UnidadeId),
        sql.Named("Guia", g.Guia),
        sql.Named("DataOcorrencia", g.DataOcorrencia),
        sql.Named("PacienteId", g.PacienteId),
        sql.Named("Mo", g.Mo),
        sql.Named("NomeContratado", g.NomeContratado),
        sql.Named("DenteRegiao", g.DenteRegiao),
        sql.Named("MotivoGlosa", g.MotivoGlosa),
        sql.Named("StatusRecurso", g.StatusRecurso),
    ).Scan(&detalheId)

    if err != nil {
        tx.Rollback()
        return fmt.Errorf("erro ao criar GlosasDetalhes: %w", err)
    }

    for _, p := range g.Procedimentos {
        _, err = tx.Exec(`
            INSERT INTO GlosaProcedimentos
            (GlosaDetalheId, ProcedimentoId, CodigoProcedimento, NomeProcedimento,
             Quantidade, ValorInformado, ValorGlosado)
            VALUES (@GID, @PID, @Cod, @Nome, @Qtd, @VInf, @VGlos)
        `,
            sql.Named("GID", detalheId),
            sql.Named("PID", p.ProcedimentoId),
            sql.Named("Cod", p.CodigoProcedimento),
            sql.Named("Nome", p.NomeProcedimento),
            sql.Named("Qtd", p.Quantidade),
            sql.Named("VInf", p.ValorInformado),
            sql.Named("VGlos", p.ValorGlosado),
        )

        if err != nil {
            tx.Rollback()
            return fmt.Errorf("erro ao inserir procedimentos: %w", err)
        }
    }

    return tx.Commit()
}

func (r *glosaDetalheRepository) EditarGlosa(g *models.GlosaDetalhe) error {
    tx, err := r.db.Begin()
    if err != nil {
        return err
    }

    var setClauses []string
    var args []any

    campos := map[string]any{
        "Guia":           g.Guia,
        "DataOcorrencia": g.DataOcorrencia,
        "PacienteId":     g.PacienteId,
        "Mo":             g.Mo,
        "NomeContratado": g.NomeContratado,
        "DenteRegiao":    g.DenteRegiao,
        "MotivoGlosa":    g.MotivoGlosa,
        "StatusRecurso":  g.StatusRecurso,
    }

    for campo, valor := range campos {
        switch v := valor.(type) {
        case string:
            if v != "" {
                setClauses = append(setClauses, fmt.Sprintf("%s = @%s", campo, campo))
                args = append(args, sql.Named(campo, v))
            }
        case int:
            if v != 0 {
                setClauses = append(setClauses, fmt.Sprintf("%s = @%s", campo, campo))
                args = append(args, sql.Named(campo, v))
            }
        }
    }

    if len(setClauses) == 0 {
        tx.Rollback()
        return fmt.Errorf("nenhum campo informado para atualização")
    }

    args = append(args, sql.Named("Id", g.GlosaDetalheId))

    query := fmt.Sprintf(
        "UPDATE GlosasDetalhes SET %s WHERE GlosaDetalheId = @Id",
        strings.Join(setClauses, ", "),
    )

    _, err = tx.Exec(query, args...)
    if err != nil {
        tx.Rollback()
        return err
    }

    _, err = tx.Exec("DELETE FROM GlosaProcedimentos WHERE GlosaDetalheId = @Id",
        sql.Named("Id", g.GlosaDetalheId),
    )
    if err != nil {
        tx.Rollback()
        return err
    }

    for _, p := range g.Procedimentos {
        _, err = tx.Exec(`
            INSERT INTO GlosaProcedimentos
            (GlosaDetalheId, ProcedimentoId, CodigoProcedimento, NomeProcedimento,
             Quantidade, ValorInformado, ValorGlosado)
            VALUES (@GID, @PID, @Cod, @Nome, @Qtd, @VInf, @VGlos)
        `,
            sql.Named("GID", g.GlosaDetalheId),
            sql.Named("PID", p.ProcedimentoId),
            sql.Named("Cod", p.CodigoProcedimento),
            sql.Named("Nome", p.NomeProcedimento),
            sql.Named("Qtd", p.Quantidade),
            sql.Named("VInf", p.ValorInformado),
            sql.Named("VGlos", p.ValorGlosado),
        )
        if err != nil {
            tx.Rollback()
            return err
        }
    }

    return tx.Commit()
}

func (r *glosaDetalheRepository) ListarGlosas() ([]models.GlosaDetalhe, error) {
    lista := []models.GlosaDetalhe{}

    rows, err := r.db.Query(`
        SELECT GlosaDetalheId, Guia, DataOcorrencia, PacienteId, Mo, NomeContratado,
               DenteRegiao, MotivoGlosa, StatusRecurso, UnidadeId
        FROM GlosasDetalhes WITH (NOLOCK)
        ORDER BY DataOcorrencia DESC
    `)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var g models.GlosaDetalhe

        rows.Scan(
            &g.GlosaDetalheId,
            &g.Guia,
            &g.DataOcorrencia,
            &g.PacienteId,
            &g.Mo,
            &g.NomeContratado,
            &g.DenteRegiao,
            &g.MotivoGlosa,
            &g.StatusRecurso,
            &g.UnidadeId,
        )

        g.Procedimentos, _ = r.buscarProcedimentos(g.GlosaDetalheId)

        lista = append(lista, g)
    }

    return lista, nil
}

func (r *glosaDetalheRepository) BuscarPorId(id int) (*models.GlosaDetalhe, error) {
    var g models.GlosaDetalhe

    query := `
        SELECT GlosaDetalheId, Guia, DataOcorrencia, PacienteId, Mo,
               NomeContratado, DenteRegiao, MotivoGlosa, StatusRecurso, UnidadeId
        FROM GlosasDetalhes
        WHERE GlosaDetalheId = @Id
    `

    err := r.db.QueryRow(query, sql.Named("Id", id)).Scan(
        &g.GlosaDetalheId,
        &g.Guia,
        &g.DataOcorrencia,
        &g.PacienteId,
        &g.Mo,
        &g.NomeContratado,
        &g.DenteRegiao,
        &g.MotivoGlosa,
        &g.StatusRecurso,
        &g.UnidadeId,
    )

    if err == sql.ErrNoRows {
        return nil, nil
    }
    if err != nil {
        return nil, err
    }

    g.Procedimentos, _ = r.buscarProcedimentos(id)

    return &g, nil
}

func (r *glosaDetalheRepository) ExisteGlosa(g *models.GlosaDetalhe) (bool, error) {
    var count int
    query := `
        SELECT COUNT(1)
        FROM GlosasDetalhes
        WHERE Guia = @Guia
          AND PacienteId = @PacienteId
          AND DataOcorrencia = @DataOcorrencia
          AND UnidadeId = @UnidadeId
    `
    err := r.db.QueryRow(query,
        sql.Named("Guia", g.Guia),
        sql.Named("PacienteId", g.PacienteId),
        sql.Named("DataOcorrencia", g.DataOcorrencia),
        sql.Named("UnidadeId", g.UnidadeId),
    ).Scan(&count)

    return count > 0, err
}

func (r *glosaDetalheRepository) buscarProcedimentos(glosaId int) ([]models.GlosaProcedimento, error) {
    lista := []models.GlosaProcedimento{}

    rows, err := r.db.Query(`
        SELECT GlosaProcedimentoId, ProcedimentoId, CodigoProcedimento, NomeProcedimento,
               Quantidade, ValorInformado, ValorGlosado
        FROM GlosaProcedimentos
        WHERE GlosaDetalheId = @Id
    `, sql.Named("Id", glosaId))

    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var p models.GlosaProcedimento
        rows.Scan(
            &p.GlosaProcedimentoId,
            &p.ProcedimentoId,
            &p.CodigoProcedimento,
            &p.NomeProcedimento,
            &p.Quantidade,
            &p.ValorInformado,
            &p.ValorGlosado,
        )
        lista = append(lista, p)
    }

    return lista, nil
}
