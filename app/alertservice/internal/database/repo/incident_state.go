package repo

import (
	"alertservice/internal/entity"
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
)

// GetIncidentState - получение состояния инцидента по ID
func (r *PostgresRepo) GetIncidentState(ctx context.Context, stateID int) (entity.IncidentState, error) {
	sql, args, err := r.db.Builder.
		Select("id, name").
		From("incident_states").
		Where(squirrel.Eq{"id": stateID}).
		ToSql()
	if err != nil {
		return entity.IncidentState{}, fmt.Errorf("IncidentStateRepo - GetIncidentState - r.Builder: %w", err)
	}

	row := r.db.Pool.QueryRow(ctx, sql, args...)

	var state entity.IncidentState
	err = row.Scan(
		&state.ID,
		&state.Name,
	)
	if err != nil {
		return entity.IncidentState{}, fmt.Errorf("IncidentStateRepo - GetIncidentState - row.Scan: %w", err)
	}

	return state, nil
}

// GetIncidentStates - получение всех состояний инцидентов
func (r *PostgresRepo) GetIncidentStates(ctx context.Context) ([]entity.IncidentState, error) {
	sql, args, err := r.db.Builder.
		Select("id, name").
		From("incident_states").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("IncidentStateRepo - GetIncidentStates - r.Builder: %w", err)
	}

	rows, err := r.db.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("IncidentStateRepo - GetIncidentStates - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	states := make([]entity.IncidentState, 0)

	for rows.Next() {
		var state entity.IncidentState
		err = rows.Scan(
			&state.ID,
			&state.Name,
		)
		if err != nil {
			return nil, fmt.Errorf("IncidentStateRepo - GetIncidentStates - rows.Scan: %w", err)
		}
		states = append(states, state)
	}

	return states, nil
}

// GetIncidentState - получение состояния инцидента по ID
func (r *PostgresRepo) GetIncidentStateByName(ctx context.Context, stateName string) (entity.IncidentState, error) {
	sql, args, err := r.db.Builder.
		Select("id, name").
		From("incident_states").
		Where(squirrel.Eq{"name": stateName}).
		ToSql()
	if err != nil {
		return entity.IncidentState{}, fmt.Errorf("IncidentStateRepo - GetIncidentState - r.Builder: %w", err)
	}

	row := r.db.Pool.QueryRow(ctx, sql, args...)

	var state entity.IncidentState
	err = row.Scan(
		&state.ID,
		&state.Name,
	)
	if err != nil {
		return entity.IncidentState{}, fmt.Errorf("IncidentStateRepo - GetIncidentState - row.Scan: %w", err)
	}

	return state, nil
}
