package fechamentolaboratorio

import (
	"clinica-api/database"
	"clinica-api/internal/models"
	"database/sql"
	"fmt"
	"strings"
)

type FechamentoRepository interface {
    Criar(f *models.LaboratorioFechamento) error
    Editar(f *models.LaboratorioFechamento) error
    Listar() ([]models.LaboratorioFechamento, error)
    BuscarPorID(id int) (*models.LaboratorioFechamento, error)
    ListarPorLaboratorio(labId int) ([]models.LaboratorioFechamento, error)
    ListarPorData(data string) ([]models.LaboratorioFechamento, error)
    ListarPorMes(mes int, ano int) ([]models.LaboratorioFechamento, error)
    ListarPorPeriodo(inicio string, fim string) ([]models.LaboratorioFechamento, error)
    ListarPorLaboratorioMes(labId int, mes int, ano int) ([]models.LaboratorioFechamento, error)
    ListarPorPaciente(id int) ([]models.LaboratorioFechamento, error)
}

type fechamentoRepository struct {
    db *sql.DB
}

func NovoFechamentoRepository(conn *database.SQLStr) FechamentoRepository {
    return &fechamentoRepository{
        db: conn.DB(),
    }
}

func (r *fechamentoRepository) Criar(f *models.LaboratorioFechamento) error {
    query := `
        INSERT INTO LaboratorioFechamento
        (LaboratorioId, DataFinalizacao, Produto, PacienteId, Quantidade, PrecoUnitario)
        VALUES (@LaboratorioId, @DataFinalizacao, @Produto, @PacienteId, @Quantidade, @PrecoUnitario)
    `

    var pacienteParam interface{}
    if f.PacienteId == nil {
        pacienteParam = nil
    } else {
        pacienteParam = *f.PacienteId
    }

    _, err := r.db.Exec(query,
        sql.Named("LaboratorioId", f.LaboratorioId),
        sql.Named("DataFinalizacao", f.DataFinalizacao),
        sql.Named("Produto", f.Produto),
        sql.Named("PacienteId", pacienteParam),
        sql.Named("Quantidade", f.Quantidade),
        sql.Named("PrecoUnitario", f.PrecoUnitario),
    )
    if err != nil {
        return fmt.Errorf("erro ao criar fechamento: %w", err)
    }
    return nil
}

func (r *fechamentoRepository) Editar(f *models.LaboratorioFechamento) error {
    if f.FechamentoId == 0 {
        return fmt.Errorf("fechamento_id é obrigatório")
    }

    var setClauses []string
    var args []interface{}

    if f.LaboratorioId != 0 {
        setClauses = append(setClauses, "LaboratorioId = @LaboratorioId")
        args = append(args, sql.Named("LaboratorioId", f.LaboratorioId))
    }
    if f.DataFinalizacao != "" {
        setClauses = append(setClauses, "DataFinalizacao = @DataFinalizacao")
        args = append(args, sql.Named("DataFinalizacao", f.DataFinalizacao))
    }
    if f.Produto != "" {
        setClauses = append(setClauses, "Produto = @Produto")
        args = append(args, sql.Named("Produto", f.Produto))
    }
    if f.PacienteId != nil {
        setClauses = append(setClauses, "PacienteId = @PacienteId")
        args = append(args, sql.Named("PacienteId", *f.PacienteId))
    }
    if f.Quantidade != 0 {
        setClauses = append(setClauses, "Quantidade = @Quantidade")
        args = append(args, sql.Named("Quantidade", f.Quantidade))
    }
    if f.PrecoUnitario != 0 {
        setClauses = append(setClauses, "PrecoUnitario = @PrecoUnitario")
        args = append(args, sql.Named("PrecoUnitario", f.PrecoUnitario))
    }

    if len(setClauses) == 0 {
        return fmt.Errorf("nenhum campo para atualizar")
    }

    args = append(args, sql.Named("FechamentoId", f.FechamentoId))

    query := fmt.Sprintf(`
        UPDATE LaboratorioFechamento
        SET %s
        WHERE FechamentoId = @FechamentoId
    `, strings.Join(setClauses, ", "))

    res, err := r.db.Exec(query, args...)
    if err != nil {
        return fmt.Errorf("erro ao editar fechamento: %w", err)
    }

    rows, err := res.RowsAffected()
    if err != nil {
        return fmt.Errorf("erro ao verificar linhas afetadas: %w", err)
    }
    if rows == 0 {
        return fmt.Errorf("nenhum fechamento atualizado (id=%d)", f.FechamentoId)
    }

    return nil
}

