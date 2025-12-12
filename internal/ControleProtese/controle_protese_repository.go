package controleprotese

import (
	"clinica-api/database"
	"clinica-api/internal/models"
	"database/sql"
	"fmt"
	"strings"
)

type ControleProteseRepository interface {
	Criar(c *models.ControleProtese) error
	ExisteDuplicacao(c *models.ControleProtese) (bool, error)
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

type controleRepository struct {
	db *sql.DB
}

func NovoControleRepository(conn *database.SQLStr) ControleProteseRepository {
	return &controleRepository{db: conn.DB()}
}
func (r *controleRepository) ExisteDuplicacao(c *models.ControleProtese) (bool, error) {
	query := `
		SELECT COUNT(1)
		FROM ControleProteses WITH (NOLOCK)
		WHERE PacienteId=@PacienteId
		AND Produto=@Produto
		AND LaboratorioId=@LaboratorioId
		AND CAST(DataEnvio AS DATE) = CAST(@DataEnvio AS DATE)
	`

	var qtd int
	err := r.db.QueryRow(
		query,
		sql.Named("PacienteId", c.PacienteId),
		sql.Named("Produto", c.Produto),
		sql.Named("LaboratorioId", c.LaboratorioId),
		sql.Named("DataEnvio", c.DataEnvio),
	).Scan(&qtd)

	if err != nil {
		return false, err
	}

	return qtd > 0, nil
}

func (r *controleRepository) Criar(c *models.ControleProtese) error {

	duplicado, err := r.ExisteDuplicacao(c)
	if err != nil {
		return err
	}
	if duplicado {
		return fmt.Errorf("já existe uma prótese igual cadastrada hoje para este paciente e laboratório")
	}

	query := `
		INSERT INTO ControleProteses
		(Ficha, PacienteId, LaboratorioId, Produto, EtapaAtual, TrabalhoFinal,
		DataEnvio, DataEntregaPrevista, DataRecebimento, Status, Observacoes,
		CustoUnitario, Quantidade)
		VALUES (
			@Ficha, @PacienteId, @LaboratorioId, @Produto, @EtapaAtual, @TrabalhoFinal,
			@DataEnvio, @DataEntregaPrevista, @DataRecebimento, @Status, @Observacoes,
			@CustoUnitario, @Quantidade
		)
	`

	_, err = r.db.Exec(query,
		sql.Named("Ficha", c.Ficha),
		sql.Named("PacienteId", c.PacienteId),
		sql.Named("LaboratorioId", c.LaboratorioId),
		sql.Named("Produto", c.Produto),
		sql.Named("EtapaAtual", c.EtapaAtual),
		sql.Named("TrabalhoFinal", c.TrabalhoFinal),
		sql.Named("DataEnvio", c.DataEnvio),
		sql.Named("DataEntregaPrevista", c.DataEntregaPrevista),
		sql.Named("DataRecebimento", c.DataRecebimento),
		sql.Named("Status", c.Status),
		sql.Named("Observacoes", c.Observacoes),
		sql.Named("CustoUnitario", c.CustoUnitario),
		sql.Named("Quantidade", c.Quantidade),
	)

	if err != nil {
		return fmt.Errorf("erro ao criar: %w", err)
	}

	return nil
}

func (r *controleRepository) Editar(c *models.ControleProtese) error {
	if c.ProteseId == 0 {
		return fmt.Errorf("ID é obrigatório")
	}

	var updates []string
	var args []any

	if c.Ficha != "" {
		updates = append(updates, "Ficha=@Ficha")
		args = append(args, sql.Named("Ficha", c.Ficha))
	}
	if c.PacienteId != 0 {
		updates = append(updates, "PacienteId=@PacienteId")
		args = append(args, sql.Named("PacienteId", c.PacienteId))
	}
	if c.LaboratorioId != 0 {
		updates = append(updates, "LaboratorioId=@LaboratorioId")
		args = append(args, sql.Named("LaboratorioId", c.LaboratorioId))
	}
	if c.Produto != "" {
		updates = append(updates, "Produto=@Produto")
		args = append(args, sql.Named("Produto", c.Produto))
	}
	if c.EtapaAtual != "" {
		updates = append(updates, "EtapaAtual=@EtapaAtual")
		args = append(args, sql.Named("EtapaAtual", c.EtapaAtual))
	}
	if c.TrabalhoFinal != "" {
		updates = append(updates, "TrabalhoFinal=@TrabalhoFinal")
		args = append(args, sql.Named("TrabalhoFinal", c.TrabalhoFinal))
	}
	if c.DataEnvio != "" {
		updates = append(updates, "DataEnvio=@DataEnvio")
		args = append(args, sql.Named("DataEnvio", c.DataEnvio))
	}
	if c.DataEntregaPrevista != "" {
		updates = append(updates, "DataEntregaPrevista=@DataEntregaPrevista")
		args = append(args, sql.Named("DataEntregaPrevista", c.DataEntregaPrevista))
	}
	if c.DataRecebimento != "" {
		updates = append(updates, "DataRecebimento=@DataRecebimento")
		args = append(args, sql.Named("DataRecebimento", c.DataRecebimento))
	}
	if c.Status != "" {
		updates = append(updates, "Status=@Status")
		args = append(args, sql.Named("Status", c.Status))
	}
	if c.Observacoes != "" {
		updates = append(updates, "Observacoes=@Observacoes")
		args = append(args, sql.Named("Observacoes", c.Observacoes))
	}
	if c.CustoUnitario != 0 {
		updates = append(updates, "CustoUnitario=@CustoUnitario")
		args = append(args, sql.Named("CustoUnitario", c.CustoUnitario))
	}
	if c.Quantidade != 0 {
		updates = append(updates, "Quantidade=@Quantidade")
		args = append(args, sql.Named("Quantidade", c.Quantidade))
	}

	if len(updates) == 0 {
		return fmt.Errorf("nenhum campo para atualizar")
	}

	args = append(args, sql.Named("ProteseId", c.ProteseId))

	query := fmt.Sprintf(`
        UPDATE ControleProteses
        SET %s
        WHERE ProteseId=@ProteseId
    `, strings.Join(updates, ", "))

	_, err := r.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("erro ao editar: %w", err)
	}

	return nil
}

