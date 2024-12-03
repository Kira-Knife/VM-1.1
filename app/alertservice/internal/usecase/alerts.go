package usecase

import (
	"alertservice/internal/entity"
	"context"
	"fmt"
	"time"
)

func ConvertNotificationToEntities(notification entity.VMAlertNotification) ([]entity.Alert, []entity.Incident) {
	var alerts []entity.Alert
	var incidents []entity.Incident

	for _, alertData := range notification.VMAlerts {
		alert := entity.Alert{
			AlertName:    alertData.Labels["alertname"],
			Severity:     alertData.Labels["severity"],
			Description:  alertData.Annotations["summary"],
			CreateAt:     time.Now(), // Или другой подходящий временной штамп
			GeneratorURL: alertData.GeneratorURL,
			Status:       alertData.Status,
			StartsAt:     alertData.StartsAt,
			EndsAt:       alertData.EndsAt,
		}
		alerts = append(alerts, alert)

		incident := entity.Incident{
			AlertID:      alert.AlertID,
			Severity:     alert.Severity,
			Description:  alert.Description,
			Status:       alert.Status,
			CreateAt:     alert.CreateAt,
			GeneratorURL: alert.GeneratorURL,
		}
		incidents = append(incidents, incident)
	}

	return alerts, incidents
}

// IncomingAlerts - обработка входящих алертов
func (u *UseCase) IncomingAlerts(vm_alert_notif entity.VMAlertNotification) error {
	// Преобразование в Alert
	alerts, incidents := ConvertNotificationToEntities(vm_alert_notif)
	ctx, _ := context.WithCancel(context.Background())
	for i, alert := range alerts {
		alertId, err := u.db.StoreAlert(ctx, alert)
		if err != nil {
			u.logger.Error("u.db.StoreAlert(ctx, alert) - alert name: %s - err: %v", alert.AlertName, err)
			continue
		}
		incidents[i].AlertID = alertId
		incidents[i].Status = "open" // статус по умолчанию
		err = u.db.StoreIncident(ctx, incidents[i])
		if err != nil {
			u.logger.Error("u.db.StoreIncident(ctx, incidents[i]) - Incident sescription: %s - err: %v", incidents[i].Description, err)
			u.db.DeleteAlert(ctx, alertId)
			continue
		}
	}
	// создание инцидента + alert в базе

	return nil
}

// GetAlerts - Вернуть все Alert
func (u *UseCase) GetAlerts() ([]entity.Alert, error) {
	ctx, _ := context.WithCancel(context.Background())
	alerts, err := u.db.GetAlerts(ctx)
	if err != nil {
		u.logger.Error("u.db.GetAlerts(ctx); %v", err)
		return nil, fmt.Errorf("u.db.GetAlerts(ctx) - %w", err)
	}
	return alerts, nil
}

// GetAlert - вернуть Alert по alertId
func (u *UseCase) GetAlert(alertID int64) (entity.Alert, error) {
	ctx, _ := context.WithCancel(context.Background())
	alert, err := u.db.GetAlert(ctx, alertID)
	if err != nil {
		u.logger.Error("u.db.GetAlert(ctx, alertID) - alertId %d - %v", alertID, err)
		return alert, fmt.Errorf("u.db.GetAlert(ctx, alertID) - %w", err)
	}
	return alert, nil
}
