package models

type LaboratorioFechamento struct {
    FechamentoId     int     `json:"fechamento_id"`
    LaboratorioId    int     `json:"laboratorio_id"`
    DataFinalizacao  string  `json:"data_finalizacao"` 
    Produto          string  `json:"produto"`
    PacienteId       *int    `json:"paciente_id,omitempty"`
    Quantidade       int     `json:"quantidade"`
    PrecoUnitario    float64 `json:"preco_unitario"`
    Total            float64 `json:"total"`
}
