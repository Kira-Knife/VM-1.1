package usecase

import (
	"alertservice/internal/entity"
	"context"
	"fmt"
	"time"
)

func (u *UseCase) GetHeatmapToday(ctx context.Context) (entity.HeatmapToday, error) {
	u.logger.Debug("usecase - GetHeatmapToday")
	currentDate := time.Now()
	// Начало дня: 00:00:01
	startOfDay := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 0, 0, 1, 0, currentDate.Location())
	// Конец дня: 23:59:59
	endOfDay := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 23, 59, 59, 0, currentDate.Location())

	ht := entity.HeatmapToday{
		BeginTime: startOfDay,
		EndTime:   endOfDay,
		Date:      currentDate,
	}

	alertCount, err := u.db.GetCountAlertsForPeriod(ctx, startOfDay, endOfDay)
	if err != nil {
		return ht, err
	}
	incidentCount, err := u.db.GetCountIncidentsForPeriod(ctx, startOfDay, endOfDay)
	if err != nil {
		return ht, err
	}

	ht.CountAlerts = alertCount
	ht.CountIncidents = incidentCount

	return ht, nil
}

func (u *UseCase) GetHeatmapHistory(ctx context.Context) (entity.HeatmapHistory, error) {
	u.logger.Debug("usecase - GetHeatmapHistory")
	hh := entity.HeatmapHistory{}

	startDate, endDate, err := u.db.GetMinMaxAlertDates(ctx)
	if err != nil {
		return hh, err
	}
	hh.StartDate = startDate
	hh.EndDate = endDate

	countAllAlerts, err := u.db.GetCountAlertsForAllTime(ctx)
	if err != nil {
		return hh, err
	}
	countAllIncidents, err := u.db.GetCountIncidentsForAllTime(ctx)
	if err != nil {
		return hh, err
	}
	hh.CountAlerts = countAllAlerts
	hh.CountIncidents = countAllIncidents

	return hh, nil
}

// Получение данных HeatmapToday за все дни
func (u *UseCase) GetHeatmapAllDays(ctx context.Context) ([]entity.HeatmapToday, error) {
	minStartDate, maxCreateDate, err := u.db.GetMinMaxAlertDates(ctx) // получаем дату первого и последнего алерта
	if err != nil {
		return nil, err
	}

	var heatmapData []entity.HeatmapToday
	maxDay := maxCreateDate.AddDate(0, 0, 1)
	for date := minStartDate; !date.After(maxDay); date = date.AddDate(0, 0, 1) { // проходимся по всем дням
		startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 1, 0, date.Location())
		endOfDay := time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 0, date.Location())

		countAlerts, err := u.db.GetCountAlertsForPeriod(ctx, startOfDay, endOfDay)
		if err != nil {
			return nil, fmt.Errorf("PostgresRepo - GenerateHeatmapData - GetCountAlertsForPeriod: %w", err)
		}

		countIncidents, err := u.db.GetCountIncidentsForPeriod(ctx, startOfDay, endOfDay)
		if err != nil {
			return nil, fmt.Errorf("PostgresRepo - GenerateHeatmapData - GetCountIncidentsForPeriod: %w", err)
		}

		heatmapData = append(heatmapData, entity.HeatmapToday{
			Date:           date,
			BeginTime:      startOfDay,
			EndTime:        endOfDay,
			CountAlerts:    countAlerts,
			CountIncidents: countIncidents,
			StatisticsUrl:  "",
		})
	}

	return heatmapData, nil
}

func (u *UseCase) GetHeatmapAllDaysWithStatusClose(ctx context.Context) ([]entity.HeatmapToday, error) {
	return u.getHeatmapAllDaysWithStatus(ctx, "Решено") // Максимально на коленке - исправить !
}

// Получение данных HeatmapToday за все дни
func (u *UseCase) getHeatmapAllDaysWithStatus(ctx context.Context, statusName string) ([]entity.HeatmapToday, error) {
	minStartDate, maxCreateDate, err := u.db.GetMinMaxAlertDates(ctx) // получаем дату первого и последнего алерта
	if err != nil {
		return nil, err
	}

	var heatmapData []entity.HeatmapToday
	maxDay := maxCreateDate.AddDate(0, 0, 1)
	for date := minStartDate; !date.After(maxDay); date = date.AddDate(0, 0, 1) { // проходимся по всем дням
		startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 1, 0, date.Location())
		endOfDay := time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 0, date.Location())

		countIncidents, err := u.db.GetCountIncidentsForPeriodWithState(ctx, startOfDay, endOfDay, statusName)
		if err != nil {
			return nil, fmt.Errorf("PostgresRepo - GenerateHeatmapData - GetCountIncidentsForPeriod: %w", err)
		}

		heatmapData = append(heatmapData, entity.HeatmapToday{
			Date:           date,
			BeginTime:      startOfDay,
			EndTime:        endOfDay,
			CountIncidents: countIncidents,
			StatisticsUrl:  "",
		})
	}

	return heatmapData, nil
}
