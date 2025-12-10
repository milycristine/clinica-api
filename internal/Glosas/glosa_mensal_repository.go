package glosa

import (
	database "clinica-api/database"
	"clinica-api/internal/models"
	"database/sql"
	"fmt"
	"strings"
)

type GlosaMensalRepository interface {
	CriarGlosaMensal(gm *models.GlosaMensal) error
	EditarGlosaMensal(gm *models.GlosaMensal) error
	ListarGlosasMensais(f models.GlosaMensalFiltro) ([]models.GlosaMensal, error)
	BuscarGlosaMensalPorID(id int) (*models.GlosaMensal, error)
	RecalcularGlosaMensal(glosasMensalId int) error
	BuscarOuCriarMensal(mes int, ano int, unidadeId int) (int, error)
}

type glosaMensalRepository struct {
	db *sql.DB
}

func NovoGlosaMensalRepository(conn *database.SQLStr) GlosaMensalRepository {
	return &glosaMensalRepository{db: conn.DB()}
}

func (r *glosaMensalRepository) CriarGlosaMensal(gm *models.GlosaMensal) error {
	query := `
        INSERT INTO GlosasMensais (MesReferencia, AnoReferencia, ValorInformado, ValorGlosa, Observacoes, UnidadeId)
        VALUES (@Mes, @Ano, @ValorInformado, @ValorGlosa, @Observacoes, @Unidade)
    `

	_, err := r.db.Exec(query,
		sql.Named("Mes", gm.MesReferencia),
		sql.Named("Ano", gm.AnoReferencia),
		sql.Named("ValorInformado", gm.ValorInformado),
		sql.Named("ValorGlosa", gm.ValorGlosa),
		sql.Named("Observacoes", gm.Observacoes),
		sql.Named("Unidade", gm.UnidadeId),
	)

	if err != nil {
		return fmt.Errorf("erro ao criar glosa mensal: %w", err)
	}

	return nil
}

func (r *glosaMensalRepository) EditarGlosaMensal(gm *models.GlosaMensal) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	set := []string{}
	args := []any{}

	if gm.ValorInformado > 0 {
		set = append(set, "ValorInformado = @ValorInformado")
		args = append(args, sql.Named("ValorInformado", gm.ValorInformado))
	}
	if gm.ValorGlosa > 0 {
		set = append(set, "ValorGlosa = @ValorGlosa")
		args = append(args, sql.Named("ValorGlosa", gm.ValorGlosa))
	}
	if gm.ValorPago > 0 {
		set = append(set, "ValorPago = @ValorPago")
		args = append(args, sql.Named("ValorPago", gm.ValorPago))
	}
	if gm.Observacoes != "" {
		set = append(set, "Observacoes = @Observacoes")
		args = append(args, sql.Named("Observacoes", gm.Observacoes))
	}

	if len(set) == 0 {
		tx.Rollback()
		return fmt.Errorf("nenhum campo para atualizar")
	}

	if gm.GlosasMensalId == 0 {
		tx.Rollback()
		return fmt.Errorf("id é obrigatório")
	}

	args = append(args, sql.Named("Id", gm.GlosasMensalId))

	query := fmt.Sprintf(`
        UPDATE GlosasMensais 
        SET %s, DataAtualizacao = GETDATE()
        WHERE GlosasMensalId = @Id
    `, strings.Join(set, ", "))

	_, err = tx.Exec(query, args...)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("erro ao atualizar glosa mensal: %w", err)
	}

	return tx.Commit()
}

