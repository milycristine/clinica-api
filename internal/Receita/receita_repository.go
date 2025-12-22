package receita

import (
	database "clinica-api/database"
	"clinica-api/internal/models"
	"database/sql"
	"fmt"
	"strings"
)

type ReceitaRepository interface {
	CriarReceita(r *models.Receita) error
	EditarReceita(r *models.Receita) error
	ListarReceitas() ([]models.Receita, error)
	BuscarReceitaPorID(id int) (*models.Receita, error)
	AtualizarStatus(id int, status string) error
}

type receitaRepository struct {
	db *sql.DB
}

func NovaReceitaRepository(conn *database.SQLStr) ReceitaRepository {
	return &receitaRepository{
		db: conn.DB(),
	}
}

func (r *receitaRepository) CriarReceita(rec *models.Receita) error {
	query := `
	INSERT INTO Receitas
	(Data, Procedimento, PacienteId, ValorBruto, ValorLiquido,
	 FormaPagamento, UnidadeId, ProfissionalId, TipoReceita, Status, Observacoes)
	VALUES
	(@Data, @Procedimento, @PacienteId, @ValorBruto, @ValorLiquido,
	 @FormaPagamento, @UnidadeId, @ProfissionalId, @TipoReceita, @Status, @Observacoes)
	`

	_, err := r.db.Exec(
		query,
		sql.Named("Data", rec.Data),
		sql.Named("Procedimento", rec.Procedimento),
		sql.Named("PacienteId", rec.PacienteId),
		sql.Named("ValorBruto", rec.ValorBruto),
		sql.Named("ValorLiquido", rec.ValorLiquido),
		sql.Named("FormaPagamento", rec.FormaPagamento),
		sql.Named("UnidadeId", rec.UnidadeId),
		sql.Named("ProfissionalId", rec.ProfissionalId),
		sql.Named("TipoReceita", rec.TipoReceita),
		sql.Named("Status", rec.Status),
		sql.Named("Observacoes", rec.Observacoes),
	)

	if err != nil {
		return fmt.Errorf("erro ao criar receita: %w", err)
	}
	return nil
}

func (r *receitaRepository) EditarReceita(rec *models.Receita) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	var sets []string
	var args []any

	if rec.Data != "" {
		sets = append(sets, "Data = @Data")
		args = append(args, sql.Named("Data", rec.Data))
	}
	if rec.Procedimento != "" {
		sets = append(sets, "Procedimento = @Procedimento")
		args = append(args, sql.Named("Procedimento", rec.Procedimento))
	}
	if rec.PacienteId != nil {
		sets = append(sets, "PacienteId = @PacienteId")
		args = append(args, sql.Named("PacienteId", rec.PacienteId))
	}
	if rec.ValorBruto > 0 {
		sets = append(sets, "ValorBruto = @ValorBruto")
		args = append(args, sql.Named("ValorBruto", rec.ValorBruto))
	}
	if rec.ValorLiquido > 0 {
		sets = append(sets, "ValorLiquido = @ValorLiquido")
		args = append(args, sql.Named("ValorLiquido", rec.ValorLiquido))
	}
	if rec.FormaPagamento != "" {
		sets = append(sets, "FormaPagamento = @FormaPagamento")
		args = append(args, sql.Named("FormaPagamento", rec.FormaPagamento))
	}
	if rec.TipoReceita != "" {
		sets = append(sets, "TipoReceita = @TipoReceita")
		args = append(args, sql.Named("TipoReceita", rec.TipoReceita))
	}
	if rec.Status != "" {
		sets = append(sets, "Status = @Status")
		args = append(args, sql.Named("Status", rec.Status))
	}
	if rec.Observacoes != "" {
		sets = append(sets, "Observacoes = @Observacoes")
		args = append(args, sql.Named("Observacoes", rec.Observacoes))
	}

	if len(sets) == 0 || rec.ReceitaId == 0 {
		tx.Rollback()
		return fmt.Errorf("dados inválidos para atualização")
	}

	args = append(args, sql.Named("ReceitaId", rec.ReceitaId))

	query := fmt.Sprintf(
		"UPDATE Receitas SET %s WHERE ReceitaId = @ReceitaId",
		strings.Join(sets, ", "),
	)

	if _, err := tx.Exec(query, args...); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *receitaRepository) ListarReceitas() ([]models.Receita, error) {
	rows, err := r.db.Query(`
		SELECT ReceitaId, Data, Procedimento, PacienteId, ValorBruto,
		       ValorLiquido, FormaPagamento, UnidadeId, ProfissionalId,
		       TipoReceita, Status, Observacoes
		FROM Receitas WITH (NOLOCK)
		ORDER BY Data DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lista []models.Receita

	for rows.Next() {
		var rec models.Receita
		if err := rows.Scan(
			&rec.ReceitaId,
			&rec.Data,
			&rec.Procedimento,
			&rec.PacienteId,
			&rec.ValorBruto,
			&rec.ValorLiquido,
			&rec.FormaPagamento,
			&rec.UnidadeId,
			&rec.ProfissionalId,
			&rec.TipoReceita,
			&rec.Status,
			&rec.Observacoes,
		); err != nil {
			return nil, err
		}
		lista = append(lista, rec)
	}
	return lista, nil
}

func (r *receitaRepository) BuscarReceitaPorID(id int) (*models.Receita, error) {
	var rec models.Receita

	err := r.db.QueryRow(`
		SELECT ReceitaId, Data, Procedimento, PacienteId, ValorBruto,
		       ValorLiquido, FormaPagamento, UnidadeId, ProfissionalId,
		       TipoReceita, Status, Observacoes
		FROM Receitas WHERE ReceitaId = @Id`,
		sql.Named("Id", id),
	).Scan(
		&rec.ReceitaId,
		&rec.Data,
		&rec.Procedimento,
		&rec.PacienteId,
		&rec.ValorBruto,
		&rec.ValorLiquido,
		&rec.FormaPagamento,
		&rec.UnidadeId,
		&rec.ProfissionalId,
		&rec.TipoReceita,
		&rec.Status,
		&rec.Observacoes,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &rec, err
}

func (r *receitaRepository) AtualizarStatus(id int, status string) error {
	_, err := r.db.Exec(
		"UPDATE Receitas SET Status = @Status WHERE ReceitaId = @Id",
		sql.Named("Id", id),
		sql.Named("Status", status),
	)
	return err
}
