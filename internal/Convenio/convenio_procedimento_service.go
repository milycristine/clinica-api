package convenio

import (
	"clinica-api/internal/models"
	"fmt"
)

type ConvenioProcedimentoService interface {
	Criar(p *models.ConvenioProcedimento) error
	ListarPorConvenio(convenioId int) ([]models.ConvenioProcedimento, error)
	AtualizarStatus(
		id int,
		status string,
		valorPago *float64,
		obs string,
	) error
}

type convenioProcedimentoService struct {
	repo ConvenioProcedimentoRepository
}

func NovoConvenioProcedimentoService(
	repo ConvenioProcedimentoRepository,
) ConvenioProcedimentoService {
	return &convenioProcedimentoService{
		repo: repo,
	}
}

func (s *convenioProcedimentoService) Criar(p *models.ConvenioProcedimento) error {
	if p.ConvenioId == 0 {
		return fmt.Errorf("convenioId é obrigatório")
	}

	if p.ProcedimentoId == 0 {
		return fmt.Errorf("procedimentoId é obrigatório")
	}

	if p.Quantidade <= 0 {
		p.Quantidade = 1
	}

	return s.repo.Criar(p)
}

func (s *convenioProcedimentoService) ListarPorConvenio(convenioId int) ([]models.ConvenioProcedimento, error) {
	return s.repo.ListarPorConvenio(convenioId)
}
func (s *convenioProcedimentoService) AtualizarStatus(
	id int,
	status string,
	valorPago *float64,
	obs string,
) error {

	switch status {
	case "AUTORIZADO":
		if valorPago == nil {
			return fmt.Errorf("valorPago é obrigatório")
		}

	case "GLOSADO":
		if obs == "" {
			return fmt.Errorf("observação obrigatória para glosa")
		}
	}

	return s.repo.AtualizarStatus(id, status, obs, valorPago)
}
