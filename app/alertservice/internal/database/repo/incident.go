package repo

import (
	"alertservice/internal/entity"
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
)

// GetIncidents retrieves all incidents from the database.
func (r *PostgresRepo) GetIncidents(ctx context.Context) ([]entity.Incident, error) {
	sql, args, err := r.db.Builder.
		Select("incident_id, severity_id, description, assigned, status_id, create_at, update_at, generator_url").
		From("incidents").
		OrderBy("create_at DESC").
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
		err = rows.Scan(
			&incident.IncidentID,
			&incident.SeverityID,
			&incident.Description,
			&incident.Assigned,
			&incident.StatusID,
			&incident.CreateAt,
			&incident.UpdateAt,
			&incident.GeneratorURL,
		)
		if err != nil {
			return nil, fmt.Errorf("IncidentRepo - GetIncidents - rows.Scan: %w", err)
		}
		incidents = append(incidents, incident)
	}

	return incidents, nil
}

func (r *PostgresRepo) GetIncidentsResponse(ctx context.Context) ([]entity.IncidentResponse, error) {
	sql, args, err := r.db.Builder.
		Select("i.incident_id, i.assigned, s.name AS severity, s.priority AS priority, i.description, is2.name AS status, i.create_at, i.update_at, i.generator_url").
		From("incidents i").
		Join("severities s ON i.severity_id = s.id").
		Join("incident_states is2 ON i.status_id = is2.id").
		OrderBy("i.create_at DESC").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("IncidentRepo - GetIncidents - r.Builder: %w", err)
	}

	rows, err := r.db.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("IncidentRepo - GetIncidents - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	incidents := make([]entity.IncidentResponse, 0)

	for rows.Next() {
		var incident entity.IncidentResponse
		err = rows.Scan(
			&incident.IncidentID,
			&incident.Assigned,
			&incident.Severity,
			&incident.Priority,
			&incident.Description,
			&incident.Status,
			&incident.CreateAt,
			&incident.UpdateAt,
			&incident.GeneratorURL,
		)
		if err != nil {
			return nil, fmt.Errorf("IncidentRepo - GetIncidents - rows.Scan: %w", err)
		}
		incidents = append(incidents, incident)
	}

	return incidents, nil
}

// StoreIncident inserts a new incident into the database and returns its ID.
func (r *PostgresRepo) StoreIncident(ctx context.Context, incident entity.Incident) (int64, error) {
	sql, args, err := r.db.Builder.
		Insert("incidents").
		Columns("severity_id, description, status_id, assigned, generator_url").
		Values(
			incident.SeverityID,
			incident.Description,
			incident.StatusID,
			incident.Assigned,
			incident.GeneratorURL,
		).
		Suffix("RETURNING incident_id"). // Добавляем RETURNING для получения incident_id
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("IncidentRepo - StoreIncident - r.Builder: %w", err)
	}

	var incidentID int64
	err = r.db.Pool.QueryRow(ctx, sql, args...).Scan(&incidentID)
	if err != nil {
		return 0, fmt.Errorf("IncidentRepo - StoreIncident - r.Pool.QueryRow: %w", err)
	}

	return incidentID, nil
}

// GetIncident retrieves a single incident by its ID.
func (r *PostgresRepo) GetIncident(ctx context.Context, incidentID int64) (entity.Incident, error) {
	sql, args, err := r.db.Builder.
		Select("incident_id, severity_id, description, assigned, status_id, create_at, update_at, generator_url").
		From("incidents").
		Where(squirrel.Eq{"incident_id": incidentID}).
		ToSql()
	if err != nil {
		return entity.Incident{}, fmt.Errorf("GetIncident - r.Builder: %w", err)
	}

	var incident entity.Incident
	err = r.db.Pool.QueryRow(ctx, sql, args...).Scan(
		&incident.IncidentID,
		&incident.SeverityID,
		&incident.Description,
		&incident.Assigned,
		&incident.StatusID,
		&incident.CreateAt,
		&incident.UpdateAt,
		&incident.GeneratorURL,
	)
	if err != nil {
		return entity.Incident{}, fmt.Errorf("GetIncident - r.Pool.QueryRow: %w", err)
	}

	return incident, nil
}

