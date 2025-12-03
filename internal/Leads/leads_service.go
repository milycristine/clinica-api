package captacao

import "clinica-api/internal/models"

type ContatoService interface {
	CriarContato(c *models.Leads) error
	EditarContato(c *models.Leads) error
	ListarContatos() ([]models.Leads, error)
	BuscarContatoPorID(id int) (*models.Leads, error)
	AtualizarStatus(id int, status string) error
}

type contatoService struct {
	repo ContatoRepository
}

func NovoContatoService(repo ContatoRepository) ContatoService {
	return &contatoService{repo: repo}
}

func (s *contatoService) CriarContato(c *models.Leads) error {
	return s.repo.CriarContato(c)
}

func (s *contatoService) EditarContato(c *models.Leads) error {
	return s.repo.EditarContato(c)
}

func (s *contatoService) ListarContatos() ([]models.Leads, error) {
	return s.repo.ListarContatos()
}

func (s *contatoService) BuscarContatoPorID(id int) (*models.Leads, error) {
	return s.repo.BuscarContatoPorID(id)
}

func (s *contatoService) AtualizarStatus(id int, status string) error {
	return s.repo.AtualizarStatus(id, status)
}