func (r *fechamentoRepository) Listar() ([]models.LaboratorioFechamento, error) {
    lista := []models.LaboratorioFechamento{}

    query := `
        SELECT FechamentoId, LaboratorioId, DataFinalizacao, Produto, PacienteId,
               Quantidade, PrecoUnitario, Total
        FROM LaboratorioFechamento WITH (NOLOCK)
        ORDER BY DataFinalizacao DESC, FechamentoId DESC
    `

    rows, err := r.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var f models.LaboratorioFechamento
        var paciente sql.NullInt64
        if err := rows.Scan(
            &f.FechamentoId, &f.LaboratorioId, &f.DataFinalizacao, &f.Produto,
            &paciente, &f.Quantidade, &f.PrecoUnitario, &f.Total,
        ); err != nil {
            return nil, err                                                                                             
        }
        if paciente.Valid {
            pid := int(paciente.Int64)
            f.PacienteId = &pid
        } else {
            f.PacienteId = nil
        }
        lista = append(lista, f)
    }

    return lista, nil
}

func (r *fechamentoRepository) BuscarPorID(id int) (*models.LaboratorioFechamento, error) {
    var f models.LaboratorioFechamento
    var paciente sql.NullInt64

    query := `
        SELECT FechamentoId, LaboratorioId, DataFinalizacao, Produto, PacienteId,
               Quantidade, PrecoUnitario, Total
        FROM LaboratorioFechamento
        WHERE FechamentoId = @FechamentoId
    `

    err := r.db.QueryRow(query, sql.Named("FechamentoId", id)).Scan(
        &f.FechamentoId, &f.LaboratorioId, &f.DataFinalizacao, &f.Produto,
        &paciente, &f.Quantidade, &f.PrecoUnitario, &f.Total,
    )
    if err == sql.ErrNoRows {
        return nil, nil
    }
    if err != nil {
        return nil, err
    }
    if paciente.Valid {
        pid := int(paciente.Int64)
        f.PacienteId = &pid
    } else {
        f.PacienteId = nil
    }
    return &f, nil
}

func (r *fechamentoRepository) ListarPorLaboratorio(labId int) ([]models.LaboratorioFechamento, error) {
    lista := []models.LaboratorioFechamento{}

    query := `
        SELECT FechamentoId, LaboratorioId, DataFinalizacao, Produto, PacienteId,
               Quantidade, PrecoUnitario, Total
        FROM LaboratorioFechamento WITH (NOLOCK)
        WHERE LaboratorioId = @LaboratorioId
        ORDER BY DataFinalizacao DESC, FechamentoId DESC
    `
    rows, err := r.db.Query(query, sql.Named("LaboratorioId", labId))
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var f models.LaboratorioFechamento
        var paciente sql.NullInt64
        if err := rows.Scan(
            &f.FechamentoId, &f.LaboratorioId, &f.DataFinalizacao, &f.Produto,
            &paciente, &f.Quantidade, &f.PrecoUnitario, &f.Total,
        ); err != nil {
            return nil, err
        }
        if paciente.Valid {
            pid := int(paciente.Int64)
            f.PacienteId = &pid
        } else {
            f.PacienteId = nil
        }
        lista = append(lista, f)
    }

    return lista, nil
}

func (r *fechamentoRepository) ListarPorData(data string) ([]models.LaboratorioFechamento, error) {
    lista := []models.LaboratorioFechamento{}

    query := `
        SELECT FechamentoId, LaboratorioId, DataFinalizacao, Produto, PacienteId,
               Quantidade, PrecoUnitario, Total
        FROM LaboratorioFechamento WITH (NOLOCK)
        WHERE DataFinalizacao = @DataFinalizacao
        ORDER BY FechamentoId DESC
    `
    rows, err := r.db.Query(query, sql.Named("DataFinalizacao", data))
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var f models.LaboratorioFechamento
        var paciente sql.NullInt64
        if err := rows.Scan(
            &f.FechamentoId, &f.LaboratorioId, &f.DataFinalizacao, &f.Produto,
            &paciente, &f.Quantidade, &f.PrecoUnitario, &f.Total,
        ); err != nil {
            return nil, err
        }
        if paciente.Valid {
            pid := int(paciente.Int64)
            f.PacienteId = &pid
        } else {
            f.PacienteId = nil
        }
        lista = append(lista, f)
    }

    return lista, nil
}


