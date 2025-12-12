package routes

import (
	dbsql "clinica-api/database"
	fechamento "clinica-api/internal/FechamentoLaboratorio"
	"net/http"
)

func RegisterFechamentoRoutes(mux *http.ServeMux, db *dbsql.SQLStr) {
	repo := fechamento.NovoFechamentoRepository(db)
	svc := fechamento.NovoFechamentoService(repo)
	handler := fechamento.NovoFechamentoHandler(svc)

	mux.HandleFunc("/fechamento", handler.Criar)
	mux.HandleFunc("/editarFechamento", handler.Editar)
	mux.HandleFunc("/listarFechamentos", handler.Listar)
	mux.HandleFunc("/buscarFechamentoPorId", handler.BuscarPorID)
	mux.HandleFunc("/listarFechamentosPorLaboratorio", handler.ListarPorLaboratorio)
	mux.HandleFunc("/listarFechamentosPorData", handler.ListarPorData)
	mux.HandleFunc("/listarFechamentosPorMes", handler.ListarPorMes)
	mux.HandleFunc("/listarFechamentosPeriodo", handler.ListarPorPeriodo)
	mux.HandleFunc("/listarFechamentosPorLaboratorioMes", handler.ListarPorLaboratorioMes)
	mux.HandleFunc("/listarFechamentosPorPaciente", handler.ListarPorPaciente)

}
