package models

type LeadHistorico struct {
    HistoricoId   int    `json:"historico_id"`
    ContatoId     int    `json:"contato_id"`
    DataRegistro  string `json:"data_registro"`
    Descricao     string `json:"descricao"`
    FuncionarioId *int   `json:"funcionario_id,omitempty"`
}
