package convenio

import (
	database "clinica-api/database"
	"clinica-api/internal/models"
	"database/sql"
	"fmt"
	"strings"
)

type ConvenioRepository interface {
	CriarConvenio(c *models.Convenio) (int, error)
	EditarConvenio(c *models.Convenio) error
	ListarConvenios() ([]models.Convenio, error)
	BuscarConvenioPorID(id int) (*models.Convenio, error)
	AtualizarStatus(
		id int,
		status string,
		obs string,
		valorPago *float64,
	) error
}

type convenioRepository struct {
	db *sql.DB
}

func NovoConvenioRepository(conn *database.SQLStr) ConvenioRepository {
	return &convenioRepository{
		db: conn.DB(),
	}
}

func (r *convenioRepository) CriarConvenio(c *models.Convenio) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	queryConvenio := `
		INSERT INTO Convenios (
			UnidadeId, PacienteId, Guia, Operadora, DataAtendimento, Status
		)
		OUTPUT INSERTED.ConvenioId
		VALUES (
			@UnidadeId, @PacienteId, @Guia, @Operadora, @DataAtendimento, 'ENVIADO'
		)
	`

	var convenioId int
	err = tx.QueryRow(
		queryConvenio,
		sql.Named("UnidadeId", c.UnidadeId),
		sql.Named("PacienteId", c.PacienteId),
		sql.Named("Guia", c.Guia),
		sql.Named("Operadora", c.Operadora),
		sql.Named("DataAtendimento", c.DataAtendimento),
	).Scan(&convenioId)

	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("erro ao criar convênio: %w", err)
	}

	queryProcedimento := `
		INSERT INTO ConvenioProcedimentos (
			ConvenioId, ProcedimentoId, CodigoProcedimento, NomeProcedimento,
			DenteRegiao, Quantidade, ValorInformado, Status, Observacoes
		)
		VALUES (
			@ConvenioId, @ProcedimentoId, @CodigoProcedimento, @NomeProcedimento,
			@DenteRegiao, @Quantidade, @ValorInformado, 'EM_ANALISE', @Observacoes
		)
	`

	for _, p := range c.Procedimentos {
		_, err = tx.Exec(
			queryProcedimento,
			sql.Named("ConvenioId", convenioId),
			sql.Named("ProcedimentoId", p.ProcedimentoId),
			sql.Named("CodigoProcedimento", p.CodigoProcedimento),
			sql.Named("NomeProcedimento", p.NomeProcedimento),
			sql.Named("DenteRegiao", p.DenteRegiao),
			sql.Named("Quantidade", p.Quantidade),
			sql.Named("ValorInformado", p.ValorInformado),
			sql.Named("Observacoes", p.Observacoes),
		)

		if err != nil {
			tx.Rollback()
			return 0, fmt.Errorf("erro ao criar procedimento do convênio: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return convenioId, nil
}

func (r *convenioRepository) EditarConvenio(c *models.Convenio) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	var setClauses []string
	var args []any

	if c.Guia != "" {
		setClauses = append(setClauses, "Guia = @Guia")
		args = append(args, sql.Named("Guia", c.Guia))
	}

	if c.Operadora != "" {
		setClauses = append(setClauses, "Operadora = @Operadora")
		args = append(args, sql.Named("Operadora", c.Operadora))
	}

	if c.DataAtendimento != "" {
		setClauses = append(setClauses, "DataAtendimento = @DataAtendimento")
		args = append(args, sql.Named("DataAtendimento", c.DataAtendimento))
	}

	if c.Status != "" {
		setClauses = append(setClauses, "Status = @Status")
		args = append(args, sql.Named("Status", c.Status))
	}

	if len(setClauses) == 0 {
		tx.Rollback()
		return fmt.Errorf("nenhum campo para atualizar")
	}

	args = append(args, sql.Named("ConvenioId", c.ConvenioId))

	query := fmt.Sprintf(
		"UPDATE Convenios SET %s WHERE ConvenioId = @ConvenioId",
		strings.Join(setClauses, ", "),
	)

	_, err = tx.Exec(query, args...)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *convenioRepository) ListarConvenios() ([]models.Convenio, error) {
	rows, err := r.db.Query(`
		SELECT ConvenioId, UnidadeId, PacienteId, Guia,
		       Operadora, DataAtendimento, Status, CreatedAt
		FROM Convenios
		ORDER BY CreatedAt DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var convenios []models.Convenio

	for rows.Next() {
		var c models.Convenio
		if err := rows.Scan(
			&c.ConvenioId,
			&c.UnidadeId,
			&c.PacienteId,
			&c.Guia,
			&c.Operadora,
			&c.DataAtendimento,
			&c.Status,
			&c.CreatedAt,
		); err != nil {
			return nil, err
		}
		convenios = append(convenios, c)
	}

	return convenios, nil
}

func (r *convenioRepository) BuscarConvenioPorID(id int) (*models.Convenio, error) {
	var c models.Convenio

	err := r.db.QueryRow(`
		SELECT ConvenioId, UnidadeId, PacienteId, Guia,
		       Operadora, DataAtendimento, Status, CreatedAt
		FROM Convenios
		WHERE ConvenioId = @Id
	`,
		sql.Named("Id", id),
	).Scan(
		&c.ConvenioId,
		&c.UnidadeId,
		&c.PacienteId,
		&c.Guia,
		&c.Operadora,
		&c.DataAtendimento,
		&c.Status,
		&c.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &c, err
}

func (r *convenioRepository) AtualizarStatus(
	id int,
	status string,
	obs string,
	valorPago *float64,
) error {
	_, err := r.db.Exec(`
		UPDATE ConvenioProcedimento 
		SET Status = @Status,
		    Observacoes = @Obs,
		    ValorPago = @ValorPago
		WHERE ConvenioProcedimentoId = @Id`,
		sql.Named("Id", id),
		sql.Named("Status", status),
		sql.Named("Obs", obs),
		sql.Named("ValorPago", valorPago),
	)
	return err
}
