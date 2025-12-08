package routes

import (
	dbsql "clinica-api/database"
	procedimentos "clinica-api/internal/Procedimento"
	"net/http"
)

func RegisterProcedimentoRoutes(mux *http.ServeMux, db *dbsql.SQLStr) {
	procRepo := procedimentos.NovoProcedimentoRepository(db)
	procSvc := procedimentos.NovoProcedimentoService(procRepo)
	procHandler := procedimentos.NovoProcedimentoHandler(procSvc)

	mux.HandleFunc("/procedimentos", procHandler.ListarProcedimentos)              
	mux.HandleFunc("/procedimento", procHandler.CriarProcedimento)                 
	mux.HandleFunc("/procedimentoEditar", procHandler.EditarProcedimento)         
	mux.HandleFunc("/procedimentoBuscar", procHandler.BuscarProcedimentoPorID)    
	mux.HandleFunc("/procedimentosAtivos", procHandler.ListarProcedimentosAtivos) 

	mux.HandleFunc("/procedimentosAssociar", procHandler.AssociarProcedimentosAoLead) 
	mux.HandleFunc("/procedimentosLead", procHandler.ListarProcedimentosPorLead)      

}
