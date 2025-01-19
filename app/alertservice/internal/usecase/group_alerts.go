package usecase

import (
	"alertservice/internal/entity"
	"context"
)

// GetSortedIncidents возвращает отсортированный список инцидентов с детализированной информацией.
func (u *UseCase) GetGroupIncidents(ctx context.Context) ([]entity.GroupIncident, error) {
	groupIncidents, err := u.db.GetIncidentsWithDetails(ctx)
	if err != nil {
		return nil, err
	}
	// Инциденты уже отсортированы по starts_at в SQL-запросе.
	return groupIncidents, nil
}
