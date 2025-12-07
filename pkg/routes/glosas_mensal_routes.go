package routes

import (
	dbsql "clinica-api/database"
	glosas "clinica-api/internal/Glosas"
	"net/http"
)

func RegisterGlosasMensalRoutes(mux *http.ServeMux, db *dbsql.SQLStr) {

	repo := glosas.NovoGlosaMensalRepository(db)
	svc := glosas.NovoGlosaMensalService(repo)
	handler := glosas.NovoGlosaMensalHandler(svc)

	mux.HandleFunc("/glosaMensal", handler.CriarGlosaMensal)
	mux.HandleFunc("/editarGlosaMensal", handler.EditarGlosaMensal)
	mux.HandleFunc("/listarGlosaMensal", handler.ListarGlosasMensais)
	mux.HandleFunc("/buscarGlosaMensal", handler.BuscarGlosaMensalPorID)
}
	