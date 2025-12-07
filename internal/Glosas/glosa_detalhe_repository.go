package glosa

import (
	database "clinica-api/database"
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
	query := `
        INSERT INTO GlosasDetalhes
        (Guia, DataOcorrencia, PacienteId, Mo, NomeContratado, CodigoProcedimento, 
        NomeProcedimento, DenteRegiao, ValorInformado, ValorGlosado, MotivoGlosa, StatusRecurso)
        VALUES
        (@Guia, @DataOcorrencia, @PacienteId, @Mo, @NomeContratado, @CodigoProcedimento,
        @NomeProcedimento, @DenteRegiao, @ValorInformado, @ValorGlosado, @MotivoGlosa, @StatusRecurso)
    `

	_, err := r.db.Exec(query,
		sql.Named("Guia", g.Guia),
		sql.Named("DataOcorrencia", g.DataOcorrencia),
		sql.Named("PacienteId", g.PacienteId),
		sql.Named("Mo", g.Mo),
		sql.Named("NomeContratado", g.NomeContratado),
		sql.Named("CodigoProcedimento", g.CodigoProcedimento),
		sql.Named("NomeProcedimento", g.NomeProcedimento),
		sql.Named("DenteRegiao", g.DenteRegiao),
		sql.Named("ValorInformado", g.ValorInformado),
		sql.Named("ValorGlosado", g.ValorGlosado),
		sql.Named("MotivoGlosa", g.MotivoGlosa),
		sql.Named("StatusRecurso", g.StatusRecurso),
	)

	if err != nil {
		return fmt.Errorf("erro ao criar glosa: %w", err)
	}

	return nil
}

func (r *glosaDetalheRepository) EditarGlosa(g *models.GlosaDetalhe) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("erro ao iniciar transação: %w", err)
	}

	var setClauses []string
	var args []any

	if g.Guia != "" {
		setClauses = append(setClauses, "Guia = @Guia")
		args = append(args, sql.Named("Guia", g.Guia))
	}
	if g.DataOcorrencia != "" {
		setClauses = append(setClauses, "DataOcorrencia = @DataOcorrencia")
		args = append(args, sql.Named("DataOcorrencia", g.DataOcorrencia))
	}
	if g.PacienteId != 0 {
		setClauses = append(setClauses, "PacienteId = @PacienteId")
		args = append(args, sql.Named("PacienteId", g.PacienteId))
	}
	if g.Mo != "" {
		setClauses = append(setClauses, "Mo = @Mo")
		args = append(args, sql.Named("Mo", g.Mo))
	}
	if g.NomeContratado != "" {
		setClauses = append(setClauses, "NomeContratado = @NomeContratado")
		args = append(args, sql.Named("NomeContratado", g.NomeContratado))
	}
	if g.CodigoProcedimento != "" {
		setClauses = append(setClauses, "CodigoProcedimento = @CodigoProcedimento")
		args = append(args, sql.Named("CodigoProcedimento", g.CodigoProcedimento))
	}
	if g.NomeProcedimento != "" {
		setClauses = append(setClauses, "NomeProcedimento = @NomeProcedimento")
		args = append(args, sql.Named("NomeProcedimento", g.NomeProcedimento))
	}
	if g.DenteRegiao != "" {
		setClauses = append(setClauses, "DenteRegiao = @DenteRegiao")
		args = append(args, sql.Named("DenteRegiao", g.DenteRegiao))
	}
	if g.ValorInformado != 0 {
		setClauses = append(setClauses, "ValorInformado = @ValorInformado")
		args = append(args, sql.Named("ValorInformado", g.ValorInformado))
	}
	if g.ValorGlosado != 0 {
		setClauses = append(setClauses, "ValorGlosado = @ValorGlosado")
		args = append(args, sql.Named("ValorGlosado", g.ValorGlosado))
	}
	if g.MotivoGlosa != "" {
		setClauses = append(setClauses, "MotivoGlosa = @MotivoGlosa")
		args = append(args, sql.Named("MotivoGlosa", g.MotivoGlosa))
	}
	if g.StatusRecurso != "" {
		setClauses = append(setClauses, "StatusRecurso = @StatusRecurso")
		args = append(args, sql.Named("StatusRecurso", g.StatusRecurso))
	}

	if len(setClauses) == 0 {
		tx.Rollback()
		return fmt.Errorf("nenhum campo informado para atualização")
	}

	if g.GlosaDetalheId == 0 {
		tx.Rollback()
		return fmt.Errorf("id da glosa é obrigatório para atualização")
	}

	args = append(args, sql.Named("Id", g.GlosaDetalheId))

	query := fmt.Sprintf("UPDATE GlosasDetalhes SET %s WHERE GlosaDetalheId = @Id",
		strings.Join(setClauses, ", "))

	_, err = tx.Exec(query, args...)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("erro ao atualizar glosa: %w", err)
	}

	return tx.Commit()
}

func (r *glosaDetalheRepository) ListarGlosas() ([]models.GlosaDetalhe, error) {
	lista := []models.GlosaDetalhe{}

	rows, err := r.db.Query(`
        SELECT GlosaDetalheId, Guia, DataOcorrencia, PacienteId, Mo,
               NomeContratado, CodigoProcedimento, NomeProcedimento,
               DenteRegiao, ValorInformado, ValorGlosado, MotivoGlosa, StatusRecurso
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
			&g.CodigoProcedimento,
			&g.NomeProcedimento,
			&g.DenteRegiao,
			&g.ValorInformado,
			&g.ValorGlosado,
			&g.MotivoGlosa,
			&g.StatusRecurso,
		)
		lista = append(lista, g)
	}

	return lista, nil
}

func (r *glosaDetalheRepository) BuscarPorId(id int) (*models.GlosaDetalhe, error) {
	var g models.GlosaDetalhe

	query := `
        SELECT GlosaDetalheId, Guia, DataOcorrencia, PacienteId, Mo,
               NomeContratado, CodigoProcedimento, NomeProcedimento,
               DenteRegiao, ValorInformado, ValorGlosado, MotivoGlosa, StatusRecurso
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
		&g.CodigoProcedimento,
		&g.NomeProcedimento,
		&g.DenteRegiao,
		&g.ValorInformado,
		&g.ValorGlosado,
		&g.MotivoGlosa,
		&g.StatusRecurso,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &g, nil
}
