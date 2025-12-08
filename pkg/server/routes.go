package server

import (
	"net/http"

	dbsql "clinica-api/database"
	routes "clinica-api/pkg/routes"
)

func SetupRoutes(mux *http.ServeMux, db *dbsql.SQLStr) {
	routes.RegisterAuthRoutes(mux, db)
	routes.RegisterPontoRoutes(mux, db)
	routes.RegisterFuncionarioRoutes(mux, db)
	routes.RegisterUnidadeRoutes(mux, db)
	routes.RegisterPacienteRoutes(mux, db)
	routes.RegisterLeadsRoutes(mux, db)
	routes.RegisterLeadsHistoricoRoutes(mux, db)
	routes.RegisterProcedimentoRoutes(mux, db)
	routes.RegisterGlosasMensalRoutes(mux, db)
	routes.RegisterGlosasDetalhesRoutes(mux, db)

}
