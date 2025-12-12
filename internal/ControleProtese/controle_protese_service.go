package controleprotese

import "clinica-api/internal/models"

type ControleProteseService interface {
	Criar(c *models.ControleProtese) error
	Editar(c *models.ControleProtese) error
	Listar(page int, limit int) ([]models.ControleProtese, error)
	BuscarPorID(id int) (*models.ControleProtese, error)
	AlterarStatus(id int, status string) error
	ListarFiltrado(
		data string,
		dataInicio string,
		dataFim string,
		status string,
		laboratorioId int,
		pacienteId int,
		produto string,
		etapa string,
		page int,
		limit int,
	) ([]models.ControleProtese, error)
}

type controleService struct {
	repo ControleProteseRepository
}

func NovoControleService(repo ControleProteseRepository) ControleProteseService {
	return &controleService{repo: repo}
}

func (s *controleService) Criar(c *models.ControleProtese) error {
	if c.Status == "" {
		c.Status = "Enviado"
	}
	return s.repo.Criar(c)
}

func (s *controleService) Editar(c *models.ControleProtese) error {
	return s.repo.Editar(c)
}

func (s *controleService) Listar(page int, limit int) ([]models.ControleProtese, error) {
	return s.repo.Listar(page, limit)
}

func (s *controleService) BuscarPorID(id int) (*models.ControleProtese, error) {
	return s.repo.BuscarPorID(id)
}

func (s *controleService) AlterarStatus(id int, status string) error {
	return s.repo.AlterarStatus(id, status)
}

func (s *controleService) ListarFiltrado(
	data string,
	dataInicio string,
	dataFim string,
	status string,
	laboratorioId int,
	pacienteId int,
	produto string,
	etapa string,
	page int,
	limit int,
) ([]models.ControleProtese, error) {

	return s.repo.ListarFiltrado(
		data, dataInicio, dataFim, status,
		laboratorioId, pacienteId, produto, etapa,
		page, limit,
	)
}
