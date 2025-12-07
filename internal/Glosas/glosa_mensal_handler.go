package glosa

import (
	"clinica-api/internal/models"
	"encoding/json"
	"net/http"
	"strconv"
)

type GlosaMensalHandler interface {
	CriarGlosaMensal(w http.ResponseWriter, r *http.Request)
	EditarGlosaMensal(w http.ResponseWriter, r *http.Request)
	ListarGlosasMensais(w http.ResponseWriter, r *http.Request)
	BuscarGlosaMensalPorID(w http.ResponseWriter, r *http.Request)
}

type glosaMensalHandler struct {
	service GlosaMensalService
}

func NovoGlosaMensalHandler(service GlosaMensalService) GlosaMensalHandler {
	return &glosaMensalHandler{service: service}
}

func (h *glosaMensalHandler) CriarGlosaMensal(w http.ResponseWriter, r *http.Request) {
	var gm models.GlosaMensal
	response := models.ResponseDefaultModel{IsSuccess: true, Data: gm}

	if err := json.NewDecoder(r.Body).Decode(&gm); err != nil {
		response.IsSuccess = false
		response.ErrorMessage = "Erro ao decodificar dados"
		w.WriteHeader(http.StatusBadRequest)
	} else if err := h.service.CriarGlosaMensal(&gm); err != nil {
		response.IsSuccess = false
		response.ErrorMessage = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *glosaMensalHandler) EditarGlosaMensal(w http.ResponseWriter, r *http.Request) {
	var gm models.GlosaMensal
	response := models.ResponseDefaultModel{IsSuccess: true}

	if err := json.NewDecoder(r.Body).Decode(&gm); err != nil {
		response.IsSuccess = false
		response.ErrorMessage = "Erro ao decodificar dados"
		w.WriteHeader(http.StatusBadRequest)
	} else if err := h.service.EditarGlosaMensal(&gm); err != nil {
		response.IsSuccess = false
		response.ErrorMessage = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		response.Data = "Glosa mensal atualizada com sucesso"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *glosaMensalHandler) ListarGlosasMensais(w http.ResponseWriter, r *http.Request) {
	mes, _ := strconv.Atoi(r.URL.Query().Get("mes"))
	ano, _ := strconv.Atoi(r.URL.Query().Get("ano"))
	unidade, _ := strconv.Atoi(r.URL.Query().Get("unidade"))

	filtro := models.GlosaMensalFiltro{
		MesReferencia: mes,
		AnoReferencia: ano,
		UnidadeId:     unidade,
	}

	glosas, err := h.service.ListarGlosasMensais(filtro)
	response := models.ResponseDefaultModel{IsSuccess: true, Data: glosas}

	if err != nil {
		response.IsSuccess = false
		response.ErrorMessage = "Erro ao listar glosas mensais"
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *glosaMensalHandler) BuscarGlosaMensalPorID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	gm, err := h.service.BuscarGlosaMensalPorID(id)

	if err != nil {
		http.Error(w, "Erro ao buscar glosa", 500)
		return
	}

	if gm == nil {
		http.Error(w, "Glosa não encontrada", 404)
		return
	}

	response := models.ResponseDefaultModel{
		IsSuccess: true,
		Data:      gm,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
