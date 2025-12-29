package repasse

import (
	database "clinica-api/database"
	"clinica-api/internal/models"
	"database/sql"
	"fmt"
)

type RepasseRepository interface {
	Criar(r *models.Repasse) error
	Listar() ([]models.Repasse, error)
	BuscarPorID(id int) (*models.Repasse, error)
	AtualizarStatus(id int, status string) error
}

type repasseRepository struct {
	db *sql.DB
}

func NovoRepasseRepository(conn *database.SQLStr) RepasseRepository {
	return &repasseRepository{db: conn.DB()}
}

func (r *repasseRepository) Criar(rep *models.Repasse) error {
	query := `
		INSERT INTO Repasses
		(ReceitaId, ProfissionalId, Percentual, ValorFixo, ValorCalculado, Status, Observacoes)
		VALUES (@ReceitaId, @ProfissionalId, @Percentual, @ValorFixo, @ValorCalculado, 'A_PAGAR', @Observacoes)
	`

	_, err := r.db.Exec(
		query,
		sql.Named("ReceitaId", rep.ReceitaId),
		sql.Named("ProfissionalId", rep.ProfissionalId),
		sql.Named("Percentual", rep.Percentual),
		sql.Named("ValorFixo", rep.ValorFixo),
		sql.Named("ValorCalculado", rep.ValorCalculado),
		sql.Named("Observacoes", rep.Observacoes),
	)

	if err != nil {
		return fmt.Errorf("erro ao criar repasse: %w", err)
	}
	return nil
}

func (r *repasseRepository) Listar() ([]models.Repasse, error) {
	var lista []models.Repasse

	rows, err := r.db.Query(`
		SELECT RepasseId, ReceitaId, ProfissionalId, Percentual, ValorFixo,
		       ValorCalculado, Status, Observacoes
		FROM Repasses WITH (NOLOCK)
		ORDER BY RepasseId DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var rep models.Repasse
		if err := rows.Scan(
			&rep.RepasseId,
			&rep.ReceitaId,
			&rep.ProfissionalId,
			&rep.Percentual,
			&rep.ValorFixo,
			&rep.ValorCalculado,
			&rep.Status,
			&rep.Observacoes,
		); err != nil {
			return nil, err
		}
		lista = append(lista, rep)
	}

	return lista, nil
}

func (r *repasseRepository) BuscarPorID(id int) (*models.Repasse, error) {
	var rep models.Repasse

	err := r.db.QueryRow(`
		SELECT RepasseId, ReceitaId, ProfissionalId, Percentual, ValorFixo,
		       ValorCalculado, Status, Observacoes
		FROM Repasses WHERE RepasseId = @Id
	`, sql.Named("Id", id)).
		Scan(
			&rep.RepasseId,
			&rep.ReceitaId,
			&rep.ProfissionalId,
			&rep.Percentual,
			&rep.ValorFixo,
			&rep.ValorCalculado,
			&rep.Status,
			&rep.Observacoes,
		)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &rep, nil
}

func (r *repasseRepository) AtualizarStatus(id int, status string) error {
	_, err := r.db.Exec(
		`UPDATE Repasses SET Status = @Status WHERE RepasseId = @Id`,
		sql.Named("Id", id),
		sql.Named("Status", status),
	)
	return err
}
