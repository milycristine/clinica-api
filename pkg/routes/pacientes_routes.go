package routes

import (
	dbsql "clinica-api/database"
	paciente "clinica-api/internal/Paciente"
	"net/http"
)

func RegisterPacienteRoutes(mux *http.ServeMux, db *dbsql.SQLStr) {

	repo := paciente.NovoPacienteRepository(db)
	svc := paciente.NovoPacienteService(repo)
	handler := paciente.NovoPacienteHandler(svc)

	mux.HandleFunc("/paciente", handler.CriarPaciente)
	mux.HandleFunc("/editarPaciente", handler.EditarPaciente)
	mux.HandleFunc("/listarPacientes", handler.ListarPacientes)
	mux.HandleFunc("/buscarPaciente", handler.BuscarPacientePorID)
	mux.HandleFunc("/listarPacientesPorUnidade", handler.ListarPorUnidade)

}
