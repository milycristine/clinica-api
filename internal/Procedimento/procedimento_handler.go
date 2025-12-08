package procedimentos

import (
    "clinica-api/internal/models"
    "encoding/json"
    "net/http"
    "strconv"
)

type ProcedimentoHandler interface {
    CriarProcedimento(w http.ResponseWriter, r *http.Request)
    EditarProcedimento(w http.ResponseWriter, r *http.Request)
    ListarProcedimentos(w http.ResponseWriter, r *http.Request)
    BuscarProcedimentoPorID(w http.ResponseWriter, r *http.Request)
    ListarProcedimentosAtivos(w http.ResponseWriter, r *http.Request)

    AssociarProcedimentosAoLead(w http.ResponseWriter, r *http.Request)
    ListarProcedimentosPorLead(w http.ResponseWriter, r *http.Request)
}

type procedimentoHandler struct {
    service ProcedimentoService
}

func NovoProcedimentoHandler(s ProcedimentoService) ProcedimentoHandler {
    return &procedimentoHandler{service: s}
}

func (h *procedimentoHandler) CriarProcedimento(w http.ResponseWriter, r *http.Request) {
    var p models.Procedimento
    resp := models.ResponseDefaultModel{IsSuccess: true, Data: p}

    if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
        resp.IsSuccess = false
        resp.ErrorMessage = "Erro ao decodificar procedimento"
        w.WriteHeader(http.StatusBadRequest)
    } else if err := h.service.CriarProcedimento(&p); err != nil {
        resp.IsSuccess = false
        resp.ErrorMessage = err.Error()
        w.WriteHeader(http.StatusInternalServerError)
    } else {
        w.WriteHeader(http.StatusCreated)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}

func (h *procedimentoHandler) EditarProcedimento(w http.ResponseWriter, r *http.Request) {
    var p models.Procedimento
    resp := models.ResponseDefaultModel{IsSuccess: true, Data: p}

    if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
        resp.IsSuccess = false
        resp.ErrorMessage = "Erro ao decodificar procedimento"
        w.WriteHeader(http.StatusBadRequest)
    } else if err := h.service.EditarProcedimento(&p); err != nil {
        resp.IsSuccess = false
        resp.ErrorMessage = err.Error()
        w.WriteHeader(http.StatusInternalServerError)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}

func (h *procedimentoHandler) ListarProcedimentos(w http.ResponseWriter, r *http.Request) {
    lista, err := h.service.ListarProcedimentos()
    resp := models.ResponseDefaultModel{IsSuccess: err == nil, Data: lista}
    if err != nil {
        resp.ErrorMessage = "Erro ao listar procedimentos"
        w.WriteHeader(http.StatusInternalServerError)
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}

func (h *procedimentoHandler) BuscarProcedimentoPorID(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
    id, _ := strconv.Atoi(idStr)

    p, err := h.service.BuscarProcedimentoPorID(id)
    if err != nil {
        http.Error(w, "Erro ao buscar procedimento", http.StatusInternalServerError)
        return
    }
    if p == nil {
        http.Error(w, "Procedimento não encontrado", http.StatusNotFound)
        return
    }

    resp := models.ResponseDefaultModel{IsSuccess: true, Data: p}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}

func (h *procedimentoHandler) ListarProcedimentosAtivos(w http.ResponseWriter, r *http.Request) {
    lista, err := h.service.ListarProcedimentosAtivos()
    resp := models.ResponseDefaultModel{IsSuccess: err == nil, Data: lista}
    if err != nil {
        resp.ErrorMessage = "Erro ao listar procedimentos ativos"
        w.WriteHeader(http.StatusInternalServerError)
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}

func (h *procedimentoHandler) AssociarProcedimentosAoLead(w http.ResponseWriter, r *http.Request) {
    var body struct {
        ContatoId     int   `json:"contato_id"`
        Procedimentos []int `json:"procedimentos"`
    }
    resp := models.ResponseDefaultModel{IsSuccess: true, Data: body}

    if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
        resp.IsSuccess = false
        resp.ErrorMessage = "Erro ao decodificar dados"
        w.WriteHeader(http.StatusBadRequest)
    } else if err := h.service.AssociarProcedimentosAoLead(body.ContatoId, body.Procedimentos); err != nil {
        resp.IsSuccess = false
        resp.ErrorMessage = err.Error()
        w.WriteHeader(http.StatusInternalServerError)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}

func (h *procedimentoHandler) ListarProcedimentosPorLead(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("contatoId")
    contatoId, _ := strconv.Atoi(idStr)

    lista, err := h.service.ListarProcedimentosPorLead(contatoId)
    resp := models.ResponseDefaultModel{IsSuccess: err == nil, Data: lista}
    if err != nil {
        resp.ErrorMessage = "Erro ao listar procedimentos do lead"
        w.WriteHeader(http.StatusInternalServerError)
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}
