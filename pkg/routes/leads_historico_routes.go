package routes

import (
	dbsql "clinica-api/database"
	leads "clinica-api/internal/Leads"
	"net/http"
)

func RegisterLeadsHistoricoRoutes(mux *http.ServeMux, db *dbsql.SQLStr) {

	repo := leads.NovoLeadHistoricoRepository(db)
	svc := leads.NovoLeadHistoricoService(repo)
	handler := leads.NovoLeadHistoricoHandler(svc)

	mux.HandleFunc("/historico", handler.CriarHistorico)
	mux.HandleFunc("/listarHistorico", handler.ListarHistoricoPorLead)
}
