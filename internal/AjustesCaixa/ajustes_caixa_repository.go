package ajustecaixa

import (
	database "clinica-api/database"
	"clinica-api/internal/models"
	"database/sql"
	"fmt"
)

type AjusteCaixaRepository interface {
	Criar(a *models.AjusteCaixa) error
	Listar() ([]models.AjusteCaixa, error)
}

type ajusteCaixaRepository struct {
	db *sql.DB
}

func NovoAjusteCaixaRepository(conn *database.SQLStr) AjusteCaixaRepository {
	return &ajusteCaixaRepository{db: conn.DB()}
}

func (r *ajusteCaixaRepository) Criar(a *models.AjusteCaixa) error {
	query := `
		INSERT INTO AjustesCaixa
		(Data, Tipo, Descricao, Valor, Origem, Destino, Status, UnidadeId)
		VALUES (@Data, @Tipo, @Descricao, @Valor, @Origem, @Destino, 'PENDENTE', @UnidadeId)
	`

	_, err := r.db.Exec(
		query,
		sql.Named("Data", a.Data),
		sql.Named("Tipo", a.Tipo),
		sql.Named("Descricao", a.Descricao),
		sql.Named("Valor", a.Valor),
		sql.Named("Origem", a.Origem),
		sql.Named("Destino", a.Destino),
		sql.Named("UnidadeId", a.UnidadeId),
	)

	if err != nil {
		return fmt.Errorf("erro ao criar ajuste de caixa: %w", err)
	}
	return nil
}

func (r *ajusteCaixaRepository) Listar() ([]models.AjusteCaixa, error) {
	var lista []models.AjusteCaixa

	rows, err := r.db.Query(`
		SELECT AjusteId, Data, Tipo, Descricao, Valor, Origem, Destino, Status, UnidadeId
		FROM AjustesCaixa WITH (NOLOCK)
		ORDER BY Data DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var a models.AjusteCaixa
		if err := rows.Scan(
			&a.AjusteId,
			&a.Data,
			&a.Tipo,
			&a.Descricao,
			&a.Valor,
			&a.Origem,
			&a.Destino,
			&a.Status,
			&a.UnidadeId,
		); err != nil {
			return nil, err
		}
		lista = append(lista, a)
	}

	return lista, nil
}
