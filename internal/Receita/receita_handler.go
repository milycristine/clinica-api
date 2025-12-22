package receita

import (
	"clinica-api/internal/models"
	"encoding/json"
	"net/http"
	"strconv"
)

type ReceitaHandler interface {
	CriarReceita(w http.ResponseWriter, r *http.Request)
	EditarReceita(w http.ResponseWriter, r *http.Request)
	ListarReceitas(w http.ResponseWriter, r *http.Request)
	BuscarReceitaPorID(w http.ResponseWriter, r *http.Request)
	AtualizarStatus(w http.ResponseWriter, r *http.Request)
}

type receitaHandler struct {
	service ReceitaService
}

func NovaReceitaHandler(service ReceitaService) ReceitaHandler {
	return &receitaHandler{service: service}
}

func (h *receitaHandler) CriarReceita(w http.ResponseWriter, r *http.Request) {
	var rec models.Receita
	resp := models.ResponseDefaultModel{IsSuccess: true}

	if err := json.NewDecoder(r.Body).Decode(&rec); err != nil {
		resp.IsSuccess = false
		resp.ErrorMessage = "Erro ao decodificar dados"
		w.WriteHeader(http.StatusBadRequest)
	} else if err := h.service.CriarReceita(&rec); err != nil {
		resp.IsSuccess = false
		resp.ErrorMessage = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *receitaHandler) EditarReceita(w http.ResponseWriter, r *http.Request) {
	var rec models.Receita
	resp := models.ResponseDefaultModel{IsSuccess: true}

	if err := json.NewDecoder(r.Body).Decode(&rec); err != nil {
		resp.IsSuccess = false
		resp.ErrorMessage = "Erro ao decodificar dados"
		w.WriteHeader(http.StatusBadRequest)
	} else if err := h.service.EditarReceita(&rec); err != nil {
		resp.IsSuccess = false
		resp.ErrorMessage = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *receitaHandler) ListarReceitas(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.ListarReceitas()
	resp := models.ResponseDefaultModel{IsSuccess: true, Data: data}

	if err != nil {
		resp.IsSuccess = false
		resp.ErrorMessage = "Erro ao listar receitas"
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *receitaHandler) BuscarReceitaPorID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	rec, err := h.service.BuscarReceitaPorID(id)

	if err != nil {
		http.Error(w, "Erro ao buscar receita", 500)
		return
	}
	if rec == nil {
		http.Error(w, "Receita não encontrada", 404)
		return
	}

	resp := models.ResponseDefaultModel{IsSuccess: true, Data: rec}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *receitaHandler) AtualizarStatus(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	status := r.URL.Query().Get("status")

	err := h.service.AtualizarStatus(id, status)
	resp := models.ResponseDefaultModel{IsSuccess: err == nil}

	if err != nil {
		resp.ErrorMessage = "Erro ao atualizar status"
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
