package fechamentolaboratorio

import "clinica-api/internal/models"

type FechamentoService interface {
    Criar(f *models.LaboratorioFechamento) error
    Editar(f *models.LaboratorioFechamento) error
    Listar() ([]models.LaboratorioFechamento, error)
    BuscarPorID(id int) (*models.LaboratorioFechamento, error)
    ListarPorLaboratorio(labId int) ([]models.LaboratorioFechamento, error)
    ListarPorData(data string) ([]models.LaboratorioFechamento, error)
    ListarPorMes(mes int, ano int) ([]models.LaboratorioFechamento, error)
    ListarPorPeriodo(inicio string, fim string) ([]models.LaboratorioFechamento, error)
    ListarPorLaboratorioMes(labId int, mes int, ano int) ([]models.LaboratorioFechamento, error)
    ListarPorPaciente(id int) ([]models.LaboratorioFechamento, error)
}

type fechamentoService struct {
    repo FechamentoRepository
}

func NovoFechamentoService(repo FechamentoRepository) FechamentoService {
    return &fechamentoService{repo: repo}
}

func (s *fechamentoService) Criar(f *models.LaboratorioFechamento) error {
    return s.repo.Criar(f)
}

func (s *fechamentoService) Editar(f *models.LaboratorioFechamento) error {
    return s.repo.Editar(f)
}

func (s *fechamentoService) Listar() ([]models.LaboratorioFechamento, error) {
    return s.repo.Listar()
}

func (s *fechamentoService) BuscarPorID(id int) (*models.LaboratorioFechamento, error) {
    return s.repo.BuscarPorID(id)
}

func (s *fechamentoService) ListarPorLaboratorio(labId int) ([]models.LaboratorioFechamento, error) {
    return s.repo.ListarPorLaboratorio(labId)
}

func (s *fechamentoService) ListarPorData(data string) ([]models.LaboratorioFechamento, error) {
    return s.repo.ListarPorData(data)
}
func (s *fechamentoService) ListarPorMes(mes int, ano int) ([]models.LaboratorioFechamento, error) {
    return s.repo.ListarPorMes(mes, ano)
}

func (s *fechamentoService) ListarPorPeriodo(inicio string, fim string) ([]models.LaboratorioFechamento, error) {
    return s.repo.ListarPorPeriodo(inicio, fim)
}

func (s *fechamentoService) ListarPorLaboratorioMes(labId int, mes int, ano int) ([]models.LaboratorioFechamento, error) {
    return s.repo.ListarPorLaboratorioMes(labId, mes, ano)
}

func (s *fechamentoService) ListarPorPaciente(id int) ([]models.LaboratorioFechamento, error) {
    return s.repo.ListarPorPaciente(id)
}
