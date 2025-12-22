package models

type ConvenioProcedimento struct {
	ConvenioProcedimentoId int `json:"convenioProcedimentoId"`
	ConvenioId             int `json:"convenioId"`

	ProcedimentoId     int    `json:"procedimentoId"`
	CodigoProcedimento string `json:"codigoProcedimento"`
	NomeProcedimento   string `json:"nomeProcedimento"`
	DenteRegiao        string `json:"denteRegiao"`
	Quantidade         int    `json:"quantidade"`

	ValorInformado float64  `json:"valorInformado"` 
	ValorPago      *float64 `json:"valorPago"`      

	Status string `json:"status"` 

	Observacoes string `json:"observacoes"`
}

