package usecase

import (
	"app_jira/config"
	"app_jira/internal/database/repo"
	"app_jira/pkg/logger"
)

type UseCase struct {
	db     *repo.PostgresRepo
	logger *logger.Logger
}

func New(cfg *config.Config, db *repo.PostgresRepo, l *logger.Logger) *UseCase {
	return &UseCase{
		db:     db,
		logger: l,
	}
}
