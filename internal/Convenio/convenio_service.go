package convenio

import (
	glosa "clinica-api/internal/Glosas"
	"clinica-api/internal/models"
	"fmt"
	"time"
)

type ConvenioService interface {
	CriarConvenio(c *models.Convenio) error
	EditarConvenio(c *models.Convenio) error
	ListarConvenios() ([]models.Convenio, error)
	BuscarConvenioPorID(id int) (*models.Convenio, error)

	AtualizarStatusProcedimento(
		convenioId int,
		convenioProcedimentoId int,
		status string,
		valorPago *float64,
		observacoes string,
		motivoOdontologico string,
		nomeContratado string,
	) error
}

type convenioService struct {
	convenioRepo     ConvenioRepository
	procedimentoRepo ConvenioProcedimentoRepository
	glosaService     glosa.GlosaDetalheService
}

func NovoConvenioService(
	convenioRepo ConvenioRepository,
	procedimentoRepo ConvenioProcedimentoRepository,
	glosaService glosa.GlosaDetalheService,
) ConvenioService {
	return &convenioService{
		convenioRepo:     convenioRepo,
		procedimentoRepo: procedimentoRepo,
		glosaService:     glosaService,
	}
}

func (s *convenioService) CriarConvenio(c *models.Convenio) error {
	_, err := s.convenioRepo.CriarConvenio(c)
	return err
}

func (s *convenioService) EditarConvenio(c *models.Convenio) error {
	return s.convenioRepo.EditarConvenio(c)
}

func (s *convenioService) ListarConvenios() ([]models.Convenio, error) {
	return s.convenioRepo.ListarConvenios()
}

func (s *convenioService) BuscarConvenioPorID(id int) (*models.Convenio, error) {
	c, err := s.convenioRepo.BuscarConvenioPorID(id)
	if err != nil || c == nil {
		return c, err
	}

	procs, err := s.procedimentoRepo.ListarPorConvenio(id)
	if err != nil {
		return nil, err
	}

	c.Procedimentos = procs
	return c, nil
}

func (s *convenioService) AtualizarStatusProcedimento(
	convenioId int,
	convenioProcedimentoId int,
	status string,
	valorPago *float64,
	observacoes string,
	motivoOdontologico string,
	nomeContratado string,
) error {

	if convenioId == 0 || convenioProcedimentoId == 0 {
		return fmt.Errorf("convenioId e convenioProcedimentoId são obrigatórios")
	}

	if status == "GLOSADO" {
		if motivoOdontologico == "" {
			return fmt.Errorf("motivo odontológico (MO) é obrigatório para glosa")
		}
		if nomeContratado == "" {
			return fmt.Errorf("nome do contratado é obrigatório para glosa")
		}
	}

	if err := s.procedimentoRepo.AtualizarStatus(
		convenioProcedimentoId,
		status,
		observacoes,
		valorPago,
	); err != nil {
		return err
	}

	if status != "GLOSADO" {
		return nil
	}

	convenio, err := s.convenioRepo.BuscarConvenioPorID(convenioId)
	if err != nil || convenio == nil {
		return fmt.Errorf("convênio não encontrado")
	}

	proc, err := s.procedimentoRepo.BuscarPorID(convenioProcedimentoId)
	if err != nil || proc == nil {
		return fmt.Errorf("procedimento do convênio não encontrado")
	}

	glosaDetalhe := models.GlosaDetalhe{
		ConvenioProcedimentoId: proc.ConvenioProcedimentoId,
		Guia:                   convenio.Guia,
		PacienteId:             convenio.PacienteId,
		UnidadeId:              convenio.UnidadeId,
		DataOcorrencia:         time.Now().Format("2006-01-02"),

		Mo:             motivoOdontologico,
		NomeContratado: nomeContratado,
		DenteRegiao:    proc.DenteRegiao,

		MotivoGlosa:   observacoes,
		StatusRecurso: "RECORRIVEL",

		Procedimentos: []models.GlosaProcedimento{
			{
				ProcedimentoId:     &proc.ProcedimentoId,
				CodigoProcedimento: proc.CodigoProcedimento,
				NomeProcedimento:   proc.NomeProcedimento,
				Quantidade:         proc.Quantidade,
				ValorInformado:     proc.ValorInformado,
				ValorGlosado:       proc.ValorInformado,
			},
		},
	}

	return s.glosaService.CriarGlosa(&glosaDetalhe)
}
