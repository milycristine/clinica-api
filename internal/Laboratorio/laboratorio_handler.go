package laboratorio

import (
	"clinica-api/internal/models"
	"encoding/json"
	"net/http"
	"strconv"
)

type LaboratorioHandler interface {
	CriarLaboratorio(w http.ResponseWriter, r *http.Request)
	EditarLaboratorio(w http.ResponseWriter, r *http.Request)
	ListarLaboratorios(w http.ResponseWriter, r *http.Request)
	BuscarLaboratorioPorID(w http.ResponseWriter, r *http.Request)
	

}

type laboratorioHandler struct {
	service LaboratorioService
}

func NovoLaboratorioHandler(service LaboratorioService) LaboratorioHandler {
	return &laboratorioHandler{service: service}
}

func (h *laboratorioHandler) CriarLaboratorio(w http.ResponseWriter, r *http.Request) {
	var l models.Laboratorio
	resp := models.ResponseDefaultModel{IsSuccess: true, Data: l}

	if err := json.NewDecoder(r.Body).Decode(&l); err != nil {
		resp.IsSuccess = false
		resp.ErrorMessage = "Erro ao decodificar dados"
		w.WriteHeader(http.StatusBadRequest)
	} else if err := h.service.CriarLaboratorio(&l); err != nil {
		resp.IsSuccess = false
		resp.ErrorMessage = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *laboratorioHandler) EditarLaboratorio(w http.ResponseWriter, r *http.Request) {
	var l models.Laboratorio
	resp := models.ResponseDefaultModel{IsSuccess: true, Data: l}

	if err := json.NewDecoder(r.Body).Decode(&l); err != nil {
		resp.IsSuccess = false
		resp.ErrorMessage = "Erro ao decodificar dados"
		w.WriteHeader(http.StatusBadRequest)
	} else if err := h.service.EditarLaboratorio(&l); err != nil {
		resp.IsSuccess = false
		resp.ErrorMessage = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *laboratorioHandler) ListarLaboratorios(w http.ResponseWriter, r *http.Request) {
	lista, err := h.service.ListarLaboratorios()

	resp := models.ResponseDefaultModel{IsSuccess: true, Data: lista}
	if err != nil {
		resp.IsSuccess = false
		resp.ErrorMessage = "Erro ao listar laboratórios"
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *laboratorioHandler) BuscarLaboratorioPorID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	l, err := h.service.BuscarLaboratorioPorID(id)
	if err != nil {
		http.Error(w, "Erro ao buscar laboratório", 500)
		return
	}
	if l == nil {
		http.Error(w, "Laboratório não encontrado", 404)
		return
	}

	resp := models.ResponseDefaultModel{
		IsSuccess: true,
		Data:      l,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
