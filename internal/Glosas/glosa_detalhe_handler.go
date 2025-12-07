package glosa

import (
    "clinica-api/internal/models"
    "encoding/json"
    "net/http"
    "strconv"
)

type GlosaDetalheHandler interface {
    CriarGlosa(w http.ResponseWriter, r *http.Request)
    EditarGlosa(w http.ResponseWriter, r *http.Request)
    ListarGlosas(w http.ResponseWriter, r *http.Request)
    BuscarPorId(w http.ResponseWriter, r *http.Request)
}

type glosaDetalheHandler struct {
    service GlosaDetalheService
}

func NovoGlosaDetalheHandler(service GlosaDetalheService) GlosaDetalheHandler {
    return &glosaDetalheHandler{service: service}
}

func (h *glosaDetalheHandler) CriarGlosa(w http.ResponseWriter, r *http.Request) {
    var g models.GlosaDetalhe
    response := models.ResponseDefaultModel{IsSuccess: true, Data: g}

    if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
        response.IsSuccess = false
        response.ErrorMessage = "Erro ao decodificar dados"
        w.WriteHeader(http.StatusBadRequest)
    } else if err := h.service.CriarGlosa(&g); err != nil {
        response.IsSuccess = false
        response.ErrorMessage = err.Error()
        w.WriteHeader(http.StatusInternalServerError)
    } else {
        w.WriteHeader(http.StatusCreated)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func (h *glosaDetalheHandler) EditarGlosa(w http.ResponseWriter, r *http.Request) {
    var g models.GlosaDetalhe
    response := models.ResponseDefaultModel{IsSuccess: true, Data: g}

    if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
        response.IsSuccess = false
        response.ErrorMessage = "Erro ao decodificar dados"
        w.WriteHeader(http.StatusBadRequest)
    } else if err := h.service.EditarGlosa(&g); err != nil {
        response.IsSuccess = false
        response.ErrorMessage = err.Error()
        w.WriteHeader(http.StatusInternalServerError)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func (h *glosaDetalheHandler) ListarGlosas(w http.ResponseWriter, r *http.Request) {
    glosas, err := h.service.ListarGlosas()
    response := models.ResponseDefaultModel{IsSuccess: true, Data: glosas}

    if err != nil {
        response.IsSuccess = false
        response.ErrorMessage = "Erro ao listar glosas"
        w.WriteHeader(http.StatusInternalServerError)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func (h *glosaDetalheHandler) BuscarPorId(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
    id, _ := strconv.Atoi(idStr)

    glosa, err := h.service.BuscarPorId(id)
    if err != nil {
        http.Error(w, "Erro ao buscar glosa", 500)
        return
    }

    if glosa == nil {
        http.Error(w, "Glosa não encontrada", 404)
        return
    }

    response := models.ResponseDefaultModel{
        IsSuccess: true,
        Data:      glosa,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}
