package models

type Repasse struct {
	RepasseId      int     `json:"repasse_id"`
	ReceitaId      int     `json:"receita_id"`
	ProfissionalId int     `json:"profissional_id"`
	Percentual     float64 `json:"percentual"`
	ValorFixo      float64 `json:"valor_fixo"`
	ValorCalculado float64 `json:"valor_calculado"`
	Status         string  `json:"status"`
	Observacoes    string  `json:"observacoes"`
}