func (r *glosaMensalRepository) ListarGlosasMensais(f models.GlosaMensalFiltro) ([]models.GlosaMensal, error) {
	lista := []models.GlosaMensal{}

	query := `
        SELECT GlosasMensalId, MesReferencia, AnoReferencia, ValorInformado, ValorGlosa, ValorPago, Observacoes, UnidadeId, DataAtualizacao
        FROM GlosasMensais WITH (NOLOCK)
        WHERE (@Mes = 0 OR MesReferencia = @Mes)
        AND (@Ano = 0 OR AnoReferencia = @Ano)
        AND (@Unidade = 0 OR UnidadeId = @Unidade)
        ORDER BY AnoReferencia DESC, MesReferencia DESC
    `

	rows, err := r.db.Query(query,
		sql.Named("Mes", f.MesReferencia),
		sql.Named("Ano", f.AnoReferencia),
		sql.Named("Unidade", f.UnidadeId),
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var gm models.GlosaMensal
		if err := rows.Scan(
			&gm.GlosasMensalId, &gm.MesReferencia, &gm.AnoReferencia,
			&gm.ValorInformado, &gm.ValorGlosa, &gm.ValorPago,
			&gm.Observacoes, &gm.UnidadeId, &gm.DataAtualizacao,
		); err != nil {
			return nil, err
		}
		lista = append(lista, gm)
	}

	return lista, nil
}

func (r *glosaMensalRepository) BuscarGlosaMensalPorID(id int) (*models.GlosaMensal, error) {
	var gm models.GlosaMensal

	query := `
        SELECT GlosasMensalId, MesReferencia, AnoReferencia, ValorInformado, ValorGlosa,
               ValorPago, Observacoes, UnidadeId, DataAtualizacao
        FROM GlosasMensais
        WHERE GlosasMensalId = @Id
    `

	err := r.db.QueryRow(query, sql.Named("Id", id)).Scan(
		&gm.GlosasMensalId, &gm.MesReferencia, &gm.AnoReferencia,
		&gm.ValorInformado, &gm.ValorGlosa, &gm.ValorPago,
		&gm.Observacoes, &gm.UnidadeId, &gm.DataAtualizacao,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &gm, nil
}

func (r *glosaMensalRepository) RecalcularGlosaMensal(glosasMensalId int) error {
	query := `
        UPDATE GlosasMensais
        SET 
            ValorInformado = (
                SELECT ISNULL(SUM(ValorInformado), 0)
                FROM GlosaProcedimentos GP
                INNER JOIN GlosasDetalhes GD ON GP.GlosaDetalheId = GD.GlosaDetalheId
                WHERE GD.GlosasMensalId = @Id
            ),
            ValorGlosa = (
                SELECT ISNULL(SUM(ValorGlosado), 0)
                FROM GlosaProcedimentos GP
                INNER JOIN GlosasDetalhes GD ON GP.GlosaDetalheId = GD.GlosaDetalheId
                WHERE GD.GlosasMensalId = @Id
            ),
            ValorPago = (
                SELECT ISNULL(SUM(GP.ValorInformado - GP.ValorGlosado), 0)
                FROM GlosaProcedimentos GP
                INNER JOIN GlosasDetalhes GD ON GP.GlosaDetalheId = GD.GlosaDetalheId
                WHERE GD.GlosasMensalId = @Id
            ),
            DataAtualizacao = GETDATE()
        WHERE GlosasMensalId = @Id
    `

	_, err := r.db.Exec(query, sql.Named("Id", glosasMensalId))
	if err != nil {
		return fmt.Errorf("erro ao recalcular mensal: %w", err)
	}

	return nil
}

func (r *glosaMensalRepository) BuscarOuCriarMensal(mes int, ano int, unidadeId int) (int, error) {
	query := `
        SELECT GlosasMensalId 
        FROM GlosasMensais 
        WHERE MesReferencia = @Mes AND AnoReferencia = @Ano AND UnidadeId = @Unidade;
    `

	var id int
	err := r.db.QueryRow(query,
		sql.Named("Mes", mes),
		sql.Named("Ano", ano),
		sql.Named("Unidade", unidadeId),
	).Scan(&id)

	if err == nil {
		return id, nil
	}

	if err != sql.ErrNoRows {
		return 0, err
	}

	insert := `
        INSERT INTO GlosasMensais (MesReferencia, AnoReferencia, ValorInformado, ValorGlosa, ValorPago, Observacoes, UnidadeId)
        OUTPUT INSERTED.GlosasMensalId
        VALUES (@Mes, @Ano, 0, 0, 0, '', @Unidade)
    `

	err = r.db.QueryRow(insert,
		sql.Named("Mes", mes),
		sql.Named("Ano", ano),
		sql.Named("Unidade", unidadeId),
	).Scan(&id)

	return id, err
}
