package despesa

import (
	"clinica-api/internal/models"
	"encoding/json"
	"net/http"
	"strconv"
)

type DespesaHandler interface {
	CriarDespesa(w http.ResponseWriter, r *http.Request)
	EditarDespesa(w http.ResponseWriter, r *http.Request)
	ListarDespesas(w http.ResponseWriter, r *http.Request)
	BuscarDespesaPorID(w http.ResponseWriter, r *http.Request)
	AtualizarStatus(w http.ResponseWriter, r *http.Request)
}

type despesaHandler struct {
	service DespesaService
}

func NovaDespesaHandler(service DespesaService) DespesaHandler {
	return &despesaHandler{service: service}
}

func (h *despesaHandler) CriarDespesa(w http.ResponseWriter, r *http.Request) {
	var despesa models.Despesa

	response := models.ResponseDefaultModel{
		IsSuccess: true,
		Data:      despesa,
	}

	if err := json.NewDecoder(r.Body).Decode(&despesa); err != nil {
		response.IsSuccess = false
		response.ErrorMessage = "Erro ao decodificar dados da despesa"
		w.WriteHeader(http.StatusBadRequest)
	} else if err := h.service.CriarDespesa(&despesa); err != nil {
		response.IsSuccess = false
		response.ErrorMessage = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *despesaHandler) EditarDespesa(w http.ResponseWriter, r *http.Request) {
	var despesa models.Despesa

	response := models.ResponseDefaultModel{
		IsSuccess: true,
		Data:      despesa,
	}

	if err := json.NewDecoder(r.Body).Decode(&despesa); err != nil {
		response.IsSuccess = false
		response.ErrorMessage = "Erro ao decodificar dados da despesa"
		w.WriteHeader(http.StatusBadRequest)
	} else if err := h.service.EditarDespesa(&despesa); err != nil {
		response.IsSuccess = false
		response.ErrorMessage = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *despesaHandler) ListarDespesas(w http.ResponseWriter, r *http.Request) {
	despesas, err := h.service.ListarDespesas()

	response := models.ResponseDefaultModel{
		IsSuccess: true,
		Data:      despesas,
	}

	if err != nil {
		response.IsSuccess = false
		response.ErrorMessage = "Erro ao listar despesas"
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *despesaHandler) BuscarDespesaPorID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)

	if err != nil || id == 0 {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	despesa, err := h.service.BuscarDespesaPorID(id)
	if err != nil {
		http.Error(w, "Erro ao buscar despesa", http.StatusInternalServerError)
		return
	}

	if despesa == nil {
		http.Error(w, "Despesa não encontrada", http.StatusNotFound)
		return
	}

	response := models.ResponseDefaultModel{
		IsSuccess: true,
		Data:      despesa,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *despesaHandler) AtualizarStatus(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	status := r.URL.Query().Get("status")

	id, err := strconv.Atoi(idStr)
	if err != nil || id == 0 {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if status == "" {
		http.Error(w, "Status é obrigatório", http.StatusBadRequest)
		return
	}

	err = h.service.AtualizarStatus(id, status)

	response := models.ResponseDefaultModel{
		IsSuccess: err == nil,
	}

	if err != nil {
		response.ErrorMessage = "Erro ao atualizar status da despesa"
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
