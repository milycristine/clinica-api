package despesa

import "clinica-api/internal/models"

type DespesaService interface {
	CriarDespesa(d *models.Despesa) error
	EditarDespesa(d *models.Despesa) error
	ListarDespesas() ([]models.Despesa, error)
	BuscarDespesaPorID(id int) (*models.Despesa, error)
	AtualizarStatus(id int, status string) error
}

type despesaService struct {
	repo DespesaRepository
}

func NovaDespesaService(repo DespesaRepository) DespesaService {
	return &despesaService{repo}
}

func (s *despesaService) CriarDespesa(d *models.Despesa) error {
	return s.repo.CriarDespesa(d)
}
func (s *despesaService) EditarDespesa(d *models.Despesa) error {
	return s.repo.EditarDespesa(d)
}
func (s *despesaService) ListarDespesas() ([]models.Despesa, error) {
	return s.repo.ListarDespesas()
}
func (s *despesaService) BuscarDespesaPorID(id int) (*models.Despesa, error) {
	return s.repo.BuscarDespesaPorID(id)
}
func (s *despesaService) AtualizarStatus(id int, status string) error {
	return s.repo.AtualizarStatus(id, status)
}
