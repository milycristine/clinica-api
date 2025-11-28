package paciente

import "clinica-api/internal/models"

type PacienteService interface {
	CriarPaciente(p *models.Paciente) error
	EditarPaciente(p *models.Paciente) error
	ListarPacientes() ([]models.Paciente, error)
	BuscarPacientePorID(id int) (*models.Paciente, error)
	ListarPorUnidade(unidadeId int) ([]models.Paciente, error) 
}

type pacienteService struct {
	repo PacienteRepository
}

func NovoPacienteService(repo PacienteRepository) PacienteService {
	return &pacienteService{repo: repo}
}

func (s *pacienteService) CriarPaciente(p *models.Paciente) error {
	return s.repo.CriarPaciente(p)
}

func (s *pacienteService) EditarPaciente(p *models.Paciente) error {
	return s.repo.EditarPaciente(p)
}

func (s *pacienteService) ListarPacientes() ([]models.Paciente, error) {
	return s.repo.ListarPacientes()
}

func (s *pacienteService) BuscarPacientePorID(id int) (*models.Paciente, error) {
	return s.repo.BuscarPacientePorID(id)
}
func (s *pacienteService) ListarPorUnidade(unidadeId int) ([]models.Paciente, error)  {
	return s.repo.ListarPorUnidade(unidadeId)
}
