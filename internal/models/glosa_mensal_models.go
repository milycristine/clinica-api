	package models

	import "time"

	type GlosaMensal struct {
		GlosasMensalId  int        `json:"id"`
		Operadora       string     `json:"operadora"`
		MesReferencia   int        `json:"mes_referencia"`
		AnoReferencia   int        `json:"ano_referencia"`
		ValorInformado  float64    `json:"valor_informado"`
		ValorGlosa      float64    `json:"valor_glosa"`
		ValorPago       float64    `json:"valor_pago"`
		Observacoes     string     `json:"observacoes"`
		UnidadeId       int        `json:"unidade_id"`
		DataAtualizacao *time.Time `json:"data_atualizacao"`
	}


	type GlosaMensalFiltro struct {
		MesReferencia int `json:"mes_referencia"`
		AnoReferencia int `json:"ano_referencia"`
		UnidadeId     int `json:"unidade_id"`
	}
