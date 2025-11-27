package unidade

import (
	database "clinica-api/database"
	"clinica-api/internal/models"
	"database/sql"
	"fmt"
	"strings"
)

type UnidadeRepository interface {
	CriarUnidade(u *models.Unidade) error
	EditarUnidade(u *models.Unidade) error
	ListarUnidades() ([]models.Unidade, error)
	BuscarUnidadePorID(id int) (*models.Unidade, error)
	AtualizarStatus(id int, status int) error
}

type unidadeRepository struct {
	db *sql.DB
}

func NovaUnidadeRepository(conn *database.SQLStr) UnidadeRepository {
	return &unidadeRepository{
		db: conn.DB(),
	}
}

func (r *unidadeRepository) CriarUnidade(u *models.Unidade) error {
	query := `
        INSERT INTO Unidades (Nome, Endereco, Telefone, Email, Status)
        VALUES (@Nome, @Endereco, @Telefone, @Email, @Status)
    `
	_, err := r.db.Exec(query,
		sql.Named("Nome", u.Nome),
		sql.Named("Endereco", u.Endereco),
		sql.Named("Telefone", u.Telefone),
		sql.Named("Email", u.Email),
		sql.Named("Status", 1),
	)
	if err != nil {
		return fmt.Errorf("erro ao criar unidade: %w", err)
	}
	return nil
}

func (r *unidadeRepository) EditarUnidade(u *models.Unidade) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("erro ao iniciar transação: %w", err)
	}

	var setClauses []string
	var args []any

	if u.Nome != "" {
		setClauses = append(setClauses, "Nome = @Nome")
		args = append(args, sql.Named("Nome", u.Nome))
	}
	if u.Endereco != "" {
		setClauses = append(setClauses, "Endereco = @Endereco")
		args = append(args, sql.Named("Endereco", u.Endereco))
	}
	if u.Telefone != "" {
		setClauses = append(setClauses, "Telefone = @Telefone")
		args = append(args, sql.Named("Telefone", u.Telefone))
	}
	if u.Email != "" {
		setClauses = append(setClauses, "Email = @Email")
		args = append(args, sql.Named("Email", u.Email))
	}
	if u.Status != -1 {
		setClauses = append(setClauses, "Status = @Status")
		args = append(args, sql.Named("Status", u.Status))
	}

	if len(setClauses) == 0 {
		tx.Rollback()
		return fmt.Errorf("nenhum campo para atualizar")
	}

	if u.Id == 0 {
		tx.Rollback()
		return fmt.Errorf("id da unidade é obrigatório para atualização")
	}

	args = append(args, sql.Named("Id", u.Id))
	query := fmt.Sprintf("UPDATE Unidades SET %s WHERE Id = @Id", strings.Join(setClauses, ", "))

	res, err := tx.Exec(query, args...)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("erro ao atualizar unidade: %w", err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("erro ao verificar linhas afetadas: %w", err)
	}
	if rows == 0 {
		tx.Rollback()
		return fmt.Errorf("nenhuma unidade atualizada (id=%d)", u.Id)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("erro ao finalizar transação: %w", err)
	}

	return nil
}

func (r *unidadeRepository) ListarUnidades() ([]models.Unidade, error) {
	unidades := []models.Unidade{}

	rows, err := r.db.Query(`
        SELECT Id, Nome, Endereco, Telefone, Email, Status
        FROM Unidades WITH (NOLOCK)
        ORDER BY Nome ASC
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var u models.Unidade
		if err := rows.Scan(&u.Id, &u.Nome, &u.Endereco, &u.Telefone, &u.Email, &u.Status); err != nil {
			return nil, err
		}
		unidades = append(unidades, u)
	}

	return unidades, nil
}

func (r *unidadeRepository) BuscarUnidadePorID(id int) (*models.Unidade, error) {
	var u models.Unidade

	query := `
        SELECT Id, Nome, Endereco, Telefone, Email, Status
        FROM Unidades WHERE Id = @Id
    `
	err := r.db.QueryRow(query, sql.Named("Id", id)).
		Scan(&u.Id, &u.Nome, &u.Endereco, &u.Telefone, &u.Email, &u.Status)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *unidadeRepository) AtualizarStatus(id int, status int) error {
	query := `
        UPDATE Unidades SET Status = @Status
        WHERE Id = @Id
    `
	_, err := r.db.Exec(query,
		sql.Named("Id", id),
		sql.Named("Status", status),
	)
	return err
}
