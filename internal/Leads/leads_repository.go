package captacao

import (
	database "clinica-api/database"
	"clinica-api/internal/models"
	"database/sql"
)

type ContatoRepository interface {
	CriarContato(c *models.Leads) error
	EditarContato(c *models.Leads) error
	ListarContatos() ([]models.Leads, error)
	BuscarContatoPorID(id int) (*models.Leads, error)
	AtualizarStatus(id int, status string) error
}

type contatoRepository struct {
	db *sql.DB
}

func NovoContatoRepository(conn *database.SQLStr) ContatoRepository {
	return &contatoRepository{
		db: conn.DB(),
	}
}
func (r *contatoRepository) CriarContato(c *models.Leads) error {
	query := `
		INSERT INTO Leads
		(DataContato, MeioContato, Nome, Telefone, ProcedimentoInteresse, Observacoes, Status, FuncionarioId)
		VALUES (@DataContato, @MeioContato, @Nome, @Telefone, @ProcedimentoInteresse, @Observacoes, @Status, @FuncionarioId)
	`

	_, err := r.db.Exec(
		query,
		sql.Named("DataContato", c.DataContato),
		sql.Named("MeioContato", c.MeioContato),
		sql.Named("Nome", c.Nome),
		sql.Named("Telefone", c.Telefone),
		sql.Named("ProcedimentoInteresse", c.ProcedimentoInteresse),
		sql.Named("Observacoes", c.Observacoes),
		sql.Named("Status", c.Status),
		sql.Named("FuncionarioId", c.FuncionarioId),
	)

	return err
}
func (r *contatoRepository) EditarContato(c *models.Leads) error {
	query := `
		UPDATE Leads SET 
			DataContato = @DataContato,
			MeioContato = @MeioContato,
			Nome = @Nome,
			Telefone = @Telefone,
			ProcedimentoInteresse = @ProcedimentoInteresse,
			Observacoes = @Observacoes,
			Status = @Status,
			FuncionarioId = @FuncionarioId,
			DataAtualizacao = GETDATE()
		WHERE ContatoId = @ContatoId
	`

	_, err := r.db.Exec(
		query,
		sql.Named("DataContato", c.DataContato),
		sql.Named("MeioContato", c.MeioContato),
		sql.Named("Nome", c.Nome),
		sql.Named("Telefone", c.Telefone),
		sql.Named("ProcedimentoInteresse", c.ProcedimentoInteresse),
		sql.Named("Observacoes", c.Observacoes),
		sql.Named("Status", c.Status),
		sql.Named("FuncionarioId", c.FuncionarioId),
		sql.Named("ContatoId", c.ContatoId),
	)

	return err
}

func (r *contatoRepository) ListarContatos() ([]models.Leads, error) {
	query := `SELECT * FROM Leads ORDER BY DataAtualizacao DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lista []models.Leads

	for rows.Next() {
		var c models.Leads
		var funcionario sql.NullInt32
		var obs sql.NullString

		err := rows.Scan(
			&c.ContatoId,
			&c.DataContato,
			&c.MeioContato,
			&c.Nome,
			&c.Telefone,
			&c.ProcedimentoInteresse,
			&obs,
			&c.Status,
			&c.DataAtualizacao,
			&funcionario,
		)

		if err != nil {
			return nil, err
		}

		if funcionario.Valid {
			id := int(funcionario.Int32)
			c.FuncionarioId = &id
		}

		if obs.Valid {
			c.Observacoes = obs.String
		}

		lista = append(lista, c)
	}

	return lista, nil
}

func (r *contatoRepository) BuscarContatoPorID(id int) (*models.Leads, error) {
	query := `SELECT * FROM Leads WHERE ContatoId = @ContatoId`

	row := r.db.QueryRow(query, sql.Named("ContatoId", id))

	var c models.Leads
	var funcionario sql.NullInt32
	var obs sql.NullString

	err := row.Scan(
		&c.ContatoId,
		&c.DataContato,
		&c.MeioContato,
		&c.Nome,
		&c.Telefone,
		&c.ProcedimentoInteresse,
		&obs,
		&c.Status,
		&c.DataAtualizacao,
		&funcionario,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	if funcionario.Valid {
		idf := int(funcionario.Int32)
		c.FuncionarioId = &idf
	}

	if obs.Valid {
		c.Observacoes = obs.String
	}

	return &c, nil
}

func (r *contatoRepository) AtualizarStatus(id int, status string) error {
	query := `
		UPDATE Leads 
		SET Status = @Status, DataAtualizacao = GETDATE()
		WHERE ContatoId = @ContatoId
	`

	_, err := r.db.Exec(
		query,
		sql.Named("Status", status),
		sql.Named("ContatoId", id),
	)

	return err
}
