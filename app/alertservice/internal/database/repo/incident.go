package repo

import (
	"context"
	"fmt"

	"alertservice/internal/entity"

	"github.com/Masterminds/squirrel"
)

// GetIncidents retrieves all incidents from the database.
func (r *PostgresRepo) GetIncidents(ctx context.Context) ([]entity.Incident, error) {
	sql, args, err := r.db.Builder.
		Select("incident_id, alert_id, severity, description, status, create_at, generator_url").
		From("incidents").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("IncidentRepo - GetIncidents - r.Builder: %w", err)
	}

	rows, err := r.db.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("IncidentRepo - GetIncidents - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	incidents := make([]entity.Incident, 0)

	for rows.Next() {
		var incident entity.Incident
		err = rows.Scan(&incident.IncidentID, &incident.AlertID, &incident.Severity, &incident.Description, &incident.Status, &incident.CreateAt, &incident.GeneratorURL)
		if err != nil {
			return nil, fmt.Errorf("IncidentRepo - GetIncidents - rows.Scan: %w", err)
		}
		incidents = append(incidents, incident)
	}

	return incidents, nil
}

// StoreIncident inserts a new incident into the database.
func (r *PostgresRepo) StoreIncident(ctx context.Context, incident entity.Incident) error {
	sql, args, err := r.db.Builder.
		Insert("incidents").
		Columns("alert_id, severity, description, status, create_at, generator_url").
		Values(incident.AlertID, incident.Severity, incident.Description, incident.Status, incident.CreateAt, incident.GeneratorURL).
		ToSql()
	if err != nil {
		return fmt.Errorf("IncidentRepo - StoreIncident - r.Builder: %w", err)
	}

	_, err = r.db.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("IncidentRepo - StoreIncident - r.Pool.Exec: %w", err)
	}

	return nil
}

// GetIncident retrieves a single incident by its ID.
func (r *PostgresRepo) GetIncident(ctx context.Context, incidentID int64) (entity.Incident, error) {
	sql, args, err := r.db.Builder.
		Select("incident_id, alert_id, severity, description, status, create_at, generator_url").
		From("incidents").
		Where(squirrel.Eq{"incident_id": incidentID}).
		ToSql()
	if err != nil {
		return entity.Incident{}, fmt.Errorf("IncidentRepo - GetIncident - r.Builder: %w", err)
	}

	var incident entity.Incident
	err = r.db.Pool.QueryRow(ctx, sql, args...).Scan(
		&incident.IncidentID,
		&incident.AlertID,
		&incident.Severity,
		&incident.Description,
		&incident.Status,
		&incident.CreateAt,
		&incident.GeneratorURL,
	)
	if err != nil {
		return entity.Incident{}, fmt.Errorf("IncidentRepo - GetIncident - r.Pool.QueryRow: %w", err)
	}

	return incident, nil
}

// UpdateIncidentStatus updates the status of an incident.
func (r *PostgresRepo) UpdateIncidentStatus(ctx context.Context, incidentID int64, newStatus string) error {
	sql, args, err := r.db.Builder.
		Update("incidents").
		Set("status", newStatus).
		Where(squirrel.Eq{"incident_id": incidentID}).
		ToSql()
	if err != nil {
		return fmt.Errorf("IncidentRepo - UpdateIncidentStatus - r.Builder: %w", err)
	}

	_, err = r.db.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("IncidentRepo - UpdateIncidentStatus - r.Pool.Exec: %w", err)
	}

	return nil
}
