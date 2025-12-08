package captacao

import "clinica-api/internal/models"

type LeadHistoricoService interface {
    CriarHistorico(h *models.LeadHistorico) error
    ListarHistoricoPorLead(contatoId int) ([]models.LeadHistorico, error)
}

type leadHistoricoService struct {
    repo LeadHistoricoRepository
}

func NovoLeadHistoricoService(repo LeadHistoricoRepository) LeadHistoricoService {
    return &leadHistoricoService{repo: repo}
}

func (s *leadHistoricoService) CriarHistorico(h *models.LeadHistorico) error {
    return s.repo.CriarHistorico(h)
}

func (s *leadHistoricoService) ListarHistoricoPorLead(contatoId int) ([]models.LeadHistorico, error) {
    return s.repo.ListarHistoricoPorLead(contatoId)
}
