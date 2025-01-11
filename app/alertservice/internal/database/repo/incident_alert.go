package repo

import (
	"alertservice/internal/entity"
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
)

// GetIncidentAlert retrieves a single incident-alert association by its ID.
func (r *PostgresRepo) GetIncidentAlert(ctx context.Context, id int64) (entity.IncidentAlert, error) {
	sql, args, err := r.db.Builder.
		Select("id, incident_id, alert_id").
		From("incident_alerts").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return entity.IncidentAlert{}, fmt.Errorf("IncidentAlertRepo - GetIncidentAlert - r.Builder: %w", err)
	}

	var incidentAlert entity.IncidentAlert
	err = r.db.Pool.QueryRow(ctx, sql, args...).Scan(
		&incidentAlert.ID,
		&incidentAlert.IncidentID,
		&incidentAlert.AlertID,
	)
	if err != nil {
		return entity.IncidentAlert{}, fmt.Errorf("IncidentAlertRepo - GetIncidentAlert - r.Pool.QueryRow: %w", err)
	}

	return incidentAlert, nil
}

// GetIncidentAlerts retrieves all incident-alert associations.
func (r *PostgresRepo) GetIncidentAlerts(ctx context.Context) ([]entity.IncidentAlert, error) {
	sql, args, err := r.db.Builder.
		Select("id, incident_id, alert_id").
		From("incident_alerts").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("IncidentAlertRepo - GetIncidentAlerts - r.Builder: %w", err)
	}

	rows, err := r.db.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("IncidentAlertRepo - GetIncidentAlerts - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	incidentAlerts := make([]entity.IncidentAlert, 0)

	for rows.Next() {
		var incidentAlert entity.IncidentAlert
		err = rows.Scan(
			&incidentAlert.ID,
			&incidentAlert.IncidentID,
			&incidentAlert.AlertID,
		)
		if err != nil {
			return nil, fmt.Errorf("IncidentAlertRepo - GetIncidentAlerts - rows.Scan: %w", err)
		}
		incidentAlerts = append(incidentAlerts, incidentAlert)
	}

	return incidentAlerts, nil
}

// StoreIncidentAlert inserts a new incident-alert association into the database.
func (r *PostgresRepo) StoreIncidentAlert(ctx context.Context, incidentAlert entity.IncidentAlert) error {
	sql, args, err := r.db.Builder.
		Insert("incident_alerts").
		Columns("incident_id, alert_id").
		Values(incidentAlert.IncidentID, incidentAlert.AlertID).
		ToSql()
	if err != nil {
		return fmt.Errorf("IncidentAlertRepo - StoreIncidentAlert - r.Builder: %w", err)
	}

	_, err = r.db.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("IncidentAlertRepo - StoreIncidentAlert - r.Pool.Exec: %w", err)
	}

	return nil
}

// GetIncidentAlertPairCount возвращает количество пар с заданными incident_id и alert_id.
func (r *PostgresRepo) GetIncidentAlertPairCount(ctx context.Context, incidentID, alertID int64) (int, error) {
	sql, args, err := r.db.Builder.
		Select("COUNT(*)").
		From("incident_alerts").
		Where(squirrel.Eq{"incident_id": incidentID, "alert_id": alertID}).
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("IncidentAlertRepo - GetIncidentAlertPairCount - r.Builder: %w", err)
	}

	var count int
	err = r.db.Pool.QueryRow(ctx, sql, args...).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("IncidentAlertRepo - GetIncidentAlertPairCount - r.Pool.QueryRow: %w", err)
	}

	return count, nil
}
