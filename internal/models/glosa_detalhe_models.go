package models

type GlosaDetalhe struct {
	GlosaDetalheId     int     `json:"glosa_detalhe_id"`
	Guia               string  `json:"guia"`
	DataOcorrencia     string  `json:"data_ocorrencia"`
	PacienteId         int     `json:"paciente_id"`
	Mo                 string  `json:"mo"`
	NomeContratado     string  `json:"nome_contratado"`
	CodigoProcedimento string  `json:"codigo_procedimento"`
	NomeProcedimento   string  `json:"nome_procedimento"`
	DenteRegiao        string  `json:"dente_regiao"`
	ValorInformado     float64 `json:"valor_informado"`
	ValorGlosado       float64 `json:"valor_glosado"`
	MotivoGlosa        string  `json:"motivo_glosa"`
	StatusRecurso      string  `json:"status_recurso"`
	GlosasMensalId     *int    `json:"glosas_mensal_id"`
	UnidadeId          int     `json:"unidade_id"`  
}
