package ajustecaixa

import "clinica-api/internal/models"

type AjusteCaixaService interface {
	Criar(a *models.AjusteCaixa) error
	Listar() ([]models.AjusteCaixa, error)
}

type ajusteCaixaService struct {
	repo AjusteCaixaRepository
}

func NovoAjusteCaixaService(repo AjusteCaixaRepository) AjusteCaixaService {
	return &ajusteCaixaService{repo: repo}
}

func (s *ajusteCaixaService) Criar(a *models.AjusteCaixa) error {
	return s.repo.Criar(a)
}

func (s *ajusteCaixaService) Listar() ([]models.AjusteCaixa, error) {
	return s.repo.Listar()
}
