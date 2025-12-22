package receita

import "clinica-api/internal/models"

type ReceitaService interface {
	CriarReceita(r *models.Receita) error
	EditarReceita(r *models.Receita) error
	ListarReceitas() ([]models.Receita, error)
	BuscarReceitaPorID(id int) (*models.Receita, error)
	AtualizarStatus(id int, status string) error
}

type receitaService struct {
	repo ReceitaRepository
}

func NovoReceitaService(repo ReceitaRepository) ReceitaService {
	return &receitaService{repo: repo}
}

func (s *receitaService) CriarReceita(r *models.Receita) error {
	if r.Status == "" {
		r.Status = "REALIZADO"
	}
	return s.repo.CriarReceita(r)
}

func (s *receitaService) EditarReceita(r *models.Receita) error {
	return s.repo.EditarReceita(r)
}

func (s *receitaService) ListarReceitas() ([]models.Receita, error) {
	return s.repo.ListarReceitas()
}

func (s *receitaService) BuscarReceitaPorID(id int) (*models.Receita, error) {
	return s.repo.BuscarReceitaPorID(id)
}

func (s *receitaService) AtualizarStatus(id int, status string) error {
	return s.repo.AtualizarStatus(id, status)
}