func (r *PostgresRepo) GetIncidentResponse(ctx context.Context, incidentID int64) (entity.IncidentResponse, error) {
	sql, args, err := r.db.Builder.
		Select("i.incident_id, s.name AS severity, i.description, i.assigned, is2.name AS status, s.priority AS priority, i.create_at, i.update_at, i.generator_url").
		From("incidents i").
		Join("severities s ON i.severity_id = s.id").
		Join("incident_states is2 ON i.status_id = is2.id").
		Where(squirrel.Eq{"i.incident_id": incidentID}).
		ToSql()
	if err != nil {
		return entity.IncidentResponse{}, fmt.Errorf("GetIncident - r.Builder: %w", err)
	}

	var incident entity.IncidentResponse
	err = r.db.Pool.QueryRow(ctx, sql, args...).Scan(
		&incident.IncidentID,
		&incident.Severity,
		&incident.Description,
		&incident.Assigned,
		&incident.Status,
		&incident.Priority,
		&incident.CreateAt,
		&incident.UpdateAt,
		&incident.GeneratorURL,
	)
	if err != nil {
		return entity.IncidentResponse{}, fmt.Errorf("GetIncident - r.Pool.QueryRow: %w", err)
	}

	return incident, nil
}

// UpdateIncidentStatus updates the status of an incident.
func (r *PostgresRepo) UpdateIncidentStatus(ctx context.Context, incidentID int64, newStatusID int) error {
	sql, args, err := r.db.Builder.
		Update("incidents").
		Set("status_id", newStatusID).
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

// UpdateIncidentStatus updates the status of an incident.
func (r *PostgresRepo) UpdateIncidentAssigned(ctx context.Context, incidentID int64, newAssigned string) error {
	sql, args, err := r.db.Builder.
		Update("incidents").
		Set("assigned", newAssigned).
		Where(squirrel.Eq{"incident_id": incidentID}).
		ToSql()
	if err != nil {
		return fmt.Errorf("IncidentRepo - UpdateIncidentAssigned - r.Builder: %w", err)
	}

	_, err = r.db.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("IncidentRepo - UpdateIncidentAssigned - r.Pool.Exec: %w", err)
	}

	return nil
}

// GetListIncidents retrieves a list of incidents starting from a specific index, returning a specific count of incidents.
func (r *PostgresRepo) GetListIncidents(ctx context.Context, begin, count int) ([]entity.Incident, error) {
	sql, args, err := r.db.Builder.
		Select("incident_id, severity_id, description, assigned, status_id, create_at, update_at, generator_url").
		From("incidents").
		OrderBy("create_at DESC").
		Offset(uint64(begin)).
		Limit(uint64(count)).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("IncidentRepo - GetListIncidents - r.Builder: %w", err)
	}

	rows, err := r.db.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("IncidentRepo - GetListIncidents - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	incidents := make([]entity.Incident, 0)

	for rows.Next() {
		var incident entity.Incident
		err = rows.Scan(
			&incident.IncidentID,
			&incident.SeverityID,
			&incident.Description,
			&incident.Assigned,
			&incident.StatusID,
			&incident.CreateAt,
			&incident.UpdateAt,
			&incident.GeneratorURL,
		)
		if err != nil {
			return nil, fmt.Errorf("IncidentRepo - GetListIncidents - rows.Scan: %w", err)
		}
		incidents = append(incidents, incident)
	}

	return incidents, nil
}

