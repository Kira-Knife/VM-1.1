package entity

import "time"

type HeatmapToday struct {
	Date           time.Time `json:"date"`            // дата, за котороую собирается статистика
	BeginTime      time.Time `json:"begin_time"`      // дата Date, но время 00:00:01
	EndTime        time.Time `json:"end_time"`        // дата Date, но время 23:59:59
	CountAlerts    int64     `json:"count_alerts"`    // число созданных алертов в диапазоне от BeginTime до EndTime
	CountIncidents int64     `json:"count_incidents"` // число созданных инцидентов в диапазоне от BeginTime до EndTime
	StatisticsUrl  string    `json:"statistics_url"`  // пока пустое поле
}

type HeatmapHistory struct {
	StartDate      time.Time `json:"begin_date"`
	EndDate        time.Time `json:"end_date"`
	CountAlerts    int64     `json:"count_alerts"`
	CountIncidents int64     `json:"count_incidents"`
	StatisticsUrl  string    `json:"statistics_url"`
}
