package convenio

import (
	"clinica-api/internal/models"
	"encoding/json"
	"net/http"
	"strconv"
)

type ConvenioProcedimentoHandler interface {
	Criar(w http.ResponseWriter, r *http.Request)
	ListarPorConvenio(w http.ResponseWriter, r *http.Request)
	AtualizarStatus(w http.ResponseWriter, r *http.Request)
}

type convenioProcedimentoHandler struct {
	service ConvenioProcedimentoService
}

func NovoConvenioProcedimentoHandler(
	service ConvenioProcedimentoService,
) ConvenioProcedimentoHandler {
	return &convenioProcedimentoHandler{
		service: service,
	}
}
func (h *convenioProcedimentoHandler) Criar(w http.ResponseWriter, r *http.Request) {
	var p models.ConvenioProcedimento
	response := models.ResponseDefaultModel{IsSuccess: true}

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		response.IsSuccess = false
		response.ErrorMessage = "Erro ao decodificar dados"
		w.WriteHeader(http.StatusBadRequest)
	} else if err := h.service.Criar(&p); err != nil {
		response.IsSuccess = false
		response.ErrorMessage = err.Error()
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusCreated)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func (h *convenioProcedimentoHandler) ListarPorConvenio(w http.ResponseWriter, r *http.Request) {
	convenioId, err := strconv.Atoi(r.URL.Query().Get("convenioId"))
	if err != nil || convenioId == 0 {
		http.Error(w, "convenioId inválido", http.StatusBadRequest)
		return
	}

	procedimentos, err := h.service.ListarPorConvenio(convenioId)
	if err != nil {
		http.Error(w, "Erro ao listar procedimentos", http.StatusInternalServerError)
		return
	}

	response := models.ResponseDefaultModel{
		IsSuccess: true,
		Data:      procedimentos,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func (h *convenioProcedimentoHandler) AtualizarStatus(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		ConvenioProcedimentoId int      `json:"convenioProcedimentoId"`
		Status                 string   `json:"status"` 
		ValorPago              *float64 `json:"valorPago"`
		Observacoes            string   `json:"observacoes"`
	}

	resp := models.ResponseDefaultModel{IsSuccess: true}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		resp.IsSuccess = false
		resp.ErrorMessage = "payload inválido"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp)
		return
	}

	if err := h.service.AtualizarStatus(payload.ConvenioProcedimentoId, payload.Status, payload.ValorPago, payload.Observacoes); err != nil {
		resp.IsSuccess = false
		resp.ErrorMessage = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(resp)
}
