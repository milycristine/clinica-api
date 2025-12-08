package routes

import (
	dbsql "clinica-api/database"
	leads "clinica-api/internal/Leads"
	procedimento "clinica-api/internal/Procedimento"
	"net/http"
)

func RegisterLeadsRoutes(mux *http.ServeMux, db *dbsql.SQLStr) {
	contatoRepo := leads.NovoContatoRepository(db)
	historicoRepo := leads.NovoLeadHistoricoRepository(db)
	procedimentoRepo := procedimento.NovoProcedimentoRepository(db)
	service := leads.NovoContatoService(contatoRepo, historicoRepo, procedimentoRepo)
	handler := leads.NovoContatoHandler(service)

	mux.HandleFunc("/leads", handler.CriarContato)
	mux.HandleFunc("/editarLead", handler.EditarContato)
	mux.HandleFunc("/listarLead", handler.ListarContatos)
	mux.HandleFunc("/buscarLead", handler.BuscarContatoPorID)
	mux.HandleFunc("/atualizarStatusLead", handler.AtualizarStatus)

}
