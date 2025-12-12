package models

type Laboratorio struct {
    Id       int    `json:"id"`
    Nome     string `json:"nome"`
    Contato  string `json:"contato"`
    Telefone string `json:"telefone"`
    Email    string `json:"email"`
}
