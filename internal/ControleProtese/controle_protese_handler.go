package controleprotese

import (
	"clinica-api/internal/models"
	"encoding/json"
	"net/http"
	"strconv"
)

type ControleProteseHandler interface {
	Criar(w http.ResponseWriter, r *http.Request)
	Editar(w http.ResponseWriter, r *http.Request)
	Listar(w http.ResponseWriter, r *http.Request)
	BuscarPorID(w http.ResponseWriter, r *http.Request)
	AlterarStatus(w http.ResponseWriter, r *http.Request)
	ListarFiltrado(w http.ResponseWriter, r *http.Request)
}

type controleHandler struct {
	service ControleProteseService
}

func NovoControleHandler(s ControleProteseService) ControleProteseHandler {
	return &controleHandler{service: s}
}

func (h *controleHandler) Criar(w http.ResponseWriter, r *http.Request) {
	var m models.ControleProtese

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, "Erro ao decodificar JSON", 400)
		return
	}

	err := h.service.Criar(&m)
	resp := models.ResponseDefaultModel{IsSuccess: err == nil, Data: m}

	if err != nil {
		resp.ErrorMessage = err.Error()
		w.WriteHeader(400)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *controleHandler) Editar(w http.ResponseWriter, r *http.Request) {
	var m models.ControleProtese

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, "Erro ao decodificar JSON", 400)
		return
	}

	err := h.service.Editar(&m)
	resp := models.ResponseDefaultModel{IsSuccess: err == nil, Data: m}

	if err != nil {
		resp.ErrorMessage = err.Error()
		w.WriteHeader(400)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *controleHandler) Listar(w http.ResponseWriter, r *http.Request) {

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	lista, err := h.service.Listar(page, limit)
	resp := models.ResponseDefaultModel{IsSuccess: err == nil, Data: lista}

	if err != nil {
		resp.ErrorMessage = "Erro ao listar próteses"
		w.WriteHeader(500)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *controleHandler) BuscarPorID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	item, err := h.service.BuscarPorID(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if item == nil {
		http.Error(w, "Prótese não encontrada", 404)
		return
	}

	resp := models.ResponseDefaultModel{IsSuccess: true, Data: item}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *controleHandler) AlterarStatus(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	status := r.URL.Query().Get("status")

	id, _ := strconv.Atoi(idStr)

	err := h.service.AlterarStatus(id, status)
	resp := models.ResponseDefaultModel{IsSuccess: err == nil}

	if err != nil {
		resp.ErrorMessage = err.Error()
		w.WriteHeader(400)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *controleHandler) ListarFiltrado(w http.ResponseWriter, r *http.Request) {

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	data := r.URL.Query().Get("data")
	dataInicio := r.URL.Query().Get("dataInicio")
	dataFim := r.URL.Query().Get("dataFim")
	status := r.URL.Query().Get("status")
	produto := r.URL.Query().Get("produto")
	etapa := r.URL.Query().Get("etapa")

	laboratorioId, _ := strconv.Atoi(r.URL.Query().Get("laboratorioId"))
	pacienteId, _ := strconv.Atoi(r.URL.Query().Get("pacienteId"))

	lista, err := h.service.ListarFiltrado(
		data, dataInicio, dataFim, status,
		laboratorioId, pacienteId, produto, etapa,
		page, limit,
	)

	resp := models.ResponseDefaultModel{
		IsSuccess: err == nil,
		Data:      lista,
	}

	if err != nil {
		resp.ErrorMessage = err.Error()
		w.WriteHeader(400)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
