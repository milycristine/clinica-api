package captacao

import (
	"clinica-api/internal/models"
	"encoding/json"
	"net/http"
	"strconv"
)

type LeadHistoricoHandler interface {
	CriarHistorico(w http.ResponseWriter, r *http.Request)
	ListarHistoricoPorLead(w http.ResponseWriter, r *http.Request)
}

type leadHistoricoHandler struct {
	service LeadHistoricoService
}

func NovoLeadHistoricoHandler(service LeadHistoricoService) LeadHistoricoHandler {
	return &leadHistoricoHandler{service: service}
}

func (h *leadHistoricoHandler) CriarHistorico(w http.ResponseWriter, r *http.Request) {
	var historico models.LeadHistorico
	response := models.ResponseDefaultModel{IsSuccess: true, Data: historico}

	if err := json.NewDecoder(r.Body).Decode(&historico); err != nil {
		response.IsSuccess = false
		response.ErrorMessage = "Erro ao decodificar dados"
		w.WriteHeader(http.StatusBadRequest)
	} else if err := h.service.CriarHistorico(&historico); err != nil {
		response.IsSuccess = false
		response.ErrorMessage = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *leadHistoricoHandler) ListarHistoricoPorLead(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("contatoId")
	contatoId, _ := strconv.Atoi(idStr)

	historico, err := h.service.ListarHistoricoPorLead(contatoId)
	response := models.ResponseDefaultModel{
		IsSuccess: err == nil,
		Data:      historico,
	}

	if err != nil {
		response.ErrorMessage = "Erro ao listar histórico"
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
