package glosa

import "clinica-api/internal/models"

type GlosaMensalService interface {
    CriarGlosaMensal(gm *models.GlosaMensal) error
    EditarGlosaMensal(gm *models.GlosaMensal) error
    ListarGlosasMensais(f models.GlosaMensalFiltro) ([]models.GlosaMensal, error)
    BuscarGlosaMensalPorID(id int) (*models.GlosaMensal, error)
}

type glosaMensalService struct {
    repo GlosaMensalRepository
}

func NovoGlosaMensalService(repo GlosaMensalRepository) GlosaMensalService {
    return &glosaMensalService{repo: repo}
}

func (s *glosaMensalService) CriarGlosaMensal(gm *models.GlosaMensal) error {
    return s.repo.CriarGlosaMensal(gm)
}

func (s *glosaMensalService) EditarGlosaMensal(gm *models.GlosaMensal) error {
    return s.repo.EditarGlosaMensal(gm)
}

func (s *glosaMensalService) ListarGlosasMensais(f models.GlosaMensalFiltro) ([]models.GlosaMensal, error) {
    return s.repo.ListarGlosasMensais(f)
}

func (s *glosaMensalService) BuscarGlosaMensalPorID(id int) (*models.GlosaMensal, error) {
    return s.repo.BuscarGlosaMensalPorID(id)
}
