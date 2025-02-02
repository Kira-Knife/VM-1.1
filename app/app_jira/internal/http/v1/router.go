package v1

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

func (s *Server) routeRegistration() {
	// swagger
	s.router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	apiRouter := s.router.PathPrefix("/api/v1").Subrouter()
	// apiRouter.Use(s.enableCORS) // включение CORS заголовков

	apiRouter.HandleFunc("/task", s.createTaskJira).Methods(http.MethodPost, http.MethodOptions)
	apiRouter.HandleFunc("/task", s.getAllTaskJira).Methods(http.MethodGet, http.MethodOptions)
	apiRouter.HandleFunc("/task/list", s.getListTasksJira).Methods(http.MethodGet, http.MethodOptions)
	apiRouter.HandleFunc("/task/{task_uuid:[0-9a-fA-F-]{36}}", s.getTaskJiraByID).Methods(http.MethodGet, http.MethodOptions)
	apiRouter.HandleFunc("/task/{task_uuid:[0-9a-fA-F-]{36}}", s.updateTaskStatus).Methods(http.MethodPatch, http.MethodOptions)
}
