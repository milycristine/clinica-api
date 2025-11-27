package server

import (
	"embed"
	"fmt"
	"log"
	"net/http"

	dbsql "clinica-api/database" 
)

func StartServer(port string, db *dbsql.SQLStr, assets embed.FS) error {
	mux := http.NewServeMux()

	SetupRoutes(mux, db)

	mux.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./public/images"))))

	if assets != (embed.FS{}) {
		mux.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.FS(assets))))
	}

	
	addr := fmt.Sprintf(":%s", port)
	log.Printf("Iniciando servidor na porta %s ", addr)
	return http.ListenAndServe(addr, mux)
}
