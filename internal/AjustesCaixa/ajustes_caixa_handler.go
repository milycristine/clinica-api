package ajustecaixa

import (
	"clinica-api/internal/models"
	"encoding/json"
	"net/http"
)

type AjusteCaixaHandler struct {
	service AjusteCaixaService
}

func NovoAjusteCaixaHandler(s AjusteCaixaService) *AjusteCaixaHandler {
	return &AjusteCaixaHandler{service: s}
}

func (h *AjusteCaixaHandler) Criar(w http.ResponseWriter, r *http.Request) {
	var a models.AjusteCaixa
	json.NewDecoder(r.Body).Decode(&a)

	err := h.service.Criar(&a)
	json.NewEncoder(w).Encode(models.ResponseDefaultModel{
		IsSuccess: err == nil,
	})
}

func (h *AjusteCaixaHandler) Listar(w http.ResponseWriter, r *http.Request) {
	lista, err := h.service.Listar()
	json.NewEncoder(w).Encode(models.ResponseDefaultModel{
		IsSuccess: err == nil,
		Data:      lista,
	})
}
