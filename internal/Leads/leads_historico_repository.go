package captacao

import (
    database "clinica-api/database"
    "clinica-api/internal/models"
    "database/sql"
)

type LeadHistoricoRepository interface {
    CriarHistorico(h *models.LeadHistorico) error
    ListarHistoricoPorLead(contatoId int) ([]models.LeadHistorico, error)
}

type leadHistoricoRepository struct {
    db *sql.DB
}

func NovoLeadHistoricoRepository(conn *database.SQLStr) LeadHistoricoRepository {
    return &leadHistoricoRepository{
        db: conn.DB(),
    }
}

func (r *leadHistoricoRepository) CriarHistorico(h *models.LeadHistorico) error {
    query := `
        INSERT INTO LeadHistorico (ContatoId, Descricao, FuncionarioId)
        VALUES (@ContatoId, @Descricao, @FuncionarioId)
    `
    _, err := r.db.Exec(
        query,
        sql.Named("ContatoId", h.ContatoId),
        sql.Named("Descricao", h.Descricao),
        sql.Named("FuncionarioId", h.FuncionarioId),
    )
    return err
}

func (r *leadHistoricoRepository) ListarHistoricoPorLead(contatoId int) ([]models.LeadHistorico, error) {
    query := `
        SELECT HistoricoId, ContatoId, DataRegistro, Descricao, FuncionarioId
        FROM LeadHistorico
        WHERE ContatoId = @ContatoId
        ORDER BY DataRegistro DESC
    `

    rows, err := r.db.Query(query, sql.Named("ContatoId", contatoId))
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var lista []models.LeadHistorico

    for rows.Next() {
        var h models.LeadHistorico
        var funcionario sql.NullInt32

        err := rows.Scan(
            &h.HistoricoId,
            &h.ContatoId,
            &h.DataRegistro,
            &h.Descricao,
            &funcionario,
        )
        if err != nil {
            return nil, err
        }

        if funcionario.Valid {
            id := int(funcionario.Int32)
            h.FuncionarioId = &id
        }

        lista = append(lista, h)
    }

    return lista, nil
}
