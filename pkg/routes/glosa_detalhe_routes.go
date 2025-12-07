package routes

import (
	dbsql "clinica-api/database"
	glosas "clinica-api/internal/Glosas"
	"net/http"
)

func RegisterGlosasDetalhesRoutes(mux *http.ServeMux, db *dbsql.SQLStr) {

	repo := glosas.NovoGlosaDetalheRepository(db)
	svc := glosas.NovoGlosaDetalheService(repo)
	handler := glosas.NovoGlosaDetalheHandler(svc)

	mux.HandleFunc("/glosaDetalhe", handler.CriarGlosa)
	mux.HandleFunc("/editarGlosaDetalhe", handler.EditarGlosa)
	mux.HandleFunc("/listarGlosaDetalhe", handler.ListarGlosas)
	mux.HandleFunc("/buscarGlosaDetalhe", handler.BuscarPorId)

}
