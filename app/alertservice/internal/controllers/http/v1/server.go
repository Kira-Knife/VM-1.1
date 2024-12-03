package v1

import (
	"alertservice/config"
	"alertservice/internal/usecase"
	"alertservice/pkg/logger"
	"context"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	url        string
	u          *usecase.UseCase
	router     *mux.Router
	httpServer *http.Server
	logger     *logger.Logger
}

func New(cfg *config.Config, u *usecase.UseCase, l *logger.Logger) *Server {
	router := mux.NewRouter()

	s := Server{
		u:      u,
		router: router,
		url:    ":" + cfg.HTTP.Port,
		httpServer: &http.Server{
			Addr:    ":" + cfg.HTTP.Port,
			Handler: router,
		},
		logger: l,
	}
	s.routeRegistration()
	return &s
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	// Устанавливаем таймаут для завершения работы сервера
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Останавливаем сервер
	if err := s.httpServer.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
