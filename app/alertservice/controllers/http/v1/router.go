package v1

import (
	"log"
	"net/http"
)

func (s *Server) routeRegistration() {

	// AlertFrontendApi
	s.router.HandleFunc("/alerts", s.handlerPass).Methods(http.MethodGet)
	s.router.HandleFunc("/incidents", s.handlerPass).Methods(http.MethodGet)
	s.router.HandleFunc("/alerts/{alert_id:[0-9]+}", s.handlerPass).Methods(http.MethodGet)
	s.router.HandleFunc("/incidents/{incidents_id:[0-9]+}", s.handlerPass).Methods(http.MethodGet)

	s.router.HandleFunc("/settings/interval", s.handlerPass).Methods(http.MethodGet)
	s.router.HandleFunc("/settings/interval", s.handlerPass).Methods(http.MethodPost)
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
