package routes

import (
	dbsql "clinica-api/database"
	controleProtese "clinica-api/internal/ControleProtese"
	"net/http"
)

func RegisterControleProteseRoutes(mux *http.ServeMux, db *dbsql.SQLStr) {
	repo := controleProtese.NovoControleRepository(db)
	svc := controleProtese.NovoControleService(repo)
	handler := controleProtese.NovoControleHandler(svc)

	mux.HandleFunc("/controleProtese", handler.Criar)
	mux.HandleFunc("/editarControleProtese", handler.Editar)
	mux.HandleFunc("/listarControleProtese", handler.Listar)
	mux.HandleFunc("/buscarControleProtese", handler.BuscarPorID)
	mux.HandleFunc("/statusControleProtese", handler.AlterarStatus)
	mux.HandleFunc("/listarFiltradoControleProtese", handler.ListarFiltrado)

}
