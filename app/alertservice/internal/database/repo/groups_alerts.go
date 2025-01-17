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
		incidents = append(incidents, incident)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetIncidentsWithDetails - rows.Err: %w", err)
	}

	return incidents, nil
}
