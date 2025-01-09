package repo

import (
	"context"
	"fmt"

	"alertservice/internal/entity"
)

// GetAllIncidentStates retrieves all alert states from the database.
func (r *PostgresRepo) GetAllIncidentStates(ctx context.Context) ([]entity.IncidentState, error) {
	sql, args, err := r.db.Builder.
		Select("id", "name", "description").
		From("alert_states").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("AlertRepo - GetAllIncidentStates - r.Builder: %w", err)
	}

	rows, err := r.db.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("AlertRepo - GetAllIncidentStates - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	var IncidentStates []entity.IncidentState
	for rows.Next() {
		var IncidentState entity.IncidentState
		if err := rows.Scan(&IncidentState.ID, &IncidentState.Name, &IncidentState.Description); err != nil {
			return nil, fmt.Errorf("AlertRepo - GetAllIncidentStates - rows.Scan: %w", err)
		}
		IncidentStates = append(IncidentStates, IncidentState)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("AlertRepo - GetAllIncidentStates - rows.Err: %w", err)
	}

	return IncidentStates, nil
}
