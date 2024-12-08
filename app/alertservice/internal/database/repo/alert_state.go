package repo

import (
	"context"
	"fmt"

	"alertservice/internal/entity"
)

// GetAllAlertStates retrieves all alert states from the database.
func (r *PostgresRepo) GetAllAlertStates(ctx context.Context) ([]entity.AlertState, error) {
	sql, args, err := r.db.Builder.
		Select("id", "name", "description").
		From("alert_states").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("AlertRepo - GetAllAlertStates - r.Builder: %w", err)
	}

	rows, err := r.db.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("AlertRepo - GetAllAlertStates - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	var alertStates []entity.AlertState
	for rows.Next() {
		var alertState entity.AlertState
		if err := rows.Scan(&alertState.ID, &alertState.Name, &alertState.Description); err != nil {
			return nil, fmt.Errorf("AlertRepo - GetAllAlertStates - rows.Scan: %w", err)
		}
		alertStates = append(alertStates, alertState)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("AlertRepo - GetAllAlertStates - rows.Err: %w", err)
	}

	return alertStates, nil
}
