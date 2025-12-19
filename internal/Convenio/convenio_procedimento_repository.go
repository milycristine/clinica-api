package convenio

import (
	database "clinica-api/database"
	"clinica-api/internal/models"
	"database/sql"
)

type ConvenioProcedimentoRepository interface {
	Criar(p *models.ConvenioProcedimento) error
	ListarPorConvenio(convenioId int) ([]models.ConvenioProcedimento, error)
	AtualizarStatus(
		id int,
		status string,
		obs string,
		valorPago *float64,
	) error
	BuscarPorID(id int) (*models.ConvenioProcedimento, error)
}


type convenioProcedimentoRepository struct {
	db *sql.DB
}

func NovoConvenioProcedimentoRepository(conn *database.SQLStr) ConvenioProcedimentoRepository {
	return &convenioProcedimentoRepository{
		db: conn.DB(),
	}
}

func (r *convenioProcedimentoRepository) Criar(p *models.ConvenioProcedimento) error {
	query := `
		INSERT INTO ConvenioProcedimentos (
			ConvenioId, ProcedimentoId, CodigoProcedimento, NomeProcedimento,
			DenteRegiao, Quantidade, ValorInformado, Status, Observacoes
		)
		VALUES (
			@ConvenioId, @ProcedimentoId, @CodigoProcedimento, @NomeProcedimento,
			@DenteRegiao, @Quantidade, @ValorInformado, 'EM_ANALISE', @Observacoes
		)
	`

	_, err := r.db.Exec(
		query,
		sql.Named("ConvenioId", p.ConvenioId),
		sql.Named("ProcedimentoId", p.ProcedimentoId),
		sql.Named("CodigoProcedimento", p.CodigoProcedimento),
		sql.Named("NomeProcedimento", p.NomeProcedimento),
		sql.Named("DenteRegiao", p.DenteRegiao),
		sql.Named("Quantidade", p.Quantidade),
		sql.Named("ValorInformado", p.ValorInformado),
		sql.Named("Observacoes", p.Observacoes),
	)

	return err
}

func (r *convenioProcedimentoRepository) ListarPorConvenio(convenioId int) ([]models.ConvenioProcedimento, error) {
	rows, err := r.db.Query(`
		SELECT ConvenioProcedimentoId, ConvenioId, ProcedimentoId,
		       CodigoProcedimento, NomeProcedimento, DenteRegiao,
		       Quantidade, ValorInformado, ValorPago, Status, Observacoes
		FROM ConvenioProcedimentos
		WHERE ConvenioId = @ConvenioId
	`,
		sql.Named("ConvenioId", convenioId),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lista []models.ConvenioProcedimento

	for rows.Next() {
		var p models.ConvenioProcedimento
		rows.Scan(
			&p.ConvenioProcedimentoId,
			&p.ConvenioId,
			&p.ProcedimentoId,
			&p.CodigoProcedimento,
			&p.NomeProcedimento,
			&p.DenteRegiao,
			&p.Quantidade,
			&p.ValorInformado,
			&p.ValorPago,
			&p.Status,
			&p.Observacoes,
		)
		lista = append(lista, p)
	}

	return lista, nil
}
func (r *convenioProcedimentoRepository) AtualizarStatus(id int, status, obs string, valorPago *float64) error {
	_, err := r.db.Exec(`
UPDATE ConvenioProcedimentos
SET Status = @Status,
Observacoes = @Obs,
ValorPago = @ValorPago
WHERE ConvenioProcedimentoId = @Id
`,
		sql.Named("Id", id),
		sql.Named("Status", status),
		sql.Named("Obs", obs),
		sql.Named("ValorPago", valorPago),
	)
	return err
}

func (r *convenioProcedimentoRepository) BuscarPorID(id int) (*models.ConvenioProcedimento, error) {
	var p models.ConvenioProcedimento

	err := r.db.QueryRow(`
		SELECT ConvenioProcedimentoId, ConvenioId, ProcedimentoId,
		       CodigoProcedimento, NomeProcedimento, DenteRegiao,
		       Quantidade, ValorInformado, ValorPago, Status, Observacoes
		FROM ConvenioProcedimentos
		WHERE ConvenioProcedimentoId = @Id
	`,
		sql.Named("Id", id),
	).Scan(
		&p.ConvenioProcedimentoId,
		&p.ConvenioId,
		&p.ProcedimentoId,
		&p.CodigoProcedimento,
		&p.NomeProcedimento,
		&p.DenteRegiao,
		&p.Quantidade,
		&p.ValorInformado,
		&p.ValorPago,
		&p.Status,
		&p.Observacoes,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &p, err
}
