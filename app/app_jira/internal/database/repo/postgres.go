package repo

import (
	"app_jira/pkg/logger"
	"app_jira/pkg/postgres"
)

const _defaultEntityCap = 64

// TranslationRepo -.
type PostgresRepo struct {
	db     *postgres.Postgres
	logger *logger.Logger
}

// New -.
func New(pg *postgres.Postgres, l *logger.Logger) *PostgresRepo {
	return &PostgresRepo{
		db:     pg,
		logger: l,
	}
}
