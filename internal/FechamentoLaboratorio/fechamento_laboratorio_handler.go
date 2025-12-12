package fechamentolaboratorio

import (
	"clinica-api/internal/models"
	"encoding/json"
	"net/http"
	"strconv"
)

type FechamentoHandler interface {
	Criar(w http.ResponseWriter, r *http.Request)
	Editar(w http.ResponseWriter, r *http.Request)
	Listar(w http.ResponseWriter, r *http.Request)
	BuscarPorID(w http.ResponseWriter, r *http.Request)
	ListarPorLaboratorio(w http.ResponseWriter, r *http.Request)
	ListarPorData(w http.ResponseWriter, r *http.Request)
	ListarPorMes(w http.ResponseWriter, r *http.Request)
	ListarPorPeriodo(w http.ResponseWriter, r *http.Request)
	ListarPorLaboratorioMes(w http.ResponseWriter, r *http.Request)
	ListarPorPaciente(w http.ResponseWriter, r *http.Request)
}

type fechamentoHandler struct {
	service FechamentoService
}

func NovoFechamentoHandler(s FechamentoService) FechamentoHandler {
	return &fechamentoHandler{service: s}
}

func (h *fechamentoHandler) Criar(w http.ResponseWriter, r *http.Request) {
	var f models.LaboratorioFechamento
	if err := json.NewDecoder(r.Body).Decode(&f); err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	if f.Quantidade == 0 {
		f.Quantidade = 1
	}

	if err := h.service.Criar(&f); err != nil {
		resp := models.ResponseDefaultModel{IsSuccess: false, ErrorMessage: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}

	resp := models.ResponseDefaultModel{IsSuccess: true, Data: f}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (h *fechamentoHandler) Editar(w http.ResponseWriter, r *http.Request) {
	var f models.LaboratorioFechamento
	if err := json.NewDecoder(r.Body).Decode(&f); err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	if err := h.service.Editar(&f); err != nil {
		resp := models.ResponseDefaultModel{IsSuccess: false, ErrorMessage: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}

	resp := models.ResponseDefaultModel{IsSuccess: true, Data: f}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *fechamentoHandler) Listar(w http.ResponseWriter, r *http.Request) {
	lista, err := h.service.Listar()
	resp := models.ResponseDefaultModel{IsSuccess: err == nil, Data: lista}
	if err != nil {
		resp.ErrorMessage = "Erro ao listar fechamentos"
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *fechamentoHandler) BuscarPorID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)
	if id == 0 {
		http.Error(w, "id é obrigatório", http.StatusBadRequest)
		return
	}

	item, err := h.service.BuscarPorID(id)
	if err != nil {
		http.Error(w, "Erro ao buscar fechamento", http.StatusInternalServerError)
		return
	}
	if item == nil {
		http.Error(w, "Fechamento não encontrado", http.StatusNotFound)
		return
	}

	resp := models.ResponseDefaultModel{IsSuccess: true, Data: item}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *fechamentoHandler) ListarPorLaboratorio(w http.ResponseWriter, r *http.Request) {
	labStr := r.URL.Query().Get("laboratorio_id")
	labId, _ := strconv.Atoi(labStr)
	if labId == 0 {
		http.Error(w, "laboratorio_id é obrigatório", http.StatusBadRequest)
		return
	}

	lista, err := h.service.ListarPorLaboratorio(labId)
	resp := models.ResponseDefaultModel{IsSuccess: err == nil, Data: lista}
	if err != nil {
		resp.ErrorMessage = "Erro ao listar fechamentos por laboratório"
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *fechamentoHandler) ListarPorData(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query().Get("data")
	if data == "" {
		http.Error(w, "data é obrigatória (YYYY-MM-DD)", http.StatusBadRequest)
		return
	}

	lista, err := h.service.ListarPorData(data)
	resp := models.ResponseDefaultModel{IsSuccess: err == nil, Data: lista}
	if err != nil {
		resp.ErrorMessage = "Erro ao listar fechamentos por data"
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
func (h *fechamentoHandler) ListarPorMes(w http.ResponseWriter, r *http.Request) {
	mes, _ := strconv.Atoi(r.URL.Query().Get("mes"))
	ano, _ := strconv.Atoi(r.URL.Query().Get("ano"))

	lista, err := h.service.ListarPorMes(mes, ano)
	response := models.ResponseDefaultModel{IsSuccess: err == nil, Data: lista}

	if err != nil {
		response.ErrorMessage = err.Error()
	}

	json.NewEncoder(w).Encode(response)
}
func (h *fechamentoHandler) ListarPorPeriodo(w http.ResponseWriter, r *http.Request) {
	inicio := r.URL.Query().Get("inicio")
	fim := r.URL.Query().Get("fim")

	lista, err := h.service.ListarPorPeriodo(inicio, fim)
	response := models.ResponseDefaultModel{IsSuccess: err == nil, Data: lista}

	if err != nil {
		response.ErrorMessage = err.Error()
	}

	json.NewEncoder(w).Encode(response)
}
func (h *fechamentoHandler) ListarPorLaboratorioMes(w http.ResponseWriter, r *http.Request) {
	labId, _ := strconv.Atoi(r.URL.Query().Get("laboratorioId"))
	mes, _ := strconv.Atoi(r.URL.Query().Get("mes"))
	ano, _ := strconv.Atoi(r.URL.Query().Get("ano"))

	lista, err := h.service.ListarPorLaboratorioMes(labId, mes, ano)
	response := models.ResponseDefaultModel{IsSuccess: err == nil, Data: lista}

	if err != nil {
		response.ErrorMessage = err.Error()
	}

	json.NewEncoder(w).Encode(response)
}
func (h *fechamentoHandler) ListarPorPaciente(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("pacienteId"))

	lista, err := h.service.ListarPorPaciente(id)
	response := models.ResponseDefaultModel{IsSuccess: err == nil, Data: lista}

	if err != nil {
		response.ErrorMessage = err.Error()
	}

	json.NewEncoder(w).Encode(response)
}
