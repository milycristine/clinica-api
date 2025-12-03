package routes

import (
	dbsql "clinica-api/database"
	leads "clinica-api/internal/Leads"
	"net/http"
)

func RegisterLeadsRoutes(mux *http.ServeMux, db *dbsql.SQLStr) {

	repo := leads.NovoContatoRepository(db)
	svc := leads.NovoContatoService(repo)
	handler := leads.NovoContatoHandler(svc)

	mux.HandleFunc("/leads", handler.CriarContato)
	mux.HandleFunc("/editarLead", handler.EditarContato)
	mux.HandleFunc("/listarLead", handler.ListarContatos)
	mux.HandleFunc("/buscarLead", handler.BuscarContatoPorID)
	mux.HandleFunc("/buscarLeadPorId", handler.AtualizarStatus)
}
