package v1

import (
	"encoding/json"
	"net/http"
)

// @Summary Эндпоинт для данных за текущий день
// @Description При данном запросе AlertService возвращает динамически обновляемую статистику по количеству алертов за текущий день - с 00:00:01 до 23:59:59.
// @Tags heatmap
// @Produce json
// @Success 200 {object} entity.HeatmapToday
// @Failure 500 {string} string "Ошибка обработки алерта"
// @Router /api/v1/heatmap/today [get]
func (s *Server) getHeatmapToday(w http.ResponseWriter, r *http.Request) {
	s.logger.Debug("getHeatmapToday")
	data, err := s.u.GetHeatmapToday(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	}
}

// @Summary Получение общих данных за все прошедшие дни
// @Description При данном запросе AlertService возвращает полученные динамически агрегированные данные для отображения Heatmap за прошедшие дни с первой даты, сохраненной в БД, до текущей даты.
// @Tags heatmap
// @Produce json
// @Success 200 {object} entity.HeatmapHistory
// @Failure 500 {string} string "Ошибка обработки алерта"
// @Router /api/v1/heatmap/history [get]
func (s *Server) getHeatmapHistory(w http.ResponseWriter, r *http.Request) {
	// Handler logic
	data, err := s.u.GetHeatmapHistory(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	}
}

// @Summary Получение данных по созданным алертами и инцидентам за все прошедшие дни в виде списка по дням
// @Description При данном запросе AlertService возвращает список HeatmapToday (с числом открытых инцидентов и алертов) по каждому дню за прошедшие дни с первой даты, сохраненной в БД, до текущей даты.
// @Tags heatmap
// @Produce json
// @Success 200 {array} entity.HeatmapToday
// @Failure 500 {string} string "Ошибка обработки алерта"
// @Router /api/v1/heatmap/all [get]
func (s *Server) getHeatmapAllDays(w http.ResponseWriter, r *http.Request) {
	// Handler logic
	data, err := s.u.GetHeatmapAllDays(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	}
}

// @Summary Получение данных по закрытым инцидентам за все прошедшие дни в виде списка по дням
// @Description При данном запросе AlertService возвращает список HeatmapToday по каждому дню  за прошедшие дни с первой даты, сохраненной в БД, до текущей даты.
// @Tags heatmap
// @Produce json
// @Success 200 {array} entity.HeatmapToday
// @Failure 500 {string} string "Ошибка обработки алерта"
// @Router /api/v1/heatmap/all/close [get]
func (s *Server) GetHeatmapAllDaysWithStatusClose(w http.ResponseWriter, r *http.Request) {
	data, err := s.u.GetHeatmapAllDaysWithStatusClose(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	}
}
