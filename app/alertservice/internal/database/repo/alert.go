package repo

import (
	"context"
	"fmt"

	"alertservice/internal/entity"

	"github.com/Masterminds/squirrel"
)

// GetAlerts - получение всех алертов
func (r *PostgresRepo) GetAlerts(ctx context.Context) ([]entity.Alert, error) {
	sql, args, err := r.db.Builder.
		Select("alert_id, alert_name, severity, description, create_at, generator_url, status, job, service, instance, starts_at, ends_at").
		From("alerts").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("AlertRepo - GetAlerts - r.Builder: %w", err)
	}

	rows, err := r.db.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("AlertRepo - GetAlerts - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	alerts := make([]entity.Alert, 0)

	for rows.Next() {
		var alert entity.Alert
		err = rows.Scan(
			&alert.AlertID,
			&alert.AlertName,
			&alert.Severity,
			&alert.Description,
			&alert.CreateAt,
			&alert.GeneratorURL,
			&alert.Status,
			&alert.Job,
			&alert.Service,
			&alert.Instance,
			&alert.StartsAt,
			&alert.EndsAt,
		)
		if err != nil {
			return nil, fmt.Errorf("AlertRepo - GetAlerts - rows.Scan: %w", err)
		}
		alerts = append(alerts, alert)
	}

	return alerts, nil
}

// GetAlerts - получение алертов начиная с индекса begin, count штук
func (r *PostgresRepo) GetListAlerts(ctx context.Context, begin, count int) ([]entity.Alert, error) {
	/*
		SELECT alert_id, alert_name, severity, description, create_at, generator_url, status, starts_at, ends_at
		FROM alerts
		ORDER BY create_at DESC
		OFFSET $1
		LIMIT $2;
	*/
	sql, args, err := r.db.Builder.
		Select("alert_id, alert_name, severity, description, create_at, generator_url, status, job, service, instance, starts_at, ends_at").
		From("alerts").
		OrderBy("create_at DESC").
		Offset(uint64(begin)).
		Limit(uint64(count)).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("AlertRepo - GetListAlerts - r.Builder: %w", err)
	}

	rows, err := r.db.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("AlertRepo - GetListAlerts - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	alerts := make([]entity.Alert, 0)

	for rows.Next() {
		var alert entity.Alert
		err = rows.Scan(
			&alert.AlertID,
			&alert.AlertName,
			&alert.Severity,
			&alert.Description,
			&alert.CreateAt,
			&alert.GeneratorURL,
			&alert.Status,
			&alert.Job,
			&alert.Service,
			&alert.Instance,
			&alert.StartsAt,
			&alert.EndsAt,
		)
		if err != nil {
			return nil, fmt.Errorf("AlertRepo - GetListAlerts - rows.Scan: %w", err)
		}
		alerts = append(alerts, alert)
	}

	return alerts, nil
}

// GetAlert - получение алерта по alertID
func (r *PostgresRepo) GetAlert(ctx context.Context, alertID int64) (entity.Alert, error) {
	sql, args, err := r.db.Builder.
		Select("alert_id, alert_name, severity, description, create_at, generator_url, status, job, service, instance, starts_at, ends_at").
		From("alerts").
		Where(squirrel.Eq{"alert_id": alertID}).
		ToSql()
	if err != nil {
		return entity.Alert{}, fmt.Errorf("AlertRepo - GetAlert - r.Builder: %w", err)
	}

	row := r.db.Pool.QueryRow(ctx, sql, args...)

	var alert entity.Alert
	err = row.Scan(
		&alert.AlertID,
		&alert.AlertName,
		&alert.Severity,
		&alert.Description,
		&alert.CreateAt,
		&alert.GeneratorURL,
		&alert.Status,
		&alert.Job,
		&alert.Service,
		&alert.Instance,
		&alert.StartsAt,
		&alert.EndsAt,
	)
	if err != nil {
		return entity.Alert{}, fmt.Errorf("AlertRepo - GetAlert - row.Scan: %w", err)
	}

	return alert, nil
}

// StoreAlert - запись алерта в таблицу alerts
func (r *PostgresRepo) StoreAlert(ctx context.Context, alert entity.Alert) (int64, error) {
	sql, args, err := r.db.Builder.
		Insert("alerts").
		Columns("alert_name, severity, description, create_at, generator_url, status, job, service, instance, starts_at, ends_at").
		Values(
			alert.AlertName,
			alert.Severity,
			alert.Description,
			alert.CreateAt,
			alert.GeneratorURL,
			alert.Status,
			alert.Job,
			alert.Service,
			alert.Instance,
			alert.StartsAt,
			alert.EndsAt,
		).
		Suffix("RETURNING alert_id").
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("AlertRepo - StoreAlert - r.Builder: %w", err)
	}

	var alertID int64
	err = r.db.Pool.QueryRow(ctx, sql, args...).Scan(&alertID)
	if err != nil {
		return 0, fmt.Errorf("AlertRepo - StoreAlert - r.Pool.QueryRow: %w", err)
	}

	return alertID, nil
}

// DeleteAlert - удаление алерта по alertID
func (r *PostgresRepo) DeleteAlert(ctx context.Context, alertID int64) error {
	sql, args, err := r.db.Builder.
		Delete("alerts").
		Where(squirrel.Eq{"alert_id": alertID}).
		ToSql()
	if err != nil {
		return fmt.Errorf("AlertRepo - DeleteAlert - r.Builder: %w", err)
	}

	_, err = r.db.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("AlertRepo - DeleteAlert - r.Pool.Exec: %w", err)
	}

	return nil
}
