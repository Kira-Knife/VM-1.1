package repo

import (
	"alertservice/internal/entity"
	"context"
	"fmt"
)

// GetIncidentsWithDetails возвращает инциденты с детализированной информацией.
func (r *PostgresRepo) GetIncidentsWithDetails(ctx context.Context) ([]entity.GroupIncident, error) {
	sql, args, err := r.db.Builder.
		Select(`
      i.incident_id,
      s.name AS severity,
      i.description,
      st.name AS status,
      i.create_at,
      i.generator_url,
      a.alert_name,
      COUNT(ia.alert_id) AS alert_count,
      MIN(a.starts_at) AS first_start_at,
      MAX(a.starts_at) AS last_start_at
    `).
		From("incidents i").
		Join("severities s ON i.severity_id = s.id").
		Join("incident_states st ON i.status_id = st.id").
		Join("incident_alerts ia ON i.incident_id = ia.incident_id").
		Join("alerts a ON ia.alert_id = a.alert_id").
		GroupBy("i.incident_id, s.name, st.name, a.alert_name").
		OrderBy("MAX(a.starts_at) DESC").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("GetIncidentsWithDetails - r.Builder: %w", err)
	}

	rows, err := r.db.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("GetIncidentsWithDetails - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	incidents := make([]entity.GroupIncident, 0)

	for rows.Next() {
		var incident entity.GroupIncident
		err = rows.Scan(
			&incident.IncidentID,
			&incident.Severity,
			&incident.Description,
			&incident.Status,
			&incident.CreateAt,
			&incident.GeneratorURL,
			&incident.AlertName,
			&incident.AlertCount,
			&incident.FirstStartAt,
			&incident.LastStartAt,
		)
		if err != nil {
			return nil, fmt.Errorf("GetIncidentsWithDetails - rows.Scan: %w", err)
		}
		incidents = append(incidents, incident)
	}

	return incidents, nil
}
