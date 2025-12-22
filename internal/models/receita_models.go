package models

type Receita struct {
	ReceitaId       int     `json:"receitaId"`
	Data            string  `json:"data"`
	Procedimento    string  `json:"procedimento"`
	PacienteId      *int    `json:"pacienteId"`
	ValorBruto      float64 `json:"valorBruto"`
	ValorLiquido    float64 `json:"valorLiquido"`
	FormaPagamento  string  `json:"formaPagamento"`
	UnidadeId       int     `json:"unidadeId"`
	ProfissionalId  int     `json:"profissionalId"`
	TipoReceita     string  `json:"tipoReceita"`
	Status          string  `json:"status"`
	Observacoes     string  `json:"observacoes"`
}
