package models

type Despesa struct {
	DespesaId        int     `json:"despesa_id"`
	Data             string  `json:"data"`
	Fornecedor       string  `json:"fornecedor"`
	Nf               string  `json:"nf"`
	ProdutoDescricao string  `json:"produto_descricao"`
	Valor            float64 `json:"valor"`
	FormaPagamento   string  `json:"forma_pagamento"`
	DataVencimento   string  `json:"data_vencimento"`
	Status           string  `json:"status"`
	UnidadeId        int     `json:"unidade_id"`
	Observacoes      string  `json:"observacoes"`
}
