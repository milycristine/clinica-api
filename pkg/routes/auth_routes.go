package routes

import (
	dbsql "clinica-api/database"
	auth "clinica-api/internal/Auth"
	"net/http"
)

func RegisterAuthRoutes(mux *http.ServeMux, db *dbsql.SQLStr) {
	repo := auth.NewAuthRepository(db)
	svc := auth.NewAuthService(repo)

	mux.HandleFunc("/login", auth.LoginHandler(svc))
}