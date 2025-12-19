package routes

import (
	dbsql "clinica-api/database"
	convenio "clinica-api/internal/Convenio"
	glosa "clinica-api/internal/Glosas"
	"net/http"
)

func RegisterConvenioRoutes(mux *http.ServeMux, db *dbsql.SQLStr) {

	convenioRepo := convenio.NovoConvenioRepository(db)
	convenioProcRepo := convenio.NovoConvenioProcedimentoRepository(db)

	glosaRepo := glosa.NovoGlosaDetalheRepository(db)
	glosaMensalRepo := glosa.NovoGlosaMensalRepository(db)
	glosaService := glosa.NovoGlosaDetalheService(glosaRepo, glosaMensalRepo)

	convenioService := convenio.NovoConvenioService(
		convenioRepo,
		convenioProcRepo,
		glosaService,
	)
	handler := convenio.NovoConvenioHandler(convenioService)

	mux.HandleFunc("/convenio", handler.CriarConvenio)
	mux.HandleFunc("/editarConvenio", handler.EditarConvenio)
	mux.HandleFunc("/listarConvenios", handler.ListarConvenios)
	mux.HandleFunc("/buscarConvenio", handler.BuscarConvenioPorID)
	mux.HandleFunc("/atualizarStatusConvenio", handler.AtualizarStatus)

}
