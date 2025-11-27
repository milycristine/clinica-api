package unidade

import (
    "clinica-api/internal/models"
)

type UnidadeService interface {
    CriarUnidade(u *models.Unidade) error
    EditarUnidade(u *models.Unidade) error
    ListarUnidades() ([]models.Unidade, error)
    BuscarUnidadePorID(id int) (*models.Unidade, error)
    AtualizarStatus(id int, status bool) error
}

type unidadeService struct {
    repo UnidadeRepository
}

func NovoUnidadeService(repo UnidadeRepository) UnidadeService {
    return &unidadeService{repo: repo}
}

func (s *unidadeService) CriarUnidade(u *models.Unidade) error {
    return s.repo.CriarUnidade(u)
}

func (s *unidadeService) EditarUnidade(u *models.Unidade) error {
    return s.repo.EditarUnidade(u)
}

func (s *unidadeService) ListarUnidades() ([]models.Unidade, error) {
    return s.repo.ListarUnidades()
}

func (s *unidadeService) BuscarUnidadePorID(id int) (*models.Unidade, error) {
    return s.repo.BuscarUnidadePorID(id)
}

func (s *unidadeService) AtualizarStatus(id int, status bool) error {
    novoStatus := 0
    if status {
        novoStatus = 1
    }
    return s.repo.AtualizarStatus(id, novoStatus)
}
