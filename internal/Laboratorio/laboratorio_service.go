package laboratorio

import "clinica-api/internal/models"

type LaboratorioService interface {
    CriarLaboratorio(l *models.Laboratorio) error
    EditarLaboratorio(l *models.Laboratorio) error
    ListarLaboratorios(filtro models.LaboratorioFiltro) ([]models.Laboratorio, int, error)
    BuscarLaboratorioPorID(id int) (*models.Laboratorio, error)

}

type laboratorioService struct {
    repo LaboratorioRepository
}

func NovoLaboratorioService(repo LaboratorioRepository) LaboratorioService {
    return &laboratorioService{repo: repo}
}

func (s *laboratorioService) CriarLaboratorio(l *models.Laboratorio) error {
    return s.repo.CriarLaboratorio(l)
}

func (s *laboratorioService) EditarLaboratorio(l *models.Laboratorio) error {
    return s.repo.EditarLaboratorio(l)
}

func (s *laboratorioService) ListarLaboratorios(filtro models.LaboratorioFiltro) ([]models.Laboratorio, int, error) {
    return s.repo.ListarLaboratorios(filtro)
}

func (s *laboratorioService) BuscarLaboratorioPorID(id int) (*models.Laboratorio, error) {
    return s.repo.BuscarLaboratorioPorID(id)
}
