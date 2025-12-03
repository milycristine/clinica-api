package captacao

import (
	"clinica-api/internal/models"
	"encoding/json"
	"net/http"
	"strconv"
)

type ContatoHandler interface {
	CriarContato(w http.ResponseWriter, r *http.Request)
	EditarContato(w http.ResponseWriter, r *http.Request)
	ListarContatos(w http.ResponseWriter, r *http.Request)
	BuscarContatoPorID(w http.ResponseWriter, r *http.Request)
	AtualizarStatus(w http.ResponseWriter, r *http.Request)
}

type contatoHandler struct {
	service ContatoService
}

func NovoContatoHandler(service ContatoService) ContatoHandler {
	return &contatoHandler{service: service}
}

func (h *contatoHandler) CriarContato(w http.ResponseWriter, r *http.Request) {
	var c models.Leads
	response := models.ResponseDefaultModel{IsSuccess: true, Data: c}

	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		response.IsSuccess = false
		response.ErrorMessage = "Erro ao decodificar dados"
		w.WriteHeader(http.StatusBadRequest)
	} else if err := h.service.CriarContato(&c); err != nil {
		response.IsSuccess = false
		response.ErrorMessage = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *contatoHandler) EditarContato(w http.ResponseWriter, r *http.Request) {
	var c models.Leads
	response := models.ResponseDefaultModel{IsSuccess: true, Data: c}

	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		response.IsSuccess = false
		response.ErrorMessage = "Erro ao decodificar dados"
		w.WriteHeader(http.StatusBadRequest)
	} else if err := h.service.EditarContato(&c); err != nil {
		response.IsSuccess = false
		response.ErrorMessage = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *contatoHandler) ListarContatos(w http.ResponseWriter, r *http.Request) {
	contatos, err := h.service.ListarContatos()
	response := models.ResponseDefaultModel{IsSuccess: err == nil, Data: contatos}

	if err != nil {
		response.ErrorMessage = "Erro ao listar contatos"
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *contatoHandler) BuscarContatoPorID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	contato, err := h.service.BuscarContatoPorID(id)
	if err != nil {
		http.Error(w, "Erro ao buscar contato", 500)
		return
	}

	if contato == nil {
		http.Error(w, "Contato não encontrado", 404)
		return
	}

	response := models.ResponseDefaultModel{
		IsSuccess: true,
		Data:      contato,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *contatoHandler) AtualizarStatus(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	status := r.URL.Query().Get("status")

	id, _ := strconv.Atoi(idStr)

	err := h.service.AtualizarStatus(id, status)
	response := models.ResponseDefaultModel{IsSuccess: err == nil}

	if err != nil {
		response.ErrorMessage = "Erro ao atualizar status"
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
