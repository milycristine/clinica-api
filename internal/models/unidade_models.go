package models

type Unidade struct {
    Id       int    `json:"id"`
    Nome     string `json:"nome"`
    Endereco string `json:"endereco"`
    Telefone string `json:"telefone"`
    Email    string `json:"email"`
    Status   int    `json:"status"`
}
