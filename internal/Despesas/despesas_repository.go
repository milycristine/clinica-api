package despesa

import (
	database "clinica-api/database"
	"clinica-api/internal/models"
	"database/sql"
	"fmt"
	"strings"
)

type DespesaRepository interface {
	CriarDespesa(d *models.Despesa) error
	EditarDespesa(d *models.Despesa) error
	ListarDespesas() ([]models.Despesa, error)
	BuscarDespesaPorID(id int) (*models.Despesa, error)
	AtualizarStatus(id int, status string) error
}

type despesaRepository struct {
	db *sql.DB
}

func NovaDespesaRepository(conn *database.SQLStr) DespesaRepository {
	return &despesaRepository{
		db: conn.DB(),
	}
}
func (r *despesaRepository) CriarDespesa(d *models.Despesa) error {
	query := `
		INSERT INTO Despesas
		(Data, Fornecedor, Nf, ProdutoDescricao, Valor, FormaPagamento,
		 DataVencimento, Status, UnidadeId, Observacoes)
		VALUES
		(@Data, @Fornecedor, @Nf, @ProdutoDescricao, @Valor, @FormaPagamento,
		 @DataVencimento, @Status, @UnidadeId, @Observacoes)
	`

	_, err := r.db.Exec(query,
		sql.Named("Data", d.Data),
		sql.Named("Fornecedor", d.Fornecedor),
		sql.Named("Nf", d.Nf),
		sql.Named("ProdutoDescricao", d.ProdutoDescricao),
		sql.Named("Valor", d.Valor),
		sql.Named("FormaPagamento", d.FormaPagamento),
		sql.Named("DataVencimento", d.DataVencimento),
		sql.Named("Status", "A_PAGAR"),
		sql.Named("UnidadeId", d.UnidadeId),
		sql.Named("Observacoes", d.Observacoes),
	)

	if err != nil {
		return fmt.Errorf("erro ao criar despesa: %w", err)
	}
	return nil
}
func (r *despesaRepository) EditarDespesa(d *models.Despesa) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	var sets []string
	var args []any

	if d.Fornecedor != "" {
		sets = append(sets, "Fornecedor = @Fornecedor")
		args = append(args, sql.Named("Fornecedor", d.Fornecedor))
	}
	if d.ProdutoDescricao != "" {
		sets = append(sets, "ProdutoDescricao = @ProdutoDescricao")
		args = append(args, sql.Named("ProdutoDescricao", d.ProdutoDescricao))
	}
	if d.Valor > 0 {
		sets = append(sets, "Valor = @Valor")
		args = append(args, sql.Named("Valor", d.Valor))
	}
	if d.FormaPagamento != "" {
		sets = append(sets, "FormaPagamento = @FormaPagamento")
		args = append(args, sql.Named("FormaPagamento", d.FormaPagamento))
	}
	if d.Status != "" {
		sets = append(sets, "Status = @Status")
		args = append(args, sql.Named("Status", d.Status))
	}
	if d.Observacoes != "" {
		sets = append(sets, "Observacoes = @Observacoes")
		args = append(args, sql.Named("Observacoes", d.Observacoes))
	}

	if len(sets) == 0 || d.DespesaId == 0 {
		tx.Rollback()
		return fmt.Errorf("dados insuficientes para atualização")
	}

	args = append(args, sql.Named("DespesaId", d.DespesaId))

	query := fmt.Sprintf(
		"UPDATE Despesas SET %s WHERE DespesaId = @DespesaId",
		strings.Join(sets, ", "),
	)

	_, err = tx.Exec(query, args...)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
func (r *despesaRepository) ListarDespesas() ([]models.Despesa, error) {
	rows, err := r.db.Query(`
		SELECT DespesaId, Data, Fornecedor, Nf, ProdutoDescricao, Valor,
		       FormaPagamento, DataVencimento, Status, UnidadeId, Observacoes
		FROM Despesas WITH (NOLOCK)
		ORDER BY Data DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var despesas []models.Despesa
	for rows.Next() {
		var d models.Despesa
		rows.Scan(
			&d.DespesaId,
			&d.Data,
			&d.Fornecedor,
			&d.Nf,
			&d.ProdutoDescricao,
			&d.Valor,
			&d.FormaPagamento,
			&d.DataVencimento,
			&d.Status,
			&d.UnidadeId,
			&d.Observacoes,
		)
		despesas = append(despesas, d)
	}
	return despesas, nil
}
func (r *despesaRepository) BuscarDespesaPorID(id int) (*models.Despesa, error) {
	var d models.Despesa

	err := r.db.QueryRow(`
		SELECT DespesaId, Data, Fornecedor, Nf, ProdutoDescricao, Valor,
		       FormaPagamento, DataVencimento, Status, UnidadeId, Observacoes
		FROM Despesas WHERE DespesaId = @Id`,
		sql.Named("Id", id),
	).Scan(
		&d.DespesaId,
		&d.Data,
		&d.Fornecedor,
		&d.Nf,
		&d.ProdutoDescricao,
		&d.Valor,
		&d.FormaPagamento,
		&d.DataVencimento,
		&d.Status,
		&d.UnidadeId,
		&d.Observacoes,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &d, err
}
func (r *despesaRepository) AtualizarStatus(id int, status string) error {
	_, err := r.db.Exec(
		"UPDATE Despesas SET Status = @Status WHERE DespesaId = @Id",
		sql.Named("Id", id),
		sql.Named("Status", status),
	)
	return err
}
