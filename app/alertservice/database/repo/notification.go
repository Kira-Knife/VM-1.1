package repo

import (
	"context"
	"fmt"

	"alertservice/entity"
)

func (r *TranslationRepo) GetNotifications(ctx context.Context) ([]entity.Notification, error) {
	sql, _, err := r.Builder.
		Select("recipient, message").
		From("notifications").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("NotificationRepo - GetNotifications - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("NotificationRepo - GetNotifications - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	notifications := make([]entity.Notification, 0, _defaultEntityCap)

	for rows.Next() {
		notification := entity.Notification{}
		err = rows.Scan(&notification.Recipient, &notification.Message)
		if err != nil {
			return nil, fmt.Errorf("NotificationRepo - GetNotifications - rows.Scan: %w", err)
		}
		notifications = append(notifications, notification)
	}

	return notifications, nil
}

func (r *TranslationRepo) StoreNotification(ctx context.Context, notification entity.Notification) error {
	sql, args, err := r.Builder.
		Insert("notifications").
		Columns("recipient, message").
		Values(notification.Recipient, notification.Message).
		ToSql()
	if err != nil {
		return fmt.Errorf("NotificationRepo - StoreNotification - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("NotificationRepo - StoreNotification - r.Pool.Exec: %w", err)
	}

	return nil
}
