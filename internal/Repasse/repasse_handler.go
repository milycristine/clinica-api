package repasse

import (
	"clinica-api/internal/models"
	"encoding/json"
	"net/http"
	"strconv"
)

type RepasseHandler struct {
	service RepasseService
}

func NovoRepasseHandler(s RepasseService) *RepasseHandler {
	return &RepasseHandler{service: s}
}

func (h *RepasseHandler) Criar(w http.ResponseWriter, r *http.Request) {
	var rep models.Repasse
	json.NewDecoder(r.Body).Decode(&rep)

	err := h.service.Criar(&rep)
	json.NewEncoder(w).Encode(models.ResponseDefaultModel{
		IsSuccess: err == nil,
		ErrorMessage: func() string {
			if err != nil {
				return err.Error()
			}
			return ""
		}(),
	})
}

func (h *RepasseHandler) Listar(w http.ResponseWriter, r *http.Request) {
	lista, err := h.service.Listar()
	json.NewEncoder(w).Encode(models.ResponseDefaultModel{
		IsSuccess: err == nil,
		Data:      lista,
	})
}

func (h *RepasseHandler) Pagar(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	err := h.service.Pagar(id)
	json.NewEncoder(w).Encode(models.ResponseDefaultModel{
		IsSuccess: err == nil,
	})
}
