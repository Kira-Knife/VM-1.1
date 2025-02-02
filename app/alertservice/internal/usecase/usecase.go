package usecase

import (
	"alertservice/config"
	"alertservice/internal/database/repo"
	"alertservice/pkg/logger"
)

type DBInterf interface {
}

type UseCase struct {
	db         *repo.PostgresRepo
	logger     *logger.Logger
	appJiraURL string
}

func New(cfg *config.Config, db *repo.PostgresRepo, l *logger.Logger) *UseCase {
	return &UseCase{
		db:         db,
		logger:     l,
		appJiraURL: cfg.AppJira.URL,
	}
}
