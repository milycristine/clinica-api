package routes

import (
	dbsql "clinica-api/database"
	despesa "clinica-api/internal/Despesas"
	"net/http"
)

func RegisterDespesaRoutes(mux *http.ServeMux, db *dbsql.SQLStr) {
	repo := despesa.NovaDespesaRepository(db)
	svc := despesa.NovaDespesaService(repo)
	handler := despesa.NovaDespesaHandler(svc)

	mux.HandleFunc("/despesa", handler.CriarDespesa)
	mux.HandleFunc("/listarDespesas", handler.ListarDespesas)
	mux.HandleFunc("/buscarDespesa", handler.BuscarDespesaPorID)
	mux.HandleFunc("/editarDespesa", handler.EditarDespesa)
	mux.HandleFunc("/statusDespesa", handler.AtualizarStatus)	
}
