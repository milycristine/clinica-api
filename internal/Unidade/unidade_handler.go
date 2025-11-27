package unidade

import (
    "clinica-api/internal/models"
    "encoding/json"
    "net/http"
    "strconv"
)

type UnidadeHandler interface {
    CriarUnidade(w http.ResponseWriter, r *http.Request)
    EditarUnidade(w http.ResponseWriter, r *http.Request)
    ListarUnidades(w http.ResponseWriter, r *http.Request)
    BuscarUnidadePorID(w http.ResponseWriter, r *http.Request)
    AtualizarStatus(w http.ResponseWriter, r *http.Request)
}

type unidadeHandler struct {
    service UnidadeService
}

func NovaUnidadeHandler(service UnidadeService) UnidadeHandler {
    return &unidadeHandler{service: service}
}

func (h *unidadeHandler) CriarUnidade(w http.ResponseWriter, r *http.Request) {
    var unidade models.Unidade
    response := models.ResponseDefaultModel{IsSuccess: true, Data: unidade}

    if err := json.NewDecoder(r.Body).Decode(&unidade); err != nil {
        response.IsSuccess = false
        response.ErrorMessage = "Erro ao decodificar dados"
        w.WriteHeader(http.StatusBadRequest)
    } else if err := h.service.CriarUnidade(&unidade); err != nil {
        response.IsSuccess = false
        response.ErrorMessage = err.Error()
        w.WriteHeader(http.StatusInternalServerError)
    } else {
        w.WriteHeader(http.StatusCreated)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func (h *unidadeHandler) EditarUnidade(w http.ResponseWriter, r *http.Request) {
    var unidade models.Unidade
    response := models.ResponseDefaultModel{IsSuccess: true, Data: unidade}

    if err := json.NewDecoder(r.Body).Decode(&unidade); err != nil {
        response.IsSuccess = false
        response.ErrorMessage = "Erro ao decodificar dados"
        w.WriteHeader(http.StatusBadRequest)
    } else if err := h.service.EditarUnidade(&unidade); err != nil {
        response.IsSuccess = false
        response.ErrorMessage = err.Error()
        w.WriteHeader(http.StatusInternalServerError)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func (h *unidadeHandler) ListarUnidades(w http.ResponseWriter, r *http.Request) {
    unidades, err := h.service.ListarUnidades()
    response := models.ResponseDefaultModel{IsSuccess: true, Data: unidades}

    if err != nil {
        response.IsSuccess = false
        response.ErrorMessage = "Erro ao listar unidades"
        w.WriteHeader(http.StatusInternalServerError)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func (h *unidadeHandler) BuscarUnidadePorID(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
    id, _ := strconv.Atoi(idStr)

    unidade, err := h.service.BuscarUnidadePorID(id)
    if err != nil {
        http.Error(w, "Erro ao buscar unidade", 500)
        return
    }

    if unidade == nil {
        http.Error(w, "Unidade não encontrada", 404)
        return
    }

    response := models.ResponseDefaultModel{
        IsSuccess: true,
        Data:      unidade,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func (h *unidadeHandler) AtualizarStatus(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
    statusStr := r.URL.Query().Get("status")

    id, _ := strconv.Atoi(idStr)
    status := statusStr == "true"

    err := h.service.AtualizarStatus(id, status)
    response := models.ResponseDefaultModel{IsSuccess: err == nil}

    if err != nil {
        response.ErrorMessage = "Erro ao atualizar status"
        w.WriteHeader(http.StatusInternalServerError)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}
