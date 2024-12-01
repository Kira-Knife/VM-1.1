package v1

import (
	"alertservice/config"
	"alertservice/usecase"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	url    string
	u      *usecase.UseCase
	router *mux.Router
}

func New(cfg *config.Config, u *usecase.UseCase) *Server {
	router := mux.NewRouter()

	return &Server{
		u:      u,
		router: router,
		url:    ":" + cfg.HTTP.Port,
	}
}

func (s *Server) Run() {
	http.ListenAndServe(s.url, s.router)
}
