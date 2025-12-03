package models

type Leads struct {
	ContatoId            int    `json:"contato_id"`
	DataContato          string `json:"data_contato"`
	MeioContato          string `json:"meio_contato"`
	Nome                 string `json:"nome"`
	Telefone             string `json:"telefone"`
	ProcedimentoInteresse string `json:"procedimento_interesse"`
	Observacoes          string `json:"observacoes"`
	Status               string `json:"status"`
	DataAtualizacao      string `json:"data_atualizacao"`
	FuncionarioId        *int   `json:"funcionario_id,omitempty"`
}
