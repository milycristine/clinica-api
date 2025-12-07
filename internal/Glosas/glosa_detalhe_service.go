package glosa

import "clinica-api/internal/models"

type GlosaDetalheService interface {
    CriarGlosa(g *models.GlosaDetalhe) error
    EditarGlosa(g *models.GlosaDetalhe) error
    ListarGlosas() ([]models.GlosaDetalhe, error)
    BuscarPorId(id int) (*models.GlosaDetalhe, error)
}

type glosaDetalheService struct {
    repo GlosaDetalheRepository
}

func NovoGlosaDetalheService(repo GlosaDetalheRepository) GlosaDetalheService {
    return &glosaDetalheService{repo: repo}
}

func (s *glosaDetalheService) CriarGlosa(g *models.GlosaDetalhe) error {
    return s.repo.CriarGlosa(g)
}

func (s *glosaDetalheService) EditarGlosa(g *models.GlosaDetalhe) error {
    return s.repo.EditarGlosa(g)
}

func (s *glosaDetalheService) ListarGlosas() ([]models.GlosaDetalhe, error) {
    return s.repo.ListarGlosas()
}

func (s *glosaDetalheService) BuscarPorId(id int) (*models.GlosaDetalhe, error) {
    return s.repo.BuscarPorId(id)
}
