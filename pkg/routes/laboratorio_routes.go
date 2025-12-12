package routes

import (
    dbsql "clinica-api/database"
    laboratorio "clinica-api/internal/Laboratorio"
    "net/http"
)

func RegisterLaboratorioRoutes(mux *http.ServeMux, db *dbsql.SQLStr) {
    repo := laboratorio.NovoLaboratorioRepository(db)
    svc := laboratorio.NovoLaboratorioService(repo)
    handler := laboratorio.NovoLaboratorioHandler(svc)

    mux.HandleFunc("/laboratorio", handler.CriarLaboratorio)
    mux.HandleFunc("/editarLaboratorio", handler.EditarLaboratorio)
    mux.HandleFunc("/listarLaboratorio", handler.ListarLaboratorios)
    mux.HandleFunc("/buscarLaboratorioPorId", handler.BuscarLaboratorioPorID)
	
}
