package routes

import (
	dbsql "clinica-api/database"
	"clinica-api/internal/Repasse"
	"net/http"
)

func RegisterRepasseRoutes(mux *http.ServeMux, db *dbsql.SQLStr) {
	repo := repasse.NovoRepasseRepository(db)
	svc := repasse.NovoRepasseService(repo)
	handler := repasse.NovoRepasseHandler(svc)

	mux.HandleFunc("/repasse", handler.Criar)
	mux.HandleFunc("/listarRepasses", handler.Listar)
	mux.HandleFunc("/pagarRepasse", handler.Pagar)
}
