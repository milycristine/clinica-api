package convenio

import (
	"clinica-api/internal/models"
	"encoding/json"
	"net/http"
	"strconv"
)

type ConvenioHandler interface {
	CriarConvenio(w http.ResponseWriter, r *http.Request)
	EditarConvenio(w http.ResponseWriter, r *http.Request)
	ListarConvenios(w http.ResponseWriter, r *http.Request)
	BuscarConvenioPorID(w http.ResponseWriter, r *http.Request)
	AtualizarStatus(w http.ResponseWriter, r *http.Request)
}

type convenioHandler struct {
	service ConvenioService
}

func NovoConvenioHandler(service ConvenioService) ConvenioHandler {
	return &convenioHandler{service: service}
}

func (h *convenioHandler) CriarConvenio(w http.ResponseWriter, r *http.Request) {
	var c models.Convenio
	response := models.ResponseDefaultModel{IsSuccess: true}

	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		response.IsSuccess = false
		response.ErrorMessage = "Erro ao decodificar dados"
		w.WriteHeader(http.StatusBadRequest)
	} else if err := h.service.CriarConvenio(&c); err != nil {
		response.IsSuccess = false
		response.ErrorMessage = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *convenioHandler) EditarConvenio(w http.ResponseWriter, r *http.Request) {
	var c models.Convenio
	response := models.ResponseDefaultModel{IsSuccess: true}

	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		response.IsSuccess = false
		response.ErrorMessage = "Erro ao decodificar dados"
		w.WriteHeader(http.StatusBadRequest)
	} else if err := h.service.EditarConvenio(&c); err != nil {
		response.IsSuccess = false
		response.ErrorMessage = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *convenioHandler) ListarConvenios(w http.ResponseWriter, r *http.Request) {
	convenios, err := h.service.ListarConvenios()
	response := models.ResponseDefaultModel{IsSuccess: true, Data: convenios}

	if err != nil {
		response.IsSuccess = false
		response.ErrorMessage = "Erro ao listar convênios"
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *convenioHandler) BuscarConvenioPorID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	convenio, err := h.service.BuscarConvenioPorID(id)
	if err != nil {
		http.Error(w, "Erro ao buscar convênio", 500)
		return
	}

	if convenio == nil {
		http.Error(w, "Convênio não encontrado", 404)
		return
	}

	response := models.ResponseDefaultModel{
		IsSuccess: true,
		Data:      convenio,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func (h *convenioHandler) AtualizarStatus(w http.ResponseWriter, r *http.Request) {

	var req struct {
		ConvenioId             int      `json:"convenioId"`
		ConvenioProcedimentoId int      `json:"convenioProcedimentoId"`
		Status                 string   `json:"status"`
		ValorPago              *float64 `json:"valorPago"`
		Observacoes            string   `json:"observacoes"`

		MotivoOdontologico string `json:"motivoOdontologico"`
		NomeContratado     string `json:"nomeContratado"`
	}

	response := models.ResponseDefaultModel{IsSuccess: true}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.IsSuccess = false
		response.ErrorMessage = "Erro ao decodificar dados"
		w.WriteHeader(http.StatusBadRequest)
	} else if err := h.service.AtualizarStatusProcedimento(
		req.ConvenioId,
		req.ConvenioProcedimentoId,
		req.Status,
		req.ValorPago,
		req.Observacoes,
		req.MotivoOdontologico,
		req.NomeContratado,
	); err != nil {
		response.IsSuccess = false
		response.ErrorMessage = err.Error()
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
