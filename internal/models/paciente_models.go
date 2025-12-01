package models

type Paciente struct {
	PacienteId             int     `json:"pacienteId"`
	Prontuario             string  `json:"prontuario"`
	Nome                   string  `json:"nome"`
	Mae                    *string `json:"mae"`
	Pai                    *string `json:"pai"`
	Nascimento             string  `json:"nascimento"`
	Sexo                   string  `json:"sexo"`
	Rg                     string  `json:"rg"`
	Cpf                    string  `json:"cpf"`
	Cns                    *string  `json:"cns"`
	Telefone1              string  `json:"telefone1"`
	Telefone2              *string `json:"telefone2"`
	Email                  string  `json:"email"`
	Observacoes            string  `json:"observacoes"`
	DataCadastro           string  `json:"dataCadastro"`
	Logradouro             string  `json:"logradouro"`
	Numero                 string  `json:"numero"`
	Complemento            *string `json:"complemento"`
	Bairro                 string  `json:"bairro"`
	Cidade                 string  `json:"cidade"`
	Uf                     string  `json:"uf"`
	Cep                    string  `json:"cep"`
	NomeResponsavel        *string `json:"nomeResponsavel"`
	CpfResponsavel         *string `json:"cpfResponsavel"`
	ProfissaoResponsavel   *string `json:"profissaoResponsavel"`
	EstadoCivilResponsavel *string `json:"estadoCivilResponsavel"`
	Foto                   *string `json:"foto"`
	Profissao              string  `json:"profissao"`
	EstadoCivil            string  `json:"estadoCivil"`
	OrigemPaciente         *string `json:"origemPaciente"`
	UnidadeId              int     `json:"unidadeId"`
}
