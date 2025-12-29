package routes

import (
	dbsql "clinica-api/database"
	"clinica-api/internal/AjustesCaixa"
	"net/http"
)

func RegisterAjusteCaixaRoutes(mux *http.ServeMux, db *dbsql.SQLStr) {
	repo := ajustecaixa.NovoAjusteCaixaRepository(db)
	svc := ajustecaixa.NovoAjusteCaixaService(repo)
	handler := ajustecaixa.NovoAjusteCaixaHandler(svc)

	mux.HandleFunc("/ajusteCaixa", handler.Criar)
	mux.HandleFunc("/listarAjustesCaixa", handler.Listar)
}
