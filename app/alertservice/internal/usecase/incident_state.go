package usecase

import (
	"alertservice/internal/entity"
	"context"
	"fmt"
)

// GetAlerts - Вернуть все Alert
func (u *UseCase) GetAllAlertStates() ([]entity.IncidentState, error) {
	ctx, _ := context.WithCancel(context.Background())
	alert_states, err := u.db.GetAllIncidentStates(ctx)
	if err != nil {
		u.logger.Error("u.db.GetAllAlertStates(ctx); %v", err)
		return nil, fmt.Errorf("u.db.GetAllAlertStates(ctx) - %w", err)
	}
	return alert_states, nil
}
