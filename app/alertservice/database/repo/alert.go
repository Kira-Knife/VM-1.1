package repo

import (
	"context"
	"fmt"

	"alertservice/entity"

	"github.com/Masterminds/squirrel"
)

func (r *TranslationRepo) GetAlerts(ctx context.Context) ([]entity.Alert, error) {
	sql, args, err := r.Builder.
		Select("alert_id, alert_name, severity, description, timestamp, generator_url, status").
		From("alerts").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("AlertRepo - GetAlerts - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("AlertRepo - GetAlerts - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	alerts := make([]entity.Alert, 0)

	for rows.Next() {
		var alert entity.Alert
		err = rows.Scan(&alert.AlertID, &alert.AlertName, &alert.Severity, &alert.Description, &alert.Timestamp, &alert.GeneratorURL, &alert.Status)
		if err != nil {
			return nil, fmt.Errorf("AlertRepo - GetAlerts - rows.Scan: %w", err)
		}
		alerts = append(alerts, alert)
	}

	return alerts, nil
}

func (r *TranslationRepo) GetAlert(ctx context.Context, alertID int) (entity.Alert, error) {
	sql, args, err := r.Builder.
		Select("alert_id, alert_name, severity, description, timestamp, generator_url, status").
		From("alerts").
		Where(squirrel.Eq{"alert_id": alertID}).
		ToSql()
	if err != nil {
		return entity.Alert{}, fmt.Errorf("AlertRepo - GetAlert - r.Builder: %w", err)
	}

	row := r.Pool.QueryRow(ctx, sql, args...)

	var alert entity.Alert
	err = row.Scan(&alert.AlertID, &alert.AlertName, &alert.Severity, &alert.Description, &alert.Timestamp, &alert.GeneratorURL, &alert.Status)
	if err != nil {
		return entity.Alert{}, fmt.Errorf("AlertRepo - GetAlert - row.Scan: %w", err)
	}

	return alert, nil
}

func (r *TranslationRepo) StoreAlert(ctx context.Context, alert entity.Alert) (int, error) {
	var alertID int

	sql, args, err := r.Builder.
		Insert("alerts").
		Columns("alert_name, severity, description, timestamp, generator_url, status").
		Values(alert.AlertName, alert.Severity, alert.Description, alert.Timestamp, alert.GeneratorURL, alert.Status).
		Suffix("RETURNING alert_id"). // Добавляем RETURNING для получения alert_id
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("AlertRepo - StoreAlert - r.Builder: %w", err)
	}

	// Выполняем запрос и получаем alert_id
	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&alertID)
	if err != nil {
		return 0, fmt.Errorf("AlertRepo - StoreAlert - r.Pool.QueryRow: %w", err)
	}

	return alertID, nil
}
