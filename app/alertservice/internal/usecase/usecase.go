package usecase

import (
	"alertservice/config"
	"alertservice/internal/database/repo"
	"alertservice/internal/entity"
	"alertservice/pkg/logger"
	"context"
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

func (u *UseCase) StoreAlert(ctx context.Context, alert entity.Alert) error {
	// alertId, err := u.db.StoreAlert(ctx, alert)
	_, err := u.db.StoreAlert(ctx, alert)
	// create Incident
	return err
}
