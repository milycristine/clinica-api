package laboratorio

import (
	"clinica-api/database"
	"clinica-api/internal/models"
	"database/sql"
	"fmt"
	"strings"
)

type LaboratorioRepository interface {
	CriarLaboratorio(l *models.Laboratorio) error
	EditarLaboratorio(l *models.Laboratorio) error
	ListarLaboratorios() ([]models.Laboratorio, error)
	BuscarLaboratorioPorID(id int) (*models.Laboratorio, error)
}

type laboratorioRepository struct {
	db *sql.DB
}

func NovoLaboratorioRepository(conn *database.SQLStr) LaboratorioRepository {
	return &laboratorioRepository{
		db: conn.DB(),
	}
}

func (r *laboratorioRepository) CriarLaboratorio(l *models.Laboratorio) error {
	query := `
        INSERT INTO Laboratorios (Nome, Contato, Telefone, Email)
        VALUES (@Nome, @Contato, @Telefone, @Email)
    `

	_, err := r.db.Exec(query,
		sql.Named("Nome", l.Nome),
		sql.Named("Contato", l.Contato),
		sql.Named("Telefone", l.Telefone),
		sql.Named("Email", l.Email),
	)

	if err != nil {
		return fmt.Errorf("erro ao criar laboratório: %w", err)
	}

	return nil
}

func (r *laboratorioRepository) EditarLaboratorio(l *models.Laboratorio) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("erro ao iniciar transação: %w", err)
	}

	var setClauses []string
	var args []any

	if l.Nome != "" {
		setClauses = append(setClauses, "Nome = @Nome")
		args = append(args, sql.Named("Nome", l.Nome))
	}
	if l.Contato != "" {
		setClauses = append(setClauses, "Contato = @Contato")
		args = append(args, sql.Named("Contato", l.Contato))
	}
	if l.Telefone != "" {
		setClauses = append(setClauses, "Telefone = @Telefone")
		args = append(args, sql.Named("Telefone", l.Telefone))
	}
	if l.Email != "" {
		setClauses = append(setClauses, "Email = @Email")
		args = append(args, sql.Named("Email", l.Email))
	}

	if len(setClauses) == 0 {
		tx.Rollback()
		return fmt.Errorf("nenhum campo para atualizar")
	}
	if l.Id == 0 {
		tx.Rollback()
		return fmt.Errorf("id do laboratório é obrigatório")
	}

	args = append(args, sql.Named("Id", l.Id))

	query := fmt.Sprintf(`
        UPDATE Laboratorios SET %s WHERE LaboratorioId = @Id
    `, strings.Join(setClauses, ", "))

	res, err := tx.Exec(query, args...)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("erro ao atualizar laboratório: %w", err)
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		tx.Rollback()
		return fmt.Errorf("nenhum laboratório atualizado (id=%d)", l.Id)
	}

	return tx.Commit()
}

func (r *laboratorioRepository) ListarLaboratorios() ([]models.Laboratorio, error) {
	lista := []models.Laboratorio{}

	rows, err := r.db.Query(`
        SELECT LaboratorioId, Nome, Contato, Telefone, Email
        FROM Laboratorios WITH (NOLOCK)
        ORDER BY Nome ASC
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var l models.Laboratorio
		if err := rows.Scan(&l.Id, &l.Nome, &l.Contato, &l.Telefone, &l.Email); err != nil {
			return nil, err
		}
		lista = append(lista, l)
	}

	return lista, nil
}

func (r *laboratorioRepository) BuscarLaboratorioPorID(id int) (*models.Laboratorio, error) {
	var l models.Laboratorio

	query := `
        SELECT LaboratorioId, Nome, Contato, Telefone, Email
        FROM Laboratorios
        WHERE LaboratorioId = @Id
    `

	err := r.db.QueryRow(query, sql.Named("Id", id)).Scan(
		&l.Id, &l.Nome, &l.Contato, &l.Telefone, &l.Email,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &l, nil
}
