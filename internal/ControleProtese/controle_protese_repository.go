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
	Editar(c *models.ControleProtese) error
	Listar() ([]models.ControleProtese, error)
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
	) ([]models.ControleProtese, error)
}

type controleRepository struct {
	db *sql.DB
}

func NovoControleRepository(conn *database.SQLStr) ControleProteseRepository {
	return &controleRepository{
		db: conn.DB(),
	}
}

func (r *controleRepository) Criar(c *models.ControleProtese) error {
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
	_, err := r.db.Exec(query,
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
		return fmt.Errorf("erro ao criar controle de prótese: %w", err)
	}

	return nil
}

func (r *controleRepository) Editar(c *models.ControleProtese) error {
	if c.ProteseId == 0 {
		return fmt.Errorf("ID da prótese é obrigatório")
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
        WHERE ProteseId = @ProteseId
    `, strings.Join(updates, ", "))

	_, err := r.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("erro ao editar prótese: %w", err)
	}

	return nil
}

func (r *controleRepository) Listar() ([]models.ControleProtese, error) {
	lista := []models.ControleProtese{}

	query := `
        SELECT ProteseId, Ficha, PacienteId, LaboratorioId, Produto, EtapaAtual,
               TrabalhoFinal, DataEnvio, DataEntregaPrevista, DataRecebimento,
               Status, Observacoes, CustoUnitario, Quantidade, CustoTotal
        FROM ControleProteses WITH (NOLOCK)
        ORDER BY ProteseId DESC
    `

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c models.ControleProtese
		err := rows.Scan(
			&c.ProteseId, &c.Ficha, &c.PacienteId, &c.LaboratorioId, &c.Produto,
			&c.EtapaAtual, &c.TrabalhoFinal, &c.DataEnvio, &c.DataEntregaPrevista,
			&c.DataRecebimento, &c.Status, &c.Observacoes, &c.CustoUnitario,
			&c.Quantidade, &c.CustoTotal,
		)
		if err != nil {
			return nil, err
		}
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
        FROM ControleProteses
        WHERE ProteseId = @ProteseId
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
) ([]models.ControleProtese, error) {

	filtros := []string{}
	params := []any{}

	// FILTRO POR DATA ESPECÍFICA
	if data != "" {
		filtros = append(filtros, "CONVERT(date, DataEnvio) = @Data")
		params = append(params, sql.Named("Data", data))
	}

	// FILTRO POR PERÍODO
	if dataInicio != "" && dataFim != "" {
		filtros = append(filtros, "CONVERT(date, DataEnvio) BETWEEN @DataInicio AND @DataFim")
		params = append(params,
			sql.Named("DataInicio", dataInicio),
			sql.Named("DataFim", dataFim),
		)
	}

	// STATUS
	if status != "" {
		filtros = append(filtros, "Status = @Status")
		params = append(params, sql.Named("Status", status))
	}

	// LABORATÓRIO
	if laboratorioId != 0 {
		filtros = append(filtros, "LaboratorioId = @LaboratorioId")
		params = append(params, sql.Named("LaboratorioId", laboratorioId))
	}

	// PACIENTE
	if pacienteId != 0 {
		filtros = append(filtros, "PacienteId = @PacienteId")
		params = append(params, sql.Named("PacienteId", pacienteId))
	}

	// PRODUTO
	if produto != "" {
		filtros = append(filtros, "Produto LIKE @Produto")
		params = append(params, sql.Named("Produto", "%"+produto+"%"))
	}

	// ETAPA ATUAL
	if etapa != "" {
		filtros = append(filtros, "EtapaAtual LIKE @EtapaAtual")
		params = append(params, sql.Named("EtapaAtual", "%"+etapa+"%"))
	}

	query := `
        SELECT ProteseId, Ficha, PacienteId, LaboratorioId, Produto, EtapaAtual,
               TrabalhoFinal, DataEnvio, DataEntregaPrevista, DataRecebimento,
               Status, Observacoes, CustoUnitario, Quantidade, CustoTotal
        FROM ControleProteses WITH (NOLOCK)
    `

	if len(filtros) > 0 {
		query += " WHERE " + strings.Join(filtros, " AND ")
	}

	query += " ORDER BY ProteseId DESC"

	rows, err := r.db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	lista := []models.ControleProtese{}

	for rows.Next() {
		var c models.ControleProtese
		err := rows.Scan(
			&c.ProteseId, &c.Ficha, &c.PacienteId, &c.LaboratorioId, &c.Produto,
			&c.EtapaAtual, &c.TrabalhoFinal, &c.DataEnvio, &c.DataEntregaPrevista,
			&c.DataRecebimento, &c.Status, &c.Observacoes,
			&c.CustoUnitario, &c.Quantidade, &c.CustoTotal,
		)
		if err != nil {
			return nil, err
		}
		lista = append(lista, c)
	}

	return lista, nil
}
