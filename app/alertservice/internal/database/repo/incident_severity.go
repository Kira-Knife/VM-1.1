package repo

import (
	"alertservice/internal/entity"
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
)

// GetSeverity - получение уровня серьезности по ID
func (r *PostgresRepo) GetSeverity(ctx context.Context, severityID int64) (entity.Severity, error) {
	sql, args, err := r.db.Builder.
		Select("id, name, priority").
		From("severities").
		Where(squirrel.Eq{"id": severityID}).
		ToSql()
	if err != nil {
		return entity.Severity{}, fmt.Errorf("SeverityRepo - GetSeverity - r.Builder: %w", err)
	}

	row := r.db.Pool.QueryRow(ctx, sql, args...)

	var severity entity.Severity
	err = row.Scan(
		&severity.ID,
		&severity.Name,
		&severity.Priority,
	)
	if err != nil {
		return entity.Severity{}, fmt.Errorf("SeverityRepo - GetSeverity - row.Scan: %w", err)
	}

	return severity, nil
}

// GetSeverities - получение всех уровней серьезности
func (r *PostgresRepo) GetSeverities(ctx context.Context) ([]entity.Severity, error) {
	sql, args, err := r.db.Builder.
		Select("id, name, priority").
		From("severities").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("SeverityRepo - GetSeverities - r.Builder: %w", err)
	}

	rows, err := r.db.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("SeverityRepo - GetSeverities - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	severities := make([]entity.Severity, 0)

	for rows.Next() {
		var severity entity.Severity
		err = rows.Scan(
			&severity.ID,
			&severity.Name,
			&severity.Priority,
		)
		if err != nil {
			return nil, fmt.Errorf("SeverityRepo - GetSeverities - rows.Scan: %w", err)
		}
		severities = append(severities, severity)
	}

	return severities, nil
}
