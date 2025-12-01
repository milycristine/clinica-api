package paciente

import (
	database "clinica-api/database"
	"clinica-api/internal/models"
	"database/sql"
	"fmt"
)

type PacienteRepository interface {
	CriarPaciente(p *models.Paciente) error
	EditarPaciente(p *models.Paciente) error
	ListarPacientes() ([]models.Paciente, error)
	BuscarPacientePorID(id int) (*models.Paciente, error)
	ListarPorUnidade(unidadeId int) ([]models.Paciente, error)
}

type pacienteRepository struct {
	db *sql.DB
}

func NovoPacienteRepository(conn *database.SQLStr) PacienteRepository {
	return &pacienteRepository{
		db: conn.DB(),
	}
}

func (r *pacienteRepository) CriarPaciente(p *models.Paciente) error {

	query := `
        INSERT INTO Pacientes (
            Prontuario, Nome, Mae, Pai, Nascimento, Sexo, Rg, Cpf, Cns, Telefone1, Telefone2,
            Email, Observacoes, DataCadastro, Logradouro, Numero, Complemento, Bairro, Cidade, Uf, Cep,
            NomeResponsavel, CpfResponsavel, ProfissaoResponsavel, EstadoCivilResponsavel,
            Foto, Profissao, EstadoCivil, OrigemPaciente, UnidadeId
        )
        VALUES (
            @Prontuario, @Nome, @Mae, @Pai, @Nascimento, @Sexo, @Rg, @Cpf, @Cns, @Telefone1, @Telefone2,
            @Email, @Observacoes, @DataCadastro, @Logradouro, @Numero, @Complemento, @Bairro, @Cidade, @Uf, @Cep,
            @NomeResponsavel, @CpfResponsavel, @ProfissaoResponsavel, @EstadoCivilResponsavel,
            @Foto, @Profissao, @EstadoCivil, @OrigemPaciente, @UnidadeId
        )
    `

	_, err := r.db.Exec(query,
		sql.Named("Prontuario", p.Prontuario),
		sql.Named("Nome", p.Nome),
		sql.Named("Mae", p.Mae),
		sql.Named("Pai", p.Pai),
		sql.Named("Nascimento", p.Nascimento),
		sql.Named("Sexo", p.Sexo),
		sql.Named("Rg", p.Rg),
		sql.Named("Cpf", p.Cpf),
		sql.Named("Cns", p.Cns),
		sql.Named("Telefone1", p.Telefone1),
		sql.Named("Telefone2", p.Telefone2),
		sql.Named("Email", p.Email),
		sql.Named("Observacoes", p.Observacoes),
		sql.Named("DataCadastro", p.DataCadastro),
		sql.Named("Logradouro", p.Logradouro),
		sql.Named("Numero", p.Numero),
		sql.Named("Complemento", p.Complemento),
		sql.Named("Bairro", p.Bairro),
		sql.Named("Cidade", p.Cidade),
		sql.Named("Uf", p.Uf),
		sql.Named("Cep", p.Cep),
		sql.Named("NomeResponsavel", p.NomeResponsavel),
		sql.Named("CpfResponsavel", p.CpfResponsavel),
		sql.Named("ProfissaoResponsavel", p.ProfissaoResponsavel),
		sql.Named("EstadoCivilResponsavel", p.EstadoCivilResponsavel),
		sql.Named("Foto", p.Foto),
		sql.Named("Profissao", p.Profissao),
		sql.Named("EstadoCivil", p.EstadoCivil),
		sql.Named("OrigemPaciente", p.OrigemPaciente),
		sql.Named("UnidadeId", p.UnidadeId),
	)

	if err != nil {
		return fmt.Errorf("erro ao criar paciente: %w", err)
	}

	return nil
}

