package repo

import (
	"app_jira/internal/entity"
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

// CreateTaskJira - создание новой задачи в базе данных и генерация уникального task_main_id
func (r *PostgresRepo) CreateTaskJira(ctx context.Context, task entity.TaskJiraRequest) (uuid.UUID, error) {
	var newTaskUUID uuid.UUID
	var err error

	// Повторяющаяся генерация UUID до получения уникального
	for {
		newTaskUUID = uuid.New()

		// Проверка на уникальность task_main_id в таблице
		var exists bool
		err = r.db.Pool.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM tasks WHERE task_main_id = $1)", newTaskUUID).Scan(&exists)
		if err != nil {
			return uuid.UUID{}, fmt.Errorf("PostgresRepo - CreateTaskJira - QueryRow (check exists): %w", err)
		}

		if !exists {
			break
		}
	}

	statusId, err := r.GetTaskStatusIDByName(ctx, task.Status)
	if err != nil {
		statusId = 1
	}

	sql, args, err := r.db.Builder.
		Insert("tasks").
		Columns("task_main_id", "task_title", "task_status_id", "assigned", "owner", "incident_id").
		Values(newTaskUUID, task.TaskTitle, statusId,
			task.Assigned, task.Owner, task.IncidentId).
		ToSql()
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("PostgresRepo - CreateTaskJira - r.Builder: %w", err)
	}

	_, err = r.db.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("PostgresRepo - CreateTaskJira - r.Pool.Exec: %w", err)
	}

	return newTaskUUID, nil
}

// GetAllTaskJira - получение всех задач
func (r *PostgresRepo) GetAllTaskJira(ctx context.Context) ([]entity.TaskJira, error) {
	sql, args, err := r.db.Builder.
		Select("t.task_id, t.task_main_id, t.task_title, ts.name, t.assigned, t.owner, t.incident_id").
		From("tasks t").
		Join("task_status ts ON t.task_status_id = ts.id").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("PostgresRepo - GetAllTaskJira - r.Builder: %w", err)
	}

	rows, err := r.db.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("PostgresRepo - GetAllTaskJira - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	tasks := make([]entity.TaskJira, 0)

	for rows.Next() {
		var task entity.TaskJira
		err = rows.Scan(&task.TaskId, &task.JiraTaskId, &task.TaskTitle, &task.Status, &task.Assigned, &task.Owner, &task.IncidentId)
		if err != nil {
			return nil, fmt.Errorf("PostgresRepo - GetAllTaskJira - rows.Scan: %w", err)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// GetTaskJiraByUUID - получение задачи по UUID
func (r *PostgresRepo) GetTaskJiraByUUID(ctx context.Context, jiraTaskId uuid.UUID) (entity.TaskJira, error) {
	sqlq, args, err := r.db.Builder.
		Select("t.task_id, t.task_main_id, t.task_title, ts.name, t.assigned, t.owner, t.incident_id").
		From("tasks t").
		Join("task_status ts ON t.task_status_id = ts.id").
		Where("t.task_main_id = ?", jiraTaskId).
		ToSql()
	if err != nil {
		return entity.TaskJira{}, fmt.Errorf("PostgresRepo - GetTaskJiraByUUID - r.Builder: %w", err)
	}

	var task entity.TaskJira
	err = r.db.Pool.QueryRow(ctx, sqlq, args...).Scan(&task.TaskId, &task.JiraTaskId, &task.TaskTitle, &task.Status, &task.Assigned, &task.Owner, &task.IncidentId)
	if err != nil {
		if err == sql.ErrNoRows {
			return task, fmt.Errorf("task not found")
		}
		return task, fmt.Errorf("PostgresRepo - GetTaskJiraByUUID - r.Pool.QueryRow: %w", err)
	}
	return task, nil
}

// UpdateTaskJiraStatus - обновление статуса задачи
func (r *PostgresRepo) UpdateTaskJiraStatus(ctx context.Context, jiraTaskId uuid.UUID, newStatusID int) error {

	sql, args, err := r.db.Builder.
		Update("tasks").
		Set("task_status_id", newStatusID).
		Where("task_main_id = ?", jiraTaskId).
		ToSql()
	if err != nil {
		return fmt.Errorf("PostgresRepo - UpdateTaskJiraStatus - r.Builder: %w", err)
	}

	_, err = r.db.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("PostgresRepo - UpdateTaskJiraStatus - r.Pool.Exec: %w", err)
	}
	return nil
}
