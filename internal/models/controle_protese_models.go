package models

type ControleProtese struct {
	ProteseId           int     `json:"protese_id"`
	Ficha               string  `json:"ficha"`
	PacienteId          int     `json:"paciente_id"`
	LaboratorioId       int     `json:"laboratorio_id"`
	Produto             string  `json:"produto"`
	EtapaAtual          string  `json:"etapa_atual"`
	TrabalhoFinal       string  `json:"trabalho_final"`
	DataEnvio           string  `json:"data_envio"`
	DataEntregaPrevista string  `json:"data_entrega_prevista"`
	DataRecebimento     string  `json:"data_recebimento"`
	Status              string  `json:"status"`
	Observacoes         string  `json:"observacoes"`
	CustoUnitario       float64 `json:"custo_unitario"`
	Quantidade          int     `json:"quantidade"`
	CustoTotal          float64 `json:"custo_total"`
}