func (r *controleRepository) Listar(page int, limit int) ([]models.ControleProtese, error) {

	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * limit

	query := `
		SELECT ProteseId, Ficha, PacienteId, LaboratorioId, Produto, EtapaAtual,
		TrabalhoFinal, DataEnvio, DataEntregaPrevista, DataRecebimento,
		Status, Observacoes, CustoUnitario, Quantidade, CustoTotal
		FROM ControleProteses WITH (NOLOCK)
		ORDER BY ProteseId DESC
		OFFSET @Offset ROWS
		FETCH NEXT @Limit ROWS ONLY
	`

	rows, err := r.db.Query(query,
		sql.Named("Offset", offset),
		sql.Named("Limit", limit),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	lista := []models.ControleProtese{}

	for rows.Next() {
		var c models.ControleProtese
		rows.Scan(
			&c.ProteseId, &c.Ficha, &c.PacienteId, &c.LaboratorioId,
			&c.Produto, &c.EtapaAtual, &c.TrabalhoFinal, &c.DataEnvio,
			&c.DataEntregaPrevista, &c.DataRecebimento, &c.Status,
			&c.Observacoes, &c.CustoUnitario, &c.Quantidade, &c.CustoTotal,
		)
		lista = append(lista, c)
	}

	return lista, nil
}

func (r *controleRepository) BuscarPorID(id int) (*models.ControleProtese, error) {
	var c models.ControleProtese

	query := `
        SELECT ProteseId, Ficha, PacienteId, LaboratorioId, Produto, EtapaAtual,
               TrabalhoFinal, DataEnvio, DataEntregaPrevista, DataRecebimento,
               Status, Observacoes, CustoUnitario, Quantidade, CustoTotal
        FROM ControleProteses WITH (NOLOCK)
        WHERE ProteseId=@ProteseId
    `

	err := r.db.QueryRow(query, sql.Named("ProteseId", id)).Scan(
		&c.ProteseId, &c.Ficha, &c.PacienteId, &c.LaboratorioId, &c.Produto,
		&c.EtapaAtual, &c.TrabalhoFinal, &c.DataEnvio, &c.DataEntregaPrevista,
		&c.DataRecebimento, &c.Status, &c.Observacoes, &c.CustoUnitario,
		&c.Quantidade, &c.CustoTotal,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *controleRepository) AlterarStatus(id int, status string) error {
	query := `
        UPDATE ControleProteses
        SET Status=@Status
        WHERE ProteseId=@ProteseId
    `

	_, err := r.db.Exec(query,
		sql.Named("Status", status),
		sql.Named("ProteseId", id),
	)

	if err != nil {
		return fmt.Errorf("erro ao alterar status: %w", err)
	}

	return nil
}

func (r *controleRepository) ListarFiltrado(
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

	filtros := []string{}
	params := []any{}

	normalizeDate := func(s string) string {
		if len(s) >= 10 {
			return s[:10] 
		}
		return s
	}
	data = normalizeDate(data)
	dataInicio = normalizeDate(dataInicio)
	dataFim = normalizeDate(dataFim)

	
	if data != "" {
		filtros = append(filtros, "CAST(DataEnvio AS DATE) = CAST(@Data AS DATE)")
		params = append(params, sql.Named("Data", data))
	}


	if dataInicio != "" && dataFim != "" {
		filtros = append(filtros, "CAST(DataEnvio AS DATE) BETWEEN CAST(@DataInicio AS DATE) AND CAST(@DataFim AS DATE)")
		params = append(params, sql.Named("DataInicio", dataInicio))
		params = append(params, sql.Named("DataFim", dataFim))
	} else if dataInicio != "" {
		filtros = append(filtros, "CAST(DataEnvio AS DATE) >= CAST(@DataInicio AS DATE)")
		params = append(params, sql.Named("DataInicio", dataInicio))
	} else if dataFim != "" {
		filtros = append(filtros, "CAST(DataEnvio AS DATE) <= CAST(@DataFim AS DATE)")
		params = append(params, sql.Named("DataFim", dataFim))
	}

	
	if status != "" {
		filtros = append(filtros, "Status LIKE @StatusFiltro")
		params = append(params, sql.Named("StatusFiltro", "%"+status+"%"))
	}


	if laboratorioId != 0 {
		filtros = append(filtros, "LaboratorioId=@LaboratorioId")
		params = append(params, sql.Named("LaboratorioId", laboratorioId))
	}


	if pacienteId != 0 {
		filtros = append(filtros, "PacienteId=@PacienteId")
		params = append(params, sql.Named("PacienteId", pacienteId))
	}

	if produto != "" {
		filtros = append(filtros, "Produto LIKE @Produto")
		params = append(params, sql.Named("Produto", "%"+produto+"%"))
	}

	if etapa != "" {
		filtros = append(filtros, "EtapaAtual LIKE @EtapaAtual")
		params = append(params, sql.Named("EtapaAtual", "%"+etapa+"%"))
	}


	if limit <= 0 {
		limit = 999999 
	}

	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * limit

	query := `
		SELECT ProteseId, Ficha, PacienteId, LaboratorioId, Produto, EtapaAtual,
		TrabalhoFinal, DataEnvio, DataEntregaPrevista, DataRecebimento,
		Status, Observacoes, CustoUnitario, Quantidade, CustoTotal
		FROM ControleProteses WITH (NOLOCK)
	`

	if len(filtros) > 0 {
		query += " WHERE " + strings.Join(filtros, " AND ")
	}

	query += `
		ORDER BY ProteseId DESC
		OFFSET @Offset ROWS
		FETCH NEXT @Limit ROWS ONLY
	`

	params = append(params, sql.Named("Offset", offset))
	params = append(params, sql.Named("Limit", limit))

	rows, err := r.db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	lista := []models.ControleProtese{}

	for rows.Next() {
		var c models.ControleProtese
		rows.Scan(
			&c.ProteseId, &c.Ficha, &c.PacienteId, &c.LaboratorioId, &c.Produto,
			&c.EtapaAtual, &c.TrabalhoFinal, &c.DataEnvio, &c.DataEntregaPrevista,
			&c.DataRecebimento, &c.Status, &c.Observacoes,
			&c.CustoUnitario, &c.Quantidade, &c.CustoTotal,
		)
		lista = append(lista, c)
	}

	return lista, nil
}
