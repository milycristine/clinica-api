package models

type GlosaMensal struct {
	GlosasMensalId  int     `json:"id"`
	MesReferencia   int     `json:"mes_referencia"`
	AnoReferencia   int     `json:"ano_referencia"`
	ValorInformado  float64 `json:"valor_informado"`
	ValorGlosa      float64 `json:"valor_glosa"`
	ValorPago       float64 `json:"valor_pago"`
	Observacoes     string  `json:"observacoes"`
	UnidadeId       int     `json:"unidade_id"`
	DataAtualizacao *string `json:"data_atualizacao"`
}

type GlosaMensalFiltro struct {
	MesReferencia int `json:"mes_referencia"`
	AnoReferencia int `json:"ano_referencia"`
	UnidadeId     int `json:"unidade_id"`
}
