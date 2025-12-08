package captacao

import (
	procedimentoRepo "clinica-api/internal/Procedimento"
	"clinica-api/internal/models"
)

type ContatoService interface {
	CriarContato(c *models.Leads) error
	EditarContato(c *models.Leads) error
	ListarContatos() ([]models.Leads, error)
	BuscarContatoPorID(id int) (*models.Leads, error)
	AtualizarStatus(id int, status string, funcionarioId *int) error
}

type contatoService struct {
	repo             ContatoRepository
	historicoRepo    LeadHistoricoRepository
	procedimentoRepo procedimentoRepo.ProcedimentoRepository
}

func NovoContatoService(
	repo ContatoRepository,
	historicoRepo LeadHistoricoRepository,
	procRepo procedimentoRepo.ProcedimentoRepository,
) ContatoService {

	return &contatoService{
		repo:             repo,
		historicoRepo:    historicoRepo,
		procedimentoRepo: procRepo,
	}
}

func (s *contatoService) CriarContato(c *models.Leads) error {

	if err := s.repo.CriarContato(c); err != nil {
		return err
	}

	for _, procID := range c.Procedimentos {
		lp := models.LeadProcedimento{
			ContatoId:      c.ContatoId,
			ProcedimentoId: procID,
		}

		if err := s.procedimentoRepo.CriarLeadProcedimento(&lp); err != nil {
			return err
		}
	}

	return nil
}

func (s *contatoService) EditarContato(c *models.Leads) error {
	return s.repo.EditarContato(c)
}

func (s *contatoService) ListarContatos() ([]models.Leads, error) {
	contatos, err := s.repo.ListarContatos()
	if err != nil {
		return nil, err
	}

	for i := range contatos {
		procs, _ := s.procedimentoRepo.ListarProcedimentosPorLead(contatos[i].ContatoId)
		var ids []int
		for _, p := range procs {
			ids = append(ids, p.ProcedimentoId)
		}
		contatos[i].Procedimentos = ids
	}

	return contatos, nil
}

func (s *contatoService) BuscarContatoPorID(id int) (*models.Leads, error) {
	contato, err := s.repo.BuscarContatoPorID(id)
	if err != nil || contato == nil {
		return contato, err
	}

	procs, _ := s.procedimentoRepo.ListarProcedimentosPorLead(contato.ContatoId)
	var ids []int
	for _, p := range procs {
		ids = append(ids, p.ProcedimentoId)
	}
	contato.Procedimentos = ids

	return contato, nil
}

func (s *contatoService) AtualizarStatus(id int, status string, funcionarioId *int) error {

	if err := s.repo.AtualizarStatus(id, status); err != nil {
		return err
	}

	h := models.LeadHistorico{
		ContatoId:     id,
		Descricao:     "Status alterado para: " + status,
		FuncionarioId: funcionarioId,
	}

	return s.historicoRepo.CriarHistorico(&h)
}
