package repo

import (
	"alertservice/internal/entity"
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
)

// InsertVMAlertServAlert добавляет новую связь между alert_id и vm_alert_id в таблицу vm_alert_serv_alert.
func (r *PostgresRepo) InsertVMAlertServAlert(ctx context.Context, alertServAlert entity.VMAlertServAlert) error {
	sql, args, err := r.db.Builder.
		Insert("vm_alert_serv_alert").
		Columns("alert_id, vm_alert_id").
		Values(alertServAlert.AlertID, alertServAlert.VMAlertJsonID).
		ToSql()
	if err != nil {
		return fmt.Errorf("VMAlertServAlertRepo - InsertVMAlertServAlert - r.Builder: %w", err)
	}

	_, err = r.db.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("VMAlertServAlertRepo - InsertVMAlertServAlert - r.Pool.Exec: %w", err)
	}

	return nil
}

// GetAllVMAlertServAlerts возвращает все записи из таблицы vm_alert_serv_alert.
func (r *PostgresRepo) GetAllVMAlertServAlerts(ctx context.Context) ([]entity.VMAlertServAlert, error) {
	sql, args, err := r.db.Builder.
		Select("alert_id, vm_alert_id").
		From("vm_alert_serv_alert").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("VMAlertServAlertRepo - GetAllVMAlertServAlerts - r.Builder: %w", err)
	}

	rows, err := r.db.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("VMAlertServAlertRepo - GetAllVMAlertServAlerts - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	alertServAlerts := make([]entity.VMAlertServAlert, 0)

	for rows.Next() {
		var alertServAlert entity.VMAlertServAlert
		err = rows.Scan(&alertServAlert.AlertID, &alertServAlert.VMAlertJsonID)
		if err != nil {
			return nil, fmt.Errorf("VMAlertServAlertRepo - GetAllVMAlertServAlerts - rows.Scan: %w", err)
		}
		alertServAlerts = append(alertServAlerts, alertServAlert)
	}

	return alertServAlerts, nil
}

// GetVMAlertServAlertByID возвращает связь из таблицы vm_alert_serv_alert по alert_id.
func (r *PostgresRepo) GetVMAlertServAlertByID(ctx context.Context, alertID int64) ([]entity.VMAlertServAlert, error) {
	sql, args, err := r.db.Builder.
		Select("alert_id, vm_alert_id").
		From("vm_alert_serv_alert").
		Where(squirrel.Eq{"alert_id": alertID}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("VMAlertServAlertRepo - GetVMAlertServAlertByID - r.Builder: %w", err)
	}

	rows, err := r.db.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("VMAlertServAlertRepo - GetVMAlertServAlertByID - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	alertServAlerts := make([]entity.VMAlertServAlert, 0)

	for rows.Next() {
		var alertServAlert entity.VMAlertServAlert
		err = rows.Scan(&alertServAlert.AlertID, &alertServAlert.VMAlertJsonID)
		if err != nil {
			return nil, fmt.Errorf("VMAlertServAlertRepo - GetVMAlertServAlertByID - rows.Scan: %w", err)
		}
		alertServAlerts = append(alertServAlerts, alertServAlert)
	}

	return alertServAlerts, nil
}

// DeleteVMAlertServAlertByID удаляет связь из таблицы vm_alert_serv_alert по alert_id.
func (r *PostgresRepo) DeleteVMAlertServAlertByID(ctx context.Context, alertID int64) error {
	sql, args, err := r.db.Builder.
		Delete("vm_alert_serv_alert").
		Where(squirrel.Eq{"alert_id": alertID}).
		ToSql()
	if err != nil {
		return fmt.Errorf("VMAlertServAlertRepo - DeleteVMAlertServAlertByID - r.Builder: %w", err)
	}

	_, err = r.db.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("VMAlertServAlertRepo - DeleteVMAlertServAlertByID - r.Pool.Exec: %w", err)
	}

	return nil
}
