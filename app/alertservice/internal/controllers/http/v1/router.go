package v1

import (
	"log"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

func (s *Server) routeRegistration() {
	// swagger
	s.router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	apiRouter := s.router.PathPrefix("/api/v1").Subrouter()
	apiRouter.Use(s.enableCORS) // включение CORS заголовков

	apiRouter.HandleFunc("/alerts", s.incomingAlerts).Methods(http.MethodPost) // входящие алерты

	// AlertFrontendApi
	apiRouter.HandleFunc("/alerts", s.getAlerts).Methods(http.MethodGet)                      // -
	apiRouter.HandleFunc("/alerts/{alert_id:[0-9]+}", s.getAlertByID).Methods(http.MethodGet) // -

	apiRouter.HandleFunc("/incidents", s.getIncidents).Methods(http.MethodGet)                                 // -
	apiRouter.HandleFunc("/incidents/{incidents_id:[0-9]+}", s.getIncidentByID).Methods(http.MethodGet)        // -
	apiRouter.HandleFunc("/incidents/{incidents_id:[0-9]+}", s.updateIncidentStatus).Methods(http.MethodPatch) // -

	apiRouter.HandleFunc("/settings/interval", s.handlerPass).Methods(http.MethodGet)  // -
	apiRouter.HandleFunc("/settings/interval", s.handlerPass).Methods(http.MethodPost) // -
}

func (s *Server) handlerPass(w http.ResponseWriter, r *http.Request) {
	// Устанавливаем код состояния 503
	w.WriteHeader(http.StatusServiceUnavailable)

	// Возвращаем сообщение о том, что функция в разработке
	_, err := w.Write([]byte("Этот функционал находится в разработке. Пожалуйста, попробуйте позже."))
	if err != nil {
		// Логируем ошибку, если не удалось записать ответ
		log.Printf("Failed to write response: %v", err)
	}
}