func (r *PostgresRepo) GetListIncidentsResponse(ctx context.Context, begin, count int) ([]entity.IncidentResponse, error) {
	sql, args, err := r.db.Builder.
		Select("i.incident_id, s.name AS severity, i.description, i.assigned, is2.name AS status, s.priority AS priority, i.create_at, i.update_at, i.generator_url").
		From("incidents i").
		Join("severities s ON i.severity_id = s.id").
		Join("incident_states is2 ON i.status_id = is2.id").
		OrderBy("i.create_at DESC").
		Offset(uint64(begin)).
		Limit(uint64(count)).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("IncidentRepo - GetListIncidents - r.Builder: %w", err)
	}

	rows, err := r.db.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("IncidentRepo - GetListIncidents - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	incidents := make([]entity.IncidentResponse, 0)

	for rows.Next() {
		var incident entity.IncidentResponse
		err = rows.Scan(
			&incident.IncidentID,
			&incident.Severity,
			&incident.Description,
			&incident.Assigned,
			&incident.Status,
			&incident.Priority,
			&incident.CreateAt,
			&incident.UpdateAt,
			&incident.GeneratorURL,
		)
		if err != nil {
			return nil, fmt.Errorf("IncidentRepo - GetListIncidents - rows.Scan: %w", err)
		}
		incidents = append(incidents, incident)
	}

	return incidents, nil
}

func (r *PostgresRepo) GetOpenIncidentIDsByAlertName(ctx context.Context, alertName string) ([]int64, error) {
	sql, args, err := r.db.Builder.
		Select("i.incident_id").
		From("incidents i").
		Where("i.incident_id IN ("+
			"SELECT ia.incident_id FROM incident_alerts ia WHERE ia.alert_id IN ("+
			"SELECT a.alert_id FROM alerts a WHERE a.alert_name = ?))", alertName).
		Where(squirrel.Or{
			squirrel.Eq{"i.status_id": 1}, // Open
			squirrel.Eq{"i.status_id": 2}, // In Progress
		}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("GetIncidentIDsByAlertName - r.Builder: %w", err)
	}

	rows, err := r.db.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("GetIncidentIDsByAlertName - r.db.QueryContext: %w", err)
	}
	defer rows.Close()

	var incidentIDs []int64
	for rows.Next() {
		var incidentID int64
		if err := rows.Scan(&incidentID); err != nil {
			return nil, fmt.Errorf("GetIncidentIDsByAlertName - rows.Scan: %w", err)
		}
		incidentIDs = append(incidentIDs, incidentID)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetIncidentIDsByAlertName - rows.Err: %w", err)
	}

	return incidentIDs, nil
}

func (r *PostgresRepo) GetCountOfAlertsForIncident(ctx context.Context, incidentID int64) (int64, error) {
	sql, args, err := r.db.Builder.
		Select("COUNT(*) AS alert_count").
		From("incident_alerts ia").
		Where("ia.incident_id = ?", incidentID).
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("GetCountOfAlertsForIncident - r.Builder: %w", err)
	}

	var alertCount int64
	err = r.db.Pool.QueryRow(ctx, sql, args...).Scan(&alertCount)
	if err != nil {
		return 0, fmt.Errorf("GetCountOfAlertsForIncident - r.db.QueryRow: %w", err)
	}

	return alertCount, nil
}

func (r *PostgresRepo) GetIncidentResponseByAssigned(ctx context.Context, assigned string) ([]entity.IncidentResponse, error) {
	sql, args, err := r.db.Builder.
		Select("i.incident_id, s.name AS severity, i.description, i.assigned, is2.name AS status, s.priority AS priority, i.create_at, i.update_at, i.generator_url").
		From("incidents i").
		Join("severities s ON i.severity_id = s.id").
		Join("incident_states is2 ON i.status_id = is2.id").
		Where(squirrel.Eq{"i.assigned": assigned}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("GetIncident - r.Builder: %w", err)
	}

	rows, err := r.db.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("IncidentRepo - GetListIncidents - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	incidents := make([]entity.IncidentResponse, 0)

	for rows.Next() {
		var incident entity.IncidentResponse
		err = rows.Scan(
			&incident.IncidentID,
			&incident.Severity,
			&incident.Description,
			&incident.Assigned,
			&incident.Status,
			&incident.Priority,
			&incident.CreateAt,
			&incident.UpdateAt,
			&incident.GeneratorURL,
		)
		if err != nil {
			return nil, fmt.Errorf("IncidentRepo - GetListIncidents - rows.Scan: %w", err)
		}
		incidents = append(incidents, incident)
	}

	return incidents, nil
}
