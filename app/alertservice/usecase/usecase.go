package usecase

import (
	"alertservice/config"
	"alertservice/database/repo"
	"alertservice/entity"
	"context"
)

type UseCase struct {
	db *repo.TranslationRepo
}

func New(cfg *config.Config, db *repo.TranslationRepo) *UseCase {
	return &UseCase{
		db: db,
	}
}

func (u *UseCase) StoreAlert(ctx context.Context, alert entity.Alert) error {
	err := u.db.StoreAlert(ctx, alert)
	return err
}
