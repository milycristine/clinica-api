package routes

import (
	dbsql "clinica-api/database"
	receita "clinica-api/internal/Receita"
	"net/http"
)

func RegisterReceitaRoutes(mux *http.ServeMux, db *dbsql.SQLStr) {
	repo := receita.NovaReceitaRepository(db)
	svc := receita.NovoReceitaService(repo)
	handler := receita.NovaReceitaHandler(svc)

	mux.HandleFunc("/receita", handler.CriarReceita)
	mux.HandleFunc("/editarReceita", handler.EditarReceita)
	mux.HandleFunc("/listarReceitas", handler.ListarReceitas)
	mux.HandleFunc("/buscarReceita", handler.BuscarReceitaPorID)
	mux.HandleFunc("/atualizarStatusReceita", handler.AtualizarStatus)
}
