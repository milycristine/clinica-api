package models

type Procedimento struct {
    ProcedimentoId int    `json:"procedimento_id"`
    Nome           string `json:"nome"`
    Ativo          bool   `json:"ativo"`
}

type LeadProcedimento struct {
    LeadProcedimentoId int `json:"lead_procedimento_id"`
    ContatoId          int `json:"contato_id"`
    ProcedimentoId     int `json:"procedimento_id"`
}
