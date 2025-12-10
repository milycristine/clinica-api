package glosa

import (
	"clinica-api/internal/models"
	"fmt"
	"time"
)

type GlosaDetalheService interface {
	CriarGlosa(g *models.GlosaDetalhe) error
	EditarGlosa(g *models.GlosaDetalhe) error
	ListarGlosas() ([]models.GlosaDetalhe, error)
	BuscarPorId(id int) (*models.GlosaDetalhe, error)
}

type glosaDetalheService struct {
	repo       GlosaDetalheRepository
	mensalRepo GlosaMensalRepository
}

func NovoGlosaDetalheService(repo GlosaDetalheRepository, mensalRepo GlosaMensalRepository) GlosaDetalheService {
	return &glosaDetalheService{
		repo:       repo,
		mensalRepo: mensalRepo,
	}
}

func (s *glosaDetalheService) CriarGlosa(g *models.GlosaDetalhe) error {
	data, err := time.Parse("2006-01-02", g.DataOcorrencia)
	if err != nil {
		return fmt.Errorf("data inválida: %v", err)
	}

	mes := int(data.Month())
	ano := data.Year()

	if g.UnidadeId == 0 {
		return fmt.Errorf("UnidadeId é obrigatório")
	}

	if len(g.Procedimentos) == 0 {
		return fmt.Errorf("é obrigatório informar ao menos um procedimento")
	}

	existe, err := s.repo.ExisteGlosa(g)
	if err != nil {
		return fmt.Errorf("erro ao verificar duplicidade: %v", err)
	}
	if existe {
		return fmt.Errorf("glosa detalhe já existe para essa guia, paciente, data e unidade")
	}

	mensalId, err := s.mensalRepo.BuscarOuCriarMensal(mes, ano, g.UnidadeId)
	if err != nil {
		return err
	}

	g.GlosasMensalId = &mensalId

	if err := s.repo.CriarGlosa(g); err != nil {
		return err
	}

	return s.mensalRepo.RecalcularGlosaMensal(mensalId)
}

func (s *glosaDetalheService) EditarGlosa(g *models.GlosaDetalhe) error {
	if len(g.Procedimentos) == 0 {
		return fmt.Errorf("é obrigatório informar ao menos um procedimento")
	}

	if err := s.repo.EditarGlosa(g); err != nil {
		return err
	}

	if g.GlosasMensalId != nil {
		return s.mensalRepo.RecalcularGlosaMensal(*g.GlosasMensalId)
	}

	return nil
}

func (s *glosaDetalheService) ListarGlosas() ([]models.GlosaDetalhe, error) {
	return s.repo.ListarGlosas()
}

func (s *glosaDetalheService) BuscarPorId(id int) (*models.GlosaDetalhe, error) {
	return s.repo.BuscarPorId(id)
}
