package server

import (
	"net/http"

	dbsql "clinica-api/database"
	routes "clinica-api/pkg/routes"
)


func SetupRoutes(mux *http.ServeMux, db *dbsql.SQLStr) {
	routes.RegisterAuthRoutes(mux, db)
}
