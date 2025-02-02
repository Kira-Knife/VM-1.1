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
	// apiRouter.Use(s.enableCORS) // включение CORS заголовков

	apiRouter.HandleFunc("/task", s.createTaskJira).Methods(http.MethodPost, http.MethodOptions)

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
