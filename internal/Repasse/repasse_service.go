package repasse

import "clinica-api/internal/models"

type RepasseService interface {
	Criar(r *models.Repasse) error
	Listar() ([]models.Repasse, error)
	BuscarPorID(id int) (*models.Repasse, error)
	Pagar(id int) error
}

type repasseService struct {
	repo RepasseRepository
}

func NovoRepasseService(repo RepasseRepository) RepasseService {
	return &repasseService{repo: repo}
}

func (s *repasseService) Criar(r *models.Repasse) error {
	return s.repo.Criar(r)
}

func (s *repasseService) Listar() ([]models.Repasse, error) {
	return s.repo.Listar()
}

func (s *repasseService) BuscarPorID(id int) (*models.Repasse, error) {
	return s.repo.BuscarPorID(id)
}

func (s *repasseService) Pagar(id int) error {
	return s.repo.AtualizarStatus(id, "PAGO")
}
