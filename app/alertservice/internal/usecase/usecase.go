package usecase

import (
	"alertservice/config"
	"alertservice/internal/database/repo"
	appjira "alertservice/internal/usecase/app_jira"
	"alertservice/pkg/logger"
)

type DBInterf interface {
}

type UseCase struct {
	db         *repo.PostgresRepo
	logger     *logger.Logger
	appJiraApi *appjira.AppJiraAPI
}

func New(cfg *config.Config, db *repo.PostgresRepo, l *logger.Logger) *UseCase {
	appJiraApi := appjira.New(cfg.AppJira.URL)
	return &UseCase{
		db:         db,
		logger:     l,
		appJiraApi: appJiraApi,
	}
}