func (r *pacienteRepository) EditarPaciente(p *models.Paciente) error {

	query := `
        UPDATE Pacientes SET
            Prontuario=@Prontuario, Nome=@Nome, Mae=@Mae, Pai=@Pai,
            Nascimento=@Nascimento, Sexo=@Sexo, Rg=@Rg, Cpf=@Cpf, Cns=@Cns,
            Telefone1=@Telefone1, Telefone2=@Telefone2, Email=@Email,
            Observacoes=@Observacoes, Logradouro=@Logradouro, Numero=@Numero,
            Complemento=@Complemento, Bairro=@Bairro, Cidade=@Cidade, Uf=@Uf,
            Cep=@Cep, NomeResponsavel=@NomeResponsavel, CpfResponsavel=@CpfResponsavel,
            ProfissaoResponsavel=@ProfissaoResponsavel, EstadoCivilResponsavel=@EstadoCivilResponsavel,
            Foto=@Foto, Profissao=@Profissao, EstadoCivil=@EstadoCivil,
            OrigemPaciente=@OrigemPaciente, UnidadeId=@UnidadeId
        WHERE PacienteId=@PacienteId
    `

	_, err := r.db.Exec(query,
		sql.Named("PacienteId", p.PacienteId),
		sql.Named("Prontuario", p.Prontuario),
		sql.Named("Nome", p.Nome),
		sql.Named("Mae", p.Mae),
		sql.Named("Pai", p.Pai),
		sql.Named("Nascimento", p.Nascimento),
		sql.Named("Sexo", p.Sexo),
		sql.Named("Rg", p.Rg),
		sql.Named("Cpf", p.Cpf),
		sql.Named("Cns", p.Cns),
		sql.Named("Telefone1", p.Telefone1),
		sql.Named("Telefone2", p.Telefone2),
		sql.Named("Email", p.Email),
		sql.Named("Observacoes", p.Observacoes),
		sql.Named("Logradouro", p.Logradouro),
		sql.Named("Numero", p.Numero),
		sql.Named("Complemento", p.Complemento),
		sql.Named("Bairro", p.Bairro),
		sql.Named("Cidade", p.Cidade),
		sql.Named("Uf", p.Uf),
		sql.Named("Cep", p.Cep),
		sql.Named("NomeResponsavel", p.NomeResponsavel),
		sql.Named("CpfResponsavel", p.CpfResponsavel),
		sql.Named("ProfissaoResponsavel", p.ProfissaoResponsavel),
		sql.Named("EstadoCivilResponsavel", p.EstadoCivilResponsavel),
		sql.Named("Foto", p.Foto),
		sql.Named("Profissao", p.Profissao),
		sql.Named("EstadoCivil", p.EstadoCivil),
		sql.Named("OrigemPaciente", p.OrigemPaciente),
		sql.Named("UnidadeId", p.UnidadeId),
	)

	return err
}

func (r *pacienteRepository) ListarPacientes() ([]models.Paciente, error) {

	rows, err := r.db.Query(`
        SELECT 
            PacienteId, Prontuario, Nome, Mae, Pai, Nascimento, Sexo, Rg, Cpf, Cns,
            Telefone1, Telefone2, Email, Observacoes, DataCadastro, Logradouro, Numero,
            Complemento, Bairro, Cidade, Uf, Cep, NomeResponsavel, CpfResponsavel,
            ProfissaoResponsavel, EstadoCivilResponsavel, Foto, Profissao, EstadoCivil,
            OrigemPaciente, UnidadeId
        FROM Pacientes WITH (NOLOCK)
        ORDER BY Nome ASC
    `)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pacientes []models.Paciente

	for rows.Next() {
		var p models.Paciente
		err := rows.Scan(
			&p.PacienteId, &p.Prontuario, &p.Nome, &p.Mae, &p.Pai, &p.Nascimento, &p.Sexo,
			&p.Rg, &p.Cpf, &p.Cns, &p.Telefone1, &p.Telefone2, &p.Email, &p.Observacoes,
			&p.DataCadastro, &p.Logradouro, &p.Numero, &p.Complemento, &p.Bairro, &p.Cidade,
			&p.Uf, &p.Cep, &p.NomeResponsavel, &p.CpfResponsavel, &p.ProfissaoResponsavel,
			&p.EstadoCivilResponsavel, &p.Foto, &p.Profissao, &p.EstadoCivil,
			&p.OrigemPaciente, &p.UnidadeId,
		)
		if err != nil {
			return nil, err
		}
		pacientes = append(pacientes, p)
	}

	return pacientes, nil
}

