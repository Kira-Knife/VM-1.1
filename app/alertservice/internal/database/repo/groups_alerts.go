package repo

import (
	"alertservice/internal/entity"
	"context"
	"fmt"
)

/*
SELECT
  i.incident_id,
  i.description,
  i.create_at,
  i.generator_url,
  s.name AS severity,
  is2.name AS status,
  in12.alert_count,
  in12.latest_start_at,
  in12.earliest_start_at
FROM
  incidents i
JOIN
  severities s ON i.severity_id = s.id
JOIN
  incident_states is2 ON i.status_id = is2.id
JOIN (
  SELECT
    ia.incident_id,
    COUNT(a.alert_name) AS alert_count,
    MAX(a.starts_at) AS latest_start_at,
    MIN(a.starts_at) AS earliest_start_at
  FROM
    incident_alerts ia
  JOIN
    alerts a ON ia.alert_id = a.alert_id
  GROUP BY
    ia.incident_id
) AS in12 ON in12.incident_id = i.incident_id
ORDER BY
  in12.latest_start_at DESC;
*/

// GetIncidentsWithDetails возвращает инциденты с детализированной информацией.
func (r *PostgresRepo) GetIncidentsWithDetails(ctx context.Context) ([]entity.GroupIncident, error) {
	query := `
	 SELECT 
	  i.incident_id, 
	  i.description, 
	  i.create_at, 
	  i.generator_url, 
	  s.name AS severity, 
	  is2.name AS status, 
	  in12.alert_count, 
	  in12.latest_start_at, 
	  in12.earliest_start_at
	 FROM 
	  incidents i
	 JOIN 
	  severities s ON i.severity_id = s.id
	 JOIN 
	  incident_states is2 ON i.status_id = is2.id
	 JOIN (
	  SELECT 
	   ia.incident_id, 
	   COUNT(a.alert_name) AS alert_count, 
	   MAX(a.starts_at) AS latest_start_at, 
	   MIN(a.starts_at) AS earliest_start_at
	  FROM 
	   incident_alerts ia
	  JOIN 
	   alerts a ON ia.alert_id = a.alert_id
	  GROUP BY 
	   ia.incident_id
	 ) AS in12 ON in12.incident_id = i.incident_id
	 ORDER BY 
	  in12.latest_start_at DESC;
	`

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("GetIncidentsWithDetails - r.db.QueryContext: %w", err)
	}
	defer rows.Close()

	var incidents []entity.GroupIncident
	for rows.Next() {
		var incident entity.GroupIncident
		if err := rows.Scan(
			&incident.IncidentID,
			&incident.Description,
			&incident.CreateAt,
			&incident.GeneratorURL,
			&incident.Severity,
			&incident.Status,
			&incident.AlertCount,
			&incident.LastStartAt,
			&incident.FirstStartAt,
		); err != nil {
			return nil, fmt.Errorf("GetIncidentsWithDetails - rows.Scan: %w", err)
		}

		sql, args, err := r.db.Builder.
			Select("a.job, a.service, a.instance").
			From("alerts a").
			Where("a.alert_id IN (SELECT ia.alert_id FROM incident_alerts ia WHERE ia.incident_id = ?)", incident.IncidentID).
			OrderBy("a.starts_at DESC").
			Limit(1).
			ToSql()
		if err != nil {
			return nil, fmt.Errorf("GetAlertDetailsByIncidentID - r.Builder: %w", err)
		}

		err = r.db.Pool.QueryRow(ctx, sql, args...).Scan(
			&incident.Job,
			&incident.Service,
			&incident.Instance,
		)
		if err != nil {
			return nil, fmt.Errorf("GetAlertDetailsByIncidentID - r.Builder: %w", err)
		}

		incidents = append(incidents, incident)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetIncidentsWithDetails - rows.Err: %w", err)
	}

	return incidents, nil
}

func (r *PostgresRepo) GetListIncidentsWithDetails(ctx context.Context, begin, count int) ([]entity.GroupIncident, error) {
	// Создание подзапроса как строка
	subquery, subargs, err := r.db.Builder.
		Select("ia.incident_id, COUNT(a.alert_name) AS alert_count, MAX(a.starts_at) AS latest_start_at, MIN(a.starts_at) AS earliest_start_at").
		From("incident_alerts ia").
		Join("alerts a ON ia.alert_id = a.alert_id").
		GroupBy("ia.incident_id").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("GetListIncidentsWithDetails - subquery Builder: %w", err)
	}

	sql, args, err := r.db.Builder.
		Select("i.incident_id, i.description, i.create_at, i.generator_url, s.name AS severity, is2.name AS status, in12.alert_count, in12.latest_start_at, in12.earliest_start_at").
		From("incidents i").
		Join("severities s ON i.severity_id = s.id").
		Join("incident_states is2 ON i.status_id = is2.id").
		Join(fmt.Sprintf("(%s) AS in12 ON in12.incident_id = i.incident_id", subquery), subargs...).
		OrderBy("in12.latest_start_at DESC").
		Offset(uint64(begin)).
		Limit(uint64(count)).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("GetListIncidentsWithDetails - r.Builder: %w", err)
	}

	rows, err := r.db.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("GetListIncidentsWithDetails - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	var incidents []entity.GroupIncident
	for rows.Next() {
		var incident entity.GroupIncident
		err = rows.Scan(
			&incident.IncidentID,
			&incident.Description,
			&incident.CreateAt,
			&incident.GeneratorURL,
			&incident.Severity,
			&incident.Status,
			&incident.AlertCount,
			&incident.LastStartAt,
			&incident.FirstStartAt,
		)
		if err != nil {
			return nil, fmt.Errorf("GetListIncidentsWithDetails - rows.Scan: %w", err)
		}

		sql, args, err = r.db.Builder.
			Select("a.job, a.service, a.instance").
			From("alerts a").
			Where("a.alert_id IN (SELECT ia.alert_id FROM incident_alerts ia WHERE ia.incident_id = ?)", incident.IncidentID).
			OrderBy("a.starts_at DESC").
			Limit(1).
			ToSql()
		if err != nil {
			return nil, fmt.Errorf("GetAlertDetailsByIncidentID - r.Builder: %w", err)
		}

		err = r.db.Pool.QueryRow(ctx, sql, args...).Scan(
			&incident.Job,
			&incident.Service,
			&incident.Instance,
		)
		if err != nil {
			return nil, fmt.Errorf("GetAlertDetailsByIncidentID - r.Builder: %w", err)
		}

		incidents = append(incidents, incident)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetListIncidentsWithDetails - rows.Err: %w", err)
	}

	if err != nil {
		return nil, fmt.Errorf("GetAlertDetailsByIncidentID - r.Pool.QueryRow: %w", err)
	}

	return incidents, nil
}
