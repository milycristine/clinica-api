package models

type AjusteCaixa struct {
	AjusteId  int     `json:"ajuste_id"`
	Data      string  `json:"data"`
	Tipo      string  `json:"tipo"`
	Descricao string  `json:"descricao"`
	Valor     float64 `json:"valor"`
	Origem    string  `json:"origem"`
	Destino   string  `json:"destino"`
	Status    string  `json:"status"`
	UnidadeId int     `json:"unidade_id"`
}