func (r *fechamentoRepository) ListarPorMes(mes int, ano int) ([]models.LaboratorioFechamento, error) {
    lista := []models.LaboratorioFechamento{}

    query := `
        SELECT FechamentoId, LaboratorioId, DataFinalizacao, Produto,
               PacienteId, Quantidade, PrecoUnitario,
               (Quantidade * PrecoUnitario) as Total
        FROM LaboratorioFechamento
        WHERE MONTH(DataFinalizacao) = @Mes
          AND YEAR(DataFinalizacao) = @Ano
        ORDER BY DataFinalizacao DESC
    `

    rows, err := r.db.Query(query,
        sql.Named("Mes", mes),
        sql.Named("Ano", ano),
    )
    if err != nil { return nil, err }
    defer rows.Close()

    for rows.Next() {
        var f models.LaboratorioFechamento
        rows.Scan(
            &f.FechamentoId, &f.LaboratorioId, &f.DataFinalizacao,
            &f.Produto, &f.PacienteId, &f.Quantidade,
            &f.PrecoUnitario, &f.Total,
        )
        lista = append(lista, f)
    }

    return lista, nil
}


func (r *fechamentoRepository) ListarPorPeriodo(inicio string, fim string) ([]models.LaboratorioFechamento, error) {
    lista := []models.LaboratorioFechamento{}

    query := `
        SELECT FechamentoId, LaboratorioId, DataFinalizacao, Produto,
               PacienteId, Quantidade, PrecoUnitario,
               (Quantidade * PrecoUnitario) as Total
        FROM LaboratorioFechamento
        WHERE CONVERT(date, DataFinalizacao)
              BETWEEN @Inicio AND @Fim
        ORDER BY DataFinalizacao DESC
    `

    rows, err := r.db.Query(query,
        sql.Named("Inicio", inicio),
        sql.Named("Fim", fim),
    )
    if err != nil { return nil, err }
    defer rows.Close()

    for rows.Next() {
        var f models.LaboratorioFechamento
        rows.Scan(
            &f.FechamentoId, &f.LaboratorioId, &f.DataFinalizacao,
            &f.Produto, &f.PacienteId, &f.Quantidade,
            &f.PrecoUnitario, &f.Total,
        )
        lista = append(lista, f)
    }

    return lista, nil
}


func (r *fechamentoRepository) ListarPorLaboratorioMes(labId int, mes int, ano int) ([]models.LaboratorioFechamento, error) {
    lista := []models.LaboratorioFechamento{}

    query := `
        SELECT FechamentoId, LaboratorioId, DataFinalizacao, Produto,
               PacienteId, Quantidade, PrecoUnitario,
               (Quantidade * PrecoUnitario) as Total
        FROM LaboratorioFechamento
        WHERE LaboratorioId = @Lab
          AND MONTH(DataFinalizacao) = @Mes
          AND YEAR(DataFinalizacao) = @Ano
        ORDER BY DataFinalizacao DESC
    `

    rows, err := r.db.Query(query,
        sql.Named("Lab", labId),
        sql.Named("Mes", mes),
        sql.Named("Ano", ano),
    )
    if err != nil { return nil, err }
    defer rows.Close()

    for rows.Next() {
        var f models.LaboratorioFechamento
        rows.Scan(
            &f.FechamentoId, &f.LaboratorioId, &f.DataFinalizacao,
            &f.Produto, &f.PacienteId, &f.Quantidade,
            &f.PrecoUnitario, &f.Total,
        )
        lista = append(lista, f)
    }

    return lista, nil
}


func (r *fechamentoRepository) ListarPorPaciente(id int) ([]models.LaboratorioFechamento, error) {
    lista := []models.LaboratorioFechamento{}

    query := `
        SELECT FechamentoId, LaboratorioId, DataFinalizacao, Produto,
               PacienteId, Quantidade, PrecoUnitario,
               (Quantidade * PrecoUnitario) as Total
        FROM LaboratorioFechamento
        WHERE PacienteId = @Id
        ORDER BY DataFinalizacao DESC
    `

    rows, err := r.db.Query(query, sql.Named("Id", id))
    if err != nil { return nil, err }
    defer rows.Close()

    for rows.Next() {
        var f models.LaboratorioFechamento
        rows.Scan(
            &f.FechamentoId, &f.LaboratorioId, &f.DataFinalizacao,
            &f.Produto, &f.PacienteId, &f.Quantidade,
            &f.PrecoUnitario, &f.Total,
        )
        lista = append(lista, f)
    }

    return lista, nil
}