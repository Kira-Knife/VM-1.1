package usecase

import (
	"alertservice/internal/entity"
	"context"
	"fmt"
)

// GetAlerts - Вернуть все Alert
func (u *UseCase) GetIncidents() ([]entity.Incident, error) {
	ctx, _ := context.WithCancel(context.Background())
	ints, err := u.db.GetIncidents(ctx)
	if err != nil {
		u.logger.Error("u.db.GetIncidents(ctx); %v", err)
		return nil, fmt.Errorf("u.db.GetIncidents(ctx) - %w", err)
	}
	return ints, nil
}

// GetAlert - вернуть Alert по alertId
func (u *UseCase) GetIncident(incidentID int64) (entity.Incident, error) {
	ctx, _ := context.WithCancel(context.Background())
	i, err := u.db.GetIncident(ctx, incidentID)
	if err != nil {
		u.logger.Error("u.db.GetAlert(ctx, alertID) - incidentID %d - %v", incidentID, err)
		return i, fmt.Errorf("u.db.GetAlert(ctx, alertID) - %w", err)
	}
	return i, nil
}

func (u *UseCase) UpdateIncidentStatus(incidentID int64, newStatus string) error {
	ctx, _ := context.WithCancel(context.Background())
	err := u.db.UpdateIncidentStatus(ctx, incidentID, newStatus)
	if err != nil {
		u.logger.Error("UpdateIncidentStatus - incidentID %d - %v", incidentID, err)
		return fmt.Errorf("UpdateIncidentStatus - %w", err)
	}
	return err
}
