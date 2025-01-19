package usecase

import (
	"alertservice/internal/entity"
	"context"
	"fmt"
)

// GetAlerts - Вернуть все Alert
func (u *UseCase) GetIncidents() ([]entity.IncidentResponse, error) {
	ctx, _ := context.WithCancel(context.Background())
	ints, err := u.db.GetIncidentsResponse(ctx)
	if err != nil {
		u.logger.Error("u.db.GetIncidents(ctx); %v", err)
		return nil, fmt.Errorf("u.db.GetIncidents(ctx) - %w", err)
	}
	return ints, nil
}

// GetAlert - вернуть Alert по alertId
func (u *UseCase) GetIncident(incidentID int64) (entity.IncidentResponse, error) {
	ctx, _ := context.WithCancel(context.Background())
	i, err := u.db.GetIncidentResponse(ctx, incidentID)
	if err != nil {
		u.logger.Error("u.db.GetAlert(ctx, alertID) - incidentID %d - %v", incidentID, err)
		return i, fmt.Errorf("u.db.GetAlert(ctx, alertID) - %w", err)
	}
	return i, nil
}

// UpdateIncidentStatus - обновить статуст инцидента
// обновление ID статуса инцидента
// Проверить что такой статус существует, если нет то вернуть ошибку с возможными статусами
func (u *UseCase) UpdateIncidentStatus(incidentID int64, newStatus string) error {
	ctx, _ := context.WithCancel(context.Background())
	// err := u.db.UpdateIncidentStatus(ctx, incidentID, newStatus)
	st, _ := u.db.GetIncidentStates(ctx)
	newStatusID := st[0].ID
	err := u.db.UpdateIncidentStatus(ctx, incidentID, newStatusID)
	if err != nil {
		u.logger.Error("UpdateIncidentStatus - incidentID %d - %v", incidentID, err)
		return fmt.Errorf("UpdateIncidentStatus - %w", err)
	}
	return err
}

// GetAlerts - Вернуть алерты начиная с индекса begin, count штук
func (u *UseCase) GetListIncidents(begin, count int) ([]entity.IncidentResponse, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	incidents, err := u.db.GetListIncidentsResponse(ctx, begin, count)
	if err != nil {
		u.logger.Error("u.db.GetListAlerts(ctx, begin, count); %v", err)
		return nil, fmt.Errorf("u.db.GetListAlerts(ctx, begin, count) - %w", err)
	}
	return incidents, nil
}
