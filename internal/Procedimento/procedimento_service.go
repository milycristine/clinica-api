package procedimentos

import "clinica-api/internal/models"

type ProcedimentoService interface {
    CriarProcedimento(p *models.Procedimento) error
    EditarProcedimento(p *models.Procedimento) error
    ListarProcedimentos() ([]models.Procedimento, error)
    BuscarProcedimentoPorID(id int) (*models.Procedimento, error)
    ListarProcedimentosAtivos() ([]models.Procedimento, error)

    AssociarProcedimentosAoLead(contatoId int, procedimentoIds []int) error
    ListarProcedimentosPorLead(contatoId int) ([]models.Procedimento, error)
}

type procedimentoService struct {
    repo ProcedimentoRepository
}

func NovoProcedimentoService(repo ProcedimentoRepository) ProcedimentoService {
    return &procedimentoService{repo: repo}
}

func (s *procedimentoService) CriarProcedimento(p *models.Procedimento) error {
    return s.repo.CriarProcedimento(p)
}

func (s *procedimentoService) EditarProcedimento(p *models.Procedimento) error {
    return s.repo.EditarProcedimento(p)
}

func (s *procedimentoService) ListarProcedimentos() ([]models.Procedimento, error) {
    return s.repo.ListarProcedimentos()
}

func (s *procedimentoService) BuscarProcedimentoPorID(id int) (*models.Procedimento, error) {
    return s.repo.BuscarProcedimentoPorID(id)
}

func (s *procedimentoService) ListarProcedimentosAtivos() ([]models.Procedimento, error) {
    return s.repo.ListarProcedimentosAtivos()
}

func (s *procedimentoService) AssociarProcedimentosAoLead(contatoId int, procedimentoIds []int) error {
    if err := s.repo.DeletarLeadProcedimentosPorLead(contatoId); err != nil {
        return err
    }

    for _, pid := range procedimentoIds {
        lp := models.LeadProcedimento{
            ContatoId:      contatoId,
            ProcedimentoId: pid,
        }
        if err := s.repo.CriarLeadProcedimento(&lp); err != nil {
            return err
        }
    }
    return nil
}

func (s *procedimentoService) ListarProcedimentosPorLead(contatoId int) ([]models.Procedimento, error) {
    return s.repo.ListarProcedimentosPorLead(contatoId)
}