func (r *pacienteRepository) BuscarPacientePorID(id int) (*models.Paciente, error) {

	query := `
        SELECT 
            PacienteId, Prontuario, Nome, Mae, Pai, Nascimento, Sexo, Rg, Cpf, Cns,
            Telefone1, Telefone2, Email, Observacoes, DataCadastro, Logradouro, Numero,
            Complemento, Bairro, Cidade, Uf, Cep, NomeResponsavel, CpfResponsavel,
            ProfissaoResponsavel, EstadoCivilResponsavel, Foto, Profissao, EstadoCivil,
            OrigemPaciente, UnidadeId
        FROM Pacientes
        WHERE PacienteId = @Id
    `

	var p models.Paciente

	err := r.db.QueryRow(query, sql.Named("Id", id)).Scan(
		&p.PacienteId, &p.Prontuario, &p.Nome, &p.Mae, &p.Pai, &p.Nascimento, &p.Sexo,
		&p.Rg, &p.Cpf, &p.Cns, &p.Telefone1, &p.Telefone2, &p.Email, &p.Observacoes,
		&p.DataCadastro, &p.Logradouro, &p.Numero, &p.Complemento, &p.Bairro, &p.Cidade,
		&p.Uf, &p.Cep, &p.NomeResponsavel, &p.CpfResponsavel, &p.ProfissaoResponsavel,
		&p.EstadoCivilResponsavel, &p.Foto, &p.Profissao, &p.EstadoCivil,
		&p.OrigemPaciente, &p.UnidadeId,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *pacienteRepository) ListarPorUnidade(unidadeId int) ([]models.Paciente, error) {
	query := `
		SELECT 
			PacienteId, Prontuario, Nome, Mae, Pai, Nascimento, Sexo, Rg, Cpf, Cns,
			Telefone1, Telefone2, Email, Observacoes, DataCadastro, Logradouro, Numero,
			Complemento, Bairro, Cidade, Uf, Cep, NomeResponsavel, CpfResponsavel,
			ProfissaoResponsavel, EstadoCivilResponsavel, Foto, Profissao, EstadoCivil,
			OrigemPaciente, UnidadeId
		FROM Pacientes
		WHERE UnidadeId = @UnidadeId
	`

	rows, err := r.db.Query(query, sql.Named("UnidadeId", unidadeId))
	if err != nil {
		return nil, fmt.Errorf("erro ao listar pacientes por unidade: %w", err)
	}
	defer rows.Close()

	var pacientes []models.Paciente

	for rows.Next() {
		var p models.Paciente
		err := rows.Scan(
			&p.PacienteId, &p.Prontuario, &p.Nome, &p.Mae, &p.Pai, &p.Nascimento,
			&p.Sexo, &p.Rg, &p.Cpf, &p.Cns, &p.Telefone1, &p.Telefone2, &p.Email,
			&p.Observacoes, &p.DataCadastro, &p.Logradouro, &p.Numero, &p.Complemento,
			&p.Bairro, &p.Cidade, &p.Uf, &p.Cep, &p.NomeResponsavel, &p.CpfResponsavel,
			&p.ProfissaoResponsavel, &p.EstadoCivilResponsavel, &p.Foto, &p.Profissao,
			&p.EstadoCivil, &p.OrigemPaciente, &p.UnidadeId,
		)

		if err != nil {
			return nil, err
		}

		pacientes = append(pacientes, p)
	}

	return pacientes, nil
}
