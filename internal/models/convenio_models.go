package models

type Convenio struct {
	ConvenioId      int    `json:"convenioId"`
	UnidadeId       int    `json:"unidadeId"`
	PacienteId      int    `json:"pacienteId"`
	Guia            string `json:"guia"`
	Operadora       string `json:"operadora"`
	DataAtendimento string `json:"dataAtendimento"`

	Status string `json:"status"` 

	CreatedAt string `json:"createdAt"`

	Procedimentos []ConvenioProcedimento `json:"procedimentos"`
}
