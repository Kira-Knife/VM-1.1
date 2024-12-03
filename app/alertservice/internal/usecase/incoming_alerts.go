package usecase

import (
	"alertservice/internal/entity"
	"fmt"
	"time"
)

func ConvertNotificationToEntities(notification entity.VMAlertNotification) ([]entity.Alert, []entity.Incident) {
	var alerts []entity.Alert
	var incidents []entity.Incident

	for _, alertData := range notification.VMAlerts {
		alert := entity.Alert{
			AlertID:      0, // Присвойте или сгенерируйте уникальный идентификатор
			AlertName:    alertData.Labels["alertname"],
			Severity:     alertData.Labels["severity"],
			Description:  alertData.Annotations["summary"],
			Timestamp:    time.Now(), // Или другой подходящий временной штамп
			GeneratorURL: alertData.GeneratorURL,
			Status:       alertData.Status,
			StartsAt:     alertData.StartsAt,
			EndsAt:       alertData.EndsAt,
		}
		alerts = append(alerts, alert)

		incident := entity.Incident{
			IncidentID:   0, // Присвойте или сгенерируйте уникальный идентификатор
			AlertID:      alert.AlertID,
			Severity:     alert.Severity,
			Description:  alert.Description,
			Status:       alert.Status,
			Timestamp:    alert.Timestamp,
			GeneratorURL: alert.GeneratorURL,
		}
		incidents = append(incidents, incident)
	}

	return alerts, incidents
}

// получеие входящих алертов
func (u *UseCase) IncomingAlerts(vm_alert_notif entity.VMAlertNotification) error {
	// получение
	u.logger.Debug("в usecase пришел алерт")
	// Преобразование в Alert
	alert, incidents := ConvertNotificationToEntities(vm_alert_notif)
	fmt.Printf("Alert: %+v\nIncidents: %+v\n", alert, incidents)
	// создание инцидента + alert в базе

	return nil
}
