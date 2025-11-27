package routes

import (
	dbsql "clinica-api/database"
	unidade "clinica-api/internal/Unidade"
	"net/http"
)

func RegisterUnidadeRoutes(mux *http.ServeMux, db *dbsql.SQLStr) {

	repo := unidade.NovaUnidadeRepository(db)
	svc := unidade.NovoUnidadeService(repo)
	handler := unidade.NovaUnidadeHandler(svc)

	mux.HandleFunc("/unidade", handler.CriarUnidade)
	mux.HandleFunc("/editarUnidade", handler.EditarUnidade)
	mux.HandleFunc("/listarUnidade", handler.ListarUnidades)
	mux.HandleFunc("/buscarUnidade", handler.BuscarUnidadePorID)
	mux.HandleFunc("/buscarUnidadePorId", handler.AtualizarStatus)
}
