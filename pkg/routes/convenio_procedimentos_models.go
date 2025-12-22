package routes

import (
	dbsql "clinica-api/database"
	cp "clinica-api/internal/Convenio"
	"net/http"
)

func RegisterConvenioProcedimentoRoutes(mux *http.ServeMux, db *dbsql.SQLStr) {
	repo := cp.NovoConvenioProcedimentoRepository(db)
	service := cp.NovoConvenioProcedimentoService(repo)
	handler := cp.NovoConvenioProcedimentoHandler(service)

	mux.HandleFunc("/convenioProcedimento", handler.Criar)
	mux.HandleFunc("/listarConvenioProcedimento", handler.ListarPorConvenio)

}
