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
	ListarLaboratorios(filtro models.LaboratorioFiltro) ([]models.Laboratorio, int, error)
	BuscarLaboratorioPorID(id int) (*models.Laboratorio, error)
	ExisteDuplicado(nome, contato string) (bool, error)
}

type laboratorioRepository struct {
	db *sql.DB
}

func NovoLaboratorioRepository(conn *database.SQLStr) LaboratorioRepository {
	return &laboratorioRepository{
		db: conn.DB(),
	}
}

func (r *laboratorioRepository) ExisteDuplicado(nome, contato string) (bool, error) {
	var count int

	err := r.db.QueryRow(`
        SELECT COUNT(1)
        FROM Laboratorios WITH (NOLOCK)
        WHERE Nome = @Nome AND Contato = @Contato
    `,
		sql.Named("Nome", nome),
		sql.Named("Contato", contato),
	).Scan(&count)

	return count > 0, err
}

func (r *laboratorioRepository) CriarLaboratorio(l *models.Laboratorio) error {

	duplicado, err := r.ExisteDuplicado(l.Nome, l.Contato)
	if err != nil {
		return err
	}
	if duplicado {
		return fmt.Errorf("laboratório já cadastrado com esse nome e contato")
	}

	query := `
        INSERT INTO Laboratorios (Nome, Contato, Telefone, Email)
        VALUES (@Nome, @Contato, @Telefone, @Email)
    `
	_, err = r.db.Exec(query,
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

func (r *laboratorioRepository) ListarLaboratorios(filtro models.LaboratorioFiltro) ([]models.Laboratorio, int, error) {

	query := `
        SELECT LaboratorioId, Nome, Contato, Telefone, Email
        FROM Laboratorios WITH (NOLOCK)
        WHERE 1=1
    `
	countQuery := `
        SELECT COUNT(1)
        FROM Laboratorios WITH (NOLOCK)
        WHERE 1=1
    `

	var args []any
	var conditions []string

	if filtro.Nome != "" {
		conditions = append(conditions, "Nome LIKE @Nome")
		args = append(args, sql.Named("Nome", "%"+filtro.Nome+"%"))
	}
	if filtro.Contato != "" {
		conditions = append(conditions, "Contato LIKE @Contato")
		args = append(args, sql.Named("Contato", "%"+filtro.Contato+"%"))
	}
	if filtro.Email != "" {
		conditions = append(conditions, "Email LIKE @Email")
		args = append(args, sql.Named("Email", "%"+filtro.Email+"%"))
	}

	if len(conditions) > 0 {
		where := " AND " + strings.Join(conditions, " AND ")
		query += where
		countQuery += where
	}

	query += `
        ORDER BY Nome ASC
        OFFSET @Offset ROWS
        FETCH NEXT @Limit ROWS ONLY
    `

	args = append(args,
		sql.Named("Offset", (filtro.Page-1)*filtro.Limit),
		sql.Named("Limit", filtro.Limit),
	)

	var total int
	err := r.db.QueryRow(countQuery, args[:len(args)-2]...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	lista := []models.Laboratorio{}
	for rows.Next() {
		var l models.Laboratorio
		if err := rows.Scan(&l.Id, &l.Nome, &l.Contato, &l.Telefone, &l.Email); err != nil {
			return nil, 0, err
		}
		lista = append(lista, l)
	}

	return lista, total, nil
}

func (r *laboratorioRepository) BuscarLaboratorioPorID(id int) (*models.Laboratorio, error) {
	var l models.Laboratorio

	err := r.db.QueryRow(`
        SELECT LaboratorioId, Nome, Contato, Telefone, Email
        FROM Laboratorios WITH (NOLOCK)
        WHERE LaboratorioId = @Id
    `, sql.Named("Id", id)).Scan(
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
