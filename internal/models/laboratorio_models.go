package models

type Laboratorio struct {
    Id       int    `json:"id"`
    Nome     string `json:"nome"`
    Contato  string `json:"contato"`
    Telefone string `json:"telefone"`
    Email    string `json:"email"`
}

type LaboratorioFiltro struct {
    Nome    string `json:"nome"`
    Contato string `json:"contato"`
    Email   string `json:"email"`
    Page    int    `json:"page"`
    Limit   int    `json:"limit"`
}
