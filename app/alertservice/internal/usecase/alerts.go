package usecase

import (
	"alertservice/internal/entity"
	appjira "alertservice/internal/usecase/app_jira"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

const (
	IncidentStateOpen       = 1
	IncidentStateInProgress = 2
	IncidentStateResolved   = 3
	IncidentStateRejected   = 4
)

// ConvertNotificationToEntities преобразует VMAlertNotification в массивы Alert и Incident.
func ConvertNotificationToEntities(notification entity.VMAlertNotification) ([]entity.Alert, []entity.Incident) {
	var alerts []entity.Alert
	var incidents []entity.Incident

	for _, alertData := range notification.VMAlerts {
		alert := entity.Alert{
			AlertName:    alertData.Labels["alertname"],
			Severity:     alertData.Labels["severity"],
			Description:  alertData.Annotations["summary"],
			CreateAt:     time.Now(),
			GeneratorURL: alertData.GeneratorURL,
			Status:       alertData.Status,
			Job:          alertData.Labels["job"],
			Service:      alertData.Labels["service"],
			Instance:     alertData.Labels["instance"],
			StartsAt:     alertData.StartsAt,
			EndsAt:       alertData.EndsAt,
		}
		alerts = append(alerts, alert)

		incident := entity.Incident{
			SeverityID:   parseSeverity(alert.Severity),
			Description:  alert.Description,
			StatusID:     parseStatus(alert.Status),
			CreateAt:     alert.CreateAt,
			GeneratorURL: alert.GeneratorURL,
		}
		incidents = append(incidents, incident)
	}

	return alerts, incidents
}

// parseSeverity преобразует строковое значение серьезности в ID.
func parseSeverity(severity string) int {
	// Пример преобразования, замените на реальную логику
	switch severity {
	case "critical":
		return 3
	case "warning":
		return 2
	default:
		return 1
	}
}

// parseStatus преобразует строковое значение статуса в ID.
func parseStatus(status string) int {
	// Пример преобразования, замените на реальную логику
	switch status {
	case "firing":
		return 3
	case "resolved":
		return 2
	default:
		return 1
	}
}

// IncomingAlerts - обработка входящих алертов
func (u *UseCase) IncomingAlerts(vm_alert_notif entity.VMAlertNotification) error {
	u.logger.Debug("IncomingAlerts: %+v", vm_alert_notif)
	// Преобразование в Alert
	alerts, incidents := ConvertNotificationToEntities(vm_alert_notif)
	u.logger.Debug("Успешная конвертация: %+v,%+v.", alerts, incidents)

	ctx, _ := context.WithCancel(context.Background())
	for i, alert := range alerts {
		openedInсidentsId, err := u.db.GetOpenIncidentIDsByAlertName(ctx, alert.AlertName) // openedInсidentsId - либо 0, либо 1
		var incidentID int64
		if len(openedInсidentsId) == 0 || errors.Is(err, sql.ErrNoRows) { // если нет открытого инцидента по данному алерту, то создается инцидент
			incidents[i].StatusID = IncidentStateOpen               // статус по умолчанию
			incidentID, err = u.db.StoreIncident(ctx, incidents[i]) // Сохранение инцидента в базу
			if err != nil {
				u.logger.Error("u.db.StoreIncident(ctx, incidents[i]) - Incident sescription: %s - err: %v", incidents[i].Description, err)
				u.db.DeleteAlert(ctx, incidentID)
				continue
			}

			// При создании нового инцидента создаём задачу в Jira
			taskReq := appjira.TaskJiraRequest{
				IncidentId: incidentID,
				Status:     "Открыт",
				TaskTitle:  incidents[i].Description,
				Assigned:   "user1",
				Owner:      "user2",
			}
			uuId, err := u.appJiraApi.CreateTaskJira(ctx, taskReq)
			if err != nil {
				u.logger.Warn("Не удалось создать задачу Jira: %v", err)
			} else {
				u.logger.Info("Создана задача Jira с uuid: %v", uuId)
			}

		} else {
			incidentID = openedInсidentsId[0] //
		}
		alertId, err := u.db.StoreAlert(ctx, alert) // сохранение алерта в базу
		if err != nil {
			u.logger.Error("u.db.StoreAlert(ctx, alert) - alert name: %s - err: %v", alert.AlertName, err)
			continue
		}

		err = u.db.StoreIncidentAlert(ctx,
			entity.IncidentAlert{IncidentID: incidentID, AlertID: alertId},
		)
		if err != nil { // подправить обработку ошибки
			u.logger.Error("u.db.StoreIncidentAlert: %w", err)
			u.db.DeleteAlert(ctx, alertId)
			continue
		}

		jsonAlert, err := json.Marshal(vm_alert_notif.VMAlerts[i])
		if err != nil {
			u.logger.Error("json.Marshal(vm_alert_notif.VMAlerts[i]): %w", err)
			continue
		}
		vmAlertJsonID, err := u.db.InsertVMAlertJson(ctx, entity.VMAlertJson{
			Alert: json.RawMessage(jsonAlert),
		})

		err = u.db.InsertVMAlertServAlert(ctx, entity.VMAlertServAlert{
			AlertID:       alertId,
			VMAlertJsonID: vmAlertJsonID,
		})
		if err != nil {
			u.logger.Error("u.db.InsertVMAlertServAlert: %w", err)
			continue
		}
	}

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

// GetAlerts - Вернуть алерты начиная с индекса begin, count штук
func (u *UseCase) GetListAlerts(begin, count int) ([]entity.Alert, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	alerts, err := u.db.GetListAlerts(ctx, begin, count)
	if err != nil {
		u.logger.Error("u.db.GetListAlerts(ctx, begin, count); %v", err)
		return nil, fmt.Errorf("u.db.GetListAlerts(ctx, begin, count) - %w", err)
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
