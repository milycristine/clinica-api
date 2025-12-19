package models

type GlosaDetalhe struct {
	GlosaDetalheId         int    `json:"glosa_detalhe_id"`
	ConvenioProcedimentoId int    `json:"convenio_procedimento_id"`
	Guia                   string `json:"guia"`
	DataOcorrencia         string `json:"data_ocorrencia"`
	PacienteId             int    `json:"paciente_id"`

	Mo             string `json:"mo"`
	NomeContratado string `json:"nome_contratado"`
	DenteRegiao    string `json:"dente_regiao"`

	MotivoGlosa   string `json:"motivo_glosa"`
	StatusRecurso string `json:"status_recurso"`

	GlosasMensalId *int `json:"glosas_mensal_id"`
	UnidadeId      int  `json:"unidade_id"`

	Procedimentos []GlosaProcedimento `json:"procedimentos"`
}

type GlosaProcedimento struct {
	GlosaProcedimentoId int     `json:"id"`
	GlosaDetalheId      int     `json:"glosa_detalhe_id"`
	ProcedimentoId      *int    `json:"procedimento_id"`
	CodigoProcedimento  string  `json:"codigo_procedimento"`
	NomeProcedimento    string  `json:"nome_procedimento"`
	Quantidade          int     `json:"quantidade"`
	ValorInformado      float64 `json:"valor_informado"`
	ValorGlosado        float64 `json:"valor_glosado"`
}
