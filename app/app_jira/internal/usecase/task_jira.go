package usecase

import (
	"app_jira/internal/entity"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (u *UseCase) CreateTaskJira(ctx context.Context, taskReq entity.TaskJiraRequest) (uuid.UUID, error) {
	uuID, err := u.db.CreateTaskJira(ctx, taskReq)
	return uuID, err
}

func (u *UseCase) GetAllTaskJira(ctx context.Context) ([]entity.TaskJira, error) {
	return u.db.GetAllTaskJira(ctx)
}

func (u *UseCase) GetListTaskJira(ctx context.Context, begin, count int64) ([]entity.TaskJira, error) {
	err := fmt.Errorf("находится в разработке")
	return nil, err
}

func (u *UseCase) GetTaskJiraByUUID(ctx context.Context, uuID uuid.UUID) (entity.TaskJira, error) {
	return u.db.GetTaskJiraByUUID(ctx, uuID)
}

func (u *UseCase) UpdateTaskJiraStatus(ctx context.Context, uuID uuid.UUID, newStatus string) error {
	statusId, err := u.db.GetTaskStatusIDByName(ctx, newStatus)
	if err != nil {
		return fmt.Errorf("UseCase - UpdateTaskJiraStatus: неверно указан статус")
		//return fmt.Errorf("UseCase - UpdateTaskJiraStatus: %w", err)
	}
	return u.db.UpdateTaskJiraStatus(ctx, uuID, statusId)
}
