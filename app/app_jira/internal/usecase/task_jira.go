package usecase

import (
	"app_jira/internal/entity"
	"context"

	"github.com/google/uuid"
)

func (u *UseCase) CreateTaskJira(ctx context.Context, task entity.TaskJira) (uuid.UUID, error) {
	uuID := uuid.New()
	return uuID, nil
}

func (u *UseCase) GetAllTaskJira(ctx context.Context) ([]entity.TaskJira, error) {
	return nil, nil
}

func (u *UseCase) GetListTaskJira(ctx context.Context, begin, count int64) ([]entity.TaskJira, error) {
	return nil, nil
}

func (u *UseCase) GetTaskJiraByUUID(ctx context.Context, uuID uuid.UUID) (entity.TaskJira, error) {
	task := entity.TaskJira{}
	return task, nil
}

func (u *UseCase) UpdateTaskJiraStatus(ctx context.Context, uuID uuid.UUID, newStatus string) error {
	return nil
}
