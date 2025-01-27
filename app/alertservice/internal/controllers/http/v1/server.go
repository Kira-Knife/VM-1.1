package v1

import (
	"alertservice/config"
	"alertservice/internal/usecase"
	"alertservice/pkg/logger"
	"context"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Server struct {
	url        string
	u          *usecase.UseCase
	router     *mux.Router
	httpServer *http.Server
	logger     *logger.Logger
	certFile   string
	keyFile    string
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
		logger:   l,
		certFile: cfg.HTTP.CertFile,
		keyFile:  cfg.HTTP.KeyFile,
	}
	s.routeRegistration()
	return &s
}

func (s *Server) Run() error {
	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow all origins
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})
	handler := corsOptions.Handler(s.router)
	s.httpServer.Handler = handler

	s.logger.Info("Сервер запущен : %s", s.url)

	// return s.httpServer.ListenAndServeTLS(s.certFile, s.keyFile) // https
	return s.httpServer.ListenAndServe() // http
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
