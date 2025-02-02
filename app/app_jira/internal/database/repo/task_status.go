package repo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
)

// GetTaskStatusIDByName - получение идентификатора статуса задачи по имени
func (r *PostgresRepo) GetTaskStatusIDByName(ctx context.Context, statusName string) (int, error) {
	var statusID int

	sqlq, args, err := r.db.Builder.Select("id").From("task_status").Where(squirrel.Eq{"name": statusName}).ToSql()
	if err != nil {
		return 0, fmt.Errorf("PostgresRepo - GetTaskStatusIDByName - r.db.Builder.Select: %w", err)
	}

	row := r.db.Pool.QueryRow(ctx, sqlq, args...)
	err = row.Scan(&statusID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("status not found")
		}
		return 0, fmt.Errorf("PostgresRepo - GetTaskStatusIDByName - QueryRowContext: %w", err)
	}

	return statusID, nil
}
