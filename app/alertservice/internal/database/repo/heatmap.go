package repo

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

// GetCountAlertsForPeriod - получение количества созданных алертов за указанный период
func (r *PostgresRepo) GetCountAlertsForPeriod(ctx context.Context, startOfDay time.Time, endOfDay time.Time) (int64, error) {
	r.logger.Debug("repo - GetCountAlertsForPeriod")
	sql, args, err := r.db.Builder.
		Select("COUNT(*)").
		From("alerts").
		Where("starts_at >= ? AND starts_at <= ?", startOfDay, endOfDay).
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("PostgresRepo - GetCountAlertsForPeriod - r.Builder: %w", err)
	}

	var countAlerts int64
	err = r.db.Pool.QueryRow(ctx, sql, args...).Scan(&countAlerts)
	if err != nil {
		return 0, fmt.Errorf("PostgresRepo - GetCountAlertsForPeriod - r.Pool.QueryRow: %w", err)
	}

	return countAlerts, nil
}

// GetCountIncidentsForPeriod - получение количества созданных инцидентов за указанный период
func (r *PostgresRepo) GetCountIncidentsForPeriod(ctx context.Context, startOfDay time.Time, endOfDay time.Time) (int64, error) {
	r.logger.Debug("repo - GetCountIncidentsForPeriod")
	sql, args, err := r.db.Builder.
		Select("COUNT(*)").
		From("incidents").
		Where("create_at >= ? AND create_at <= ?", startOfDay, endOfDay).
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("PostgresRepo - GetCountIncidentsForPeriod - r.Builder: %w", err)
	}

	var countIncidents int64
	err = r.db.Pool.QueryRow(ctx, sql, args...).Scan(&countIncidents)
	if err != nil {
		return 0, fmt.Errorf("PostgresRepo - GetCountIncidentsForPeriod - r.Pool.QueryRow: %w", err)
	}

	return countIncidents, nil
}

// GetCountAlertsForPeriod - получение количества алертов за всё время
func (r *PostgresRepo) GetCountAlertsForAllTime(ctx context.Context) (int64, error) {
	r.logger.Debug("repo - GetCountAlertsForPeriod")
	sql, args, err := r.db.Builder.
		Select("COUNT(*)").
		From("alerts").
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("PostgresRepo - GetCountAlertsForAllTime - r.Builder: %w", err)
	}

	var countAlerts int64
	err = r.db.Pool.QueryRow(ctx, sql, args...).Scan(&countAlerts)
	if err != nil {
		return 0, fmt.Errorf("PostgresRepo - GetCountAlertsForAllTime - r.Pool.QueryRow: %w", err)
	}

	return countAlerts, nil
}

// GetCountIncidentsForPeriod - получение количества инцидентов за всё время
func (r *PostgresRepo) GetCountIncidentsForAllTime(ctx context.Context) (int64, error) {
	r.logger.Debug("repo - GetCountIncidentsForAllTime")
	sql, args, err := r.db.Builder.
		Select("COUNT(*)").
		From("incidents").
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("PostgresRepo - GetCountIncidentsForAllTime - r.Builder: %w", err)
	}

	var countIncidents int64
	err = r.db.Pool.QueryRow(ctx, sql, args...).Scan(&countIncidents)
	if err != nil {
		return 0, fmt.Errorf("PostgresRepo - GetCountIncidentsForAllTime - r.Pool.QueryRow: %w", err)
	}

	return countIncidents, nil
}

// GetMinMaxAlertDates - получение минимальной и максимальной даты создания алертов
func (r *PostgresRepo) GetMinMaxAlertDates(ctx context.Context) (time.Time, time.Time, error) {
	r.logger.Debug("repo - GetMinMaxAlertDates")

	var minDate, maxDate time.Time

	// starts_at - время создания алерта в VM
	// create_at - время создания алерта в системе AlertService
	// так получаем максимальный диапазон по времени
	sqlq := `
        SELECT MIN(starts_at), MAX(create_at)
        FROM alerts;
    `

	err := r.db.Pool.QueryRow(ctx, sqlq).Scan(&minDate, &maxDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return time.Time{}, time.Time{}, fmt.Errorf("no alerts found")
		}
		return time.Time{}, time.Time{}, fmt.Errorf("PostgresRepo - GetMinMaxAlertDates - QueryRow: %w", err)
	}

	return minDate, maxDate, nil
}

// GetCountIncidentsForPeriod - получение количества закрытых инцидентов за указанный период
// Находим инциденты за указанный период по update_at со статусом "Закрыт"
func (r *PostgresRepo) GetCountIncidentsForPeriodWithState(ctx context.Context, startOfDay time.Time, endOfDay time.Time, stateName string) (int64, error) {
	r.logger.Debug("repo - GetCountIncidentsForPeriod")

	state, err := r.GetIncidentStateByName(ctx, stateName)
	if err != nil {
		return 0, err
	}

	sql, args, err := r.db.Builder.
		Select("COUNT(*)").
		From("incidents").
		Where("create_at >= ? AND create_at <= ? AND status_id = ?", startOfDay, endOfDay, state.Name).
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("PostgresRepo - GetCountIncidentsForPeriod - r.Builder: %w", err)
	}

	var countIncidents int64
	err = r.db.Pool.QueryRow(ctx, sql, args...).Scan(&countIncidents)
	if err != nil {
		return 0, fmt.Errorf("PostgresRepo - GetCountIncidentsForPeriod - r.Pool.QueryRow: %w", err)
	}

	return countIncidents, nil
}
