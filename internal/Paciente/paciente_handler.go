package paciente

import (
	"clinica-api/internal/models"
	"encoding/json"
	"net/http"
	"strconv"
)

type PacienteHandler interface {
	CriarPaciente(w http.ResponseWriter, r *http.Request)
	EditarPaciente(w http.ResponseWriter, r *http.Request)
	ListarPacientes(w http.ResponseWriter, r *http.Request)
	BuscarPacientePorID(w http.ResponseWriter, r *http.Request)
	ListarPorUnidade(w http.ResponseWriter, r *http.Request)
}

type pacienteHandler struct {
	service PacienteService
}

func NovoPacienteHandler(service PacienteService) PacienteHandler {
	return &pacienteHandler{service: service}
}

func (h *pacienteHandler) CriarPaciente(w http.ResponseWriter, r *http.Request) {

	var paciente models.Paciente
	response := models.ResponseDefaultModel{IsSuccess: true, Data: paciente}

	if err := json.NewDecoder(r.Body).Decode(&paciente); err != nil {
		response.IsSuccess = false
		response.ErrorMessage = "Erro ao decodificar JSON"
		w.WriteHeader(http.StatusBadRequest)
	} else if err := h.service.CriarPaciente(&paciente); err != nil {
		response.IsSuccess = false
		response.ErrorMessage = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *pacienteHandler) EditarPaciente(w http.ResponseWriter, r *http.Request) {

	var paciente models.Paciente
	response := models.ResponseDefaultModel{IsSuccess: true, Data: paciente}

	if err := json.NewDecoder(r.Body).Decode(&paciente); err != nil {
		response.IsSuccess = false
		response.ErrorMessage = "Erro ao decodificar JSON"
		w.WriteHeader(http.StatusBadRequest)
	} else if err := h.service.EditarPaciente(&paciente); err != nil {
		response.IsSuccess = false
		response.ErrorMessage = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *pacienteHandler) ListarPacientes(w http.ResponseWriter, r *http.Request) {

	pacientes, err := h.service.ListarPacientes()
	response := models.ResponseDefaultModel{IsSuccess: true, Data: pacientes}

	if err != nil {
		response.IsSuccess = false
		response.ErrorMessage = "Erro ao listar pacientes"
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *pacienteHandler) BuscarPacientePorID(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	paciente, err := h.service.BuscarPacientePorID(id)
	if err != nil {
		http.Error(w, "Erro ao buscar paciente", 500)
		return
	}

	if paciente == nil {
		http.Error(w, "Paciente não encontrado", 404)
		return
	}

	response := models.ResponseDefaultModel{
		IsSuccess: true,
		Data:      paciente,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func (h *pacienteHandler) ListarPorUnidade(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	unidadeId, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "ID da unidade inválido", http.StatusBadRequest)
		return
	}

	pacientes, err := h.service.ListarPorUnidade(unidadeId)
	if err != nil {
		http.Error(w, "Erro ao listar pacientes por unidade", http.StatusInternalServerError)
		return
	}

	response := models.ResponseDefaultModel{
		IsSuccess: true,
		Data:      pacientes,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
