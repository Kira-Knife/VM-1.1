package repo

import (
	"alertservice/internal/entity"
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
)

// InsertVMAlertJson добавляет новую запись в таблицу vm_alert.
func (r *PostgresRepo) InsertVMAlertJson(ctx context.Context, alert entity.VMAlertJson) (int64, error) {
	sql, args, err := r.db.Builder.
		Insert("vm_alert").
		Columns("alert").
		Values(alert.Alert).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("VMAlertJsonRepo - InsertVMAlertJson - r.Builder: %w", err)
	}

	var id int64
	err = r.db.Pool.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("VMAlertJsonRepo - InsertVMAlertJson - r.Pool.QueryRow: %w", err)
	}

	return id, nil
}

// GetAllVMAlertJsons возвращает все записи из таблицы vm_alert.
func (r *PostgresRepo) GetAllVMAlertJsons(ctx context.Context) ([]entity.VMAlertJson, error) {
	sql, args, err := r.db.Builder.
		Select("id, alert").
		From("vm_alert").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("VMAlertJsonRepo - GetAllVMAlertJsons - r.Builder: %w", err)
	}

	rows, err := r.db.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("VMAlertJsonRepo - GetAllVMAlertJsons - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	alerts := make([]entity.VMAlertJson, 0)

	for rows.Next() {
		var alert entity.VMAlertJson
		err = rows.Scan(&alert.ID, &alert.Alert)
		if err != nil {
			return nil, fmt.Errorf("VMAlertJsonRepo - GetAllVMAlertJsons - rows.Scan: %w", err)
		}
		alerts = append(alerts, alert)
	}

	return alerts, nil
}

// GetVMAlertJsonByID возвращает запись из таблицы vm_alert по ID.
func (r *PostgresRepo) GetVMAlertJsonByID(ctx context.Context, id int64) (entity.VMAlertJson, error) {
	sql, args, err := r.db.Builder.
		Select("id, alert").
		From("vm_alert").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return entity.VMAlertJson{}, fmt.Errorf("VMAlertJsonRepo - GetVMAlertJsonByID - r.Builder: %w", err)
	}

	var alert entity.VMAlertJson
	err = r.db.Pool.QueryRow(ctx, sql, args...).Scan(&alert.ID, &alert.Alert)
	if err != nil {
		return entity.VMAlertJson{}, fmt.Errorf("VMAlertJsonRepo - GetVMAlertJsonByID - r.Pool.QueryRow: %w", err)
	}

	return alert, nil
}

// DeleteVMAlertJsonByID удаляет запись из таблицы vm_alert по ID.
func (r *PostgresRepo) DeleteVMAlertJsonByID(ctx context.Context, id int64) error {
	sql, args, err := r.db.Builder.
		Delete("vm_alert").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return fmt.Errorf("VMAlertJsonRepo - DeleteVMAlertJsonByID - r.Builder: %w", err)
	}

	_, err = r.db.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("VMAlertJsonRepo - DeleteVMAlertJsonByID - r.Pool.Exec: %w", err)
	}

	return nil
}
