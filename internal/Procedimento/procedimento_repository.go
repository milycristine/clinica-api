package procedimentos

import (
	database "clinica-api/database"
	"clinica-api/internal/models"
	"database/sql"
)

type ProcedimentoRepository interface {
	CriarProcedimento(p *models.Procedimento) error
	EditarProcedimento(p *models.Procedimento) error
	ListarProcedimentos() ([]models.Procedimento, error)
	BuscarProcedimentoPorID(id int) (*models.Procedimento, error)
	ListarProcedimentosAtivos() ([]models.Procedimento, error)

	CriarLeadProcedimento(lp *models.LeadProcedimento) error
	DeletarLeadProcedimentosPorLead(contatoId int) error
	ListarProcedimentosPorLead(contatoId int) ([]models.Procedimento, error)
}

type procedimentoRepository struct {
	db *sql.DB
}

func NovoProcedimentoRepository(conn *database.SQLStr) ProcedimentoRepository {
	return &procedimentoRepository{
		db: conn.DB(),
	}
}

func (r *procedimentoRepository) CriarProcedimento(p *models.Procedimento) error {
	query := `
        INSERT INTO Procedimentos (Nome)
        VALUES (@Nome)
    `
	_, err := r.db.Exec(query,
		sql.Named("Nome", p.Nome),
	)
	return err
}

func (r *procedimentoRepository) EditarProcedimento(p *models.Procedimento) error {
	query := `
        UPDATE Procedimentos
        SET Nome = @Nome, Ativo = @Ativo
        WHERE ProcedimentoId = @ProcedimentoId
    `
	_, err := r.db.Exec(query,
		sql.Named("Nome", p.Nome),
		sql.Named("Ativo", p.Ativo),
		sql.Named("ProcedimentoId", p.ProcedimentoId),
	)
	return err
}

func (r *procedimentoRepository) ListarProcedimentos() ([]models.Procedimento, error) {
	query := `SELECT ProcedimentoId, Nome, Ativo FROM Procedimentos ORDER BY Nome`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lista []models.Procedimento
	for rows.Next() {
		var p models.Procedimento
		var ativo bool
		if err := rows.Scan(&p.ProcedimentoId, &p.Nome, &ativo); err != nil {
			return nil, err
		}
		p.Ativo = ativo
		lista = append(lista, p)
	}
	return lista, nil
}

func (r *procedimentoRepository) BuscarProcedimentoPorID(id int) (*models.Procedimento, error) {
	query := `SELECT ProcedimentoId, Nome, Ativo FROM Procedimentos WHERE ProcedimentoId = @ProcedimentoId`
	row := r.db.QueryRow(query, sql.Named("ProcedimentoId", id))

	var p models.Procedimento
	var ativo bool
	if err := row.Scan(&p.ProcedimentoId, &p.Nome, &ativo); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	p.Ativo = ativo
	return &p, nil
}

func (r *procedimentoRepository) ListarProcedimentosAtivos() ([]models.Procedimento, error) {
	query := `SELECT ProcedimentoId, Nome, Ativo FROM Procedimentos WHERE Ativo = 1 ORDER BY Nome`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lista []models.Procedimento
	for rows.Next() {
		var p models.Procedimento
		var ativo bool
		if err := rows.Scan(&p.ProcedimentoId, &p.Nome, &ativo); err != nil {
			return nil, err
		}
		p.Ativo = ativo
		lista = append(lista, p)
	}
	return lista, nil
}


func (r *procedimentoRepository) CriarLeadProcedimento(lp *models.LeadProcedimento) error {
	query := `
        INSERT INTO LeadProcedimentos (ContatoId, ProcedimentoId)
        VALUES (@ContatoId, @ProcedimentoId)
    `
	_, err := r.db.Exec(query,
		sql.Named("ContatoId", lp.ContatoId),
		sql.Named("ProcedimentoId", lp.ProcedimentoId),
	)
	return err
}

func (r *procedimentoRepository) DeletarLeadProcedimentosPorLead(contatoId int) error {
	query := `DELETE FROM LeadProcedimentos WHERE ContatoId = @ContatoId`
	_, err := r.db.Exec(query, sql.Named("ContatoId", contatoId))
	return err
}

func (r *procedimentoRepository) ListarProcedimentosPorLead(contatoId int) ([]models.Procedimento, error) {
	query := `
        SELECT p.ProcedimentoId, p.Nome, p.Ativo
        FROM LeadProcedimentos lp
        JOIN Procedimentos p ON lp.ProcedimentoId = p.ProcedimentoId
        WHERE lp.ContatoId = @ContatoId
        ORDER BY p.Nome
    `
	rows, err := r.db.Query(query, sql.Named("ContatoId", contatoId))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lista []models.Procedimento
	for rows.Next() {
		var p models.Procedimento
		var ativo bool
		if err := rows.Scan(&p.ProcedimentoId, &p.Nome, &ativo); err != nil {
			return nil, err
		}
		p.Ativo = ativo
		lista = append(lista, p)
	}
	return lista, nil
}
