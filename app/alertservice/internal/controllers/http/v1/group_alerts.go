package v1

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// GetGroupedIncidents хендлер для получения сгруппированных инцидентов.
// @Summary Получить сгруппированные инциденты
// @Description Возвращает список инцидентов с детализированной информацией, отсортированных по времени начала
// @Tags groups
// @Accept json
// @Produce json
// @Success 200 {array} entity.GroupIncident
// @Router /api/v1/alerts/groups [get]
func (s *Server) GetGroupedIncidents(w http.ResponseWriter, r *http.Request) {
	s.logger.Debug("GetGroupedIncidents")

	incidents, err := s.u.GetGroupIncidents(r.Context())
	if err != nil {
		http.Error(w, "Failed to retrieve incidents", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(incidents); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// GetListGroupedIncidents хендлер для получения сгруппированных инцидентов от begin в размере count.
// @Summary Получить сгруппированные инциденты
// @Description Возвращает список инцидентов с детализированной информацией, отсортированных по времени начала. Начиная с begin в колличестве count.
// @Tags groups
// @Accept json
// @Produce json
// @Param begin query int true "Starting alert index"
// @Param count query int true "Number of alerts"
// @Success 200 {array} entity.GroupIncident
// @Router /api/v1/alerts/groups [get]
func (s *Server) GetListGroupedIncidents(w http.ResponseWriter, r *http.Request) {
	begin, err := strconv.Atoi(r.URL.Query().Get("begin"))
	if err != nil || begin < 0 {
		http.Error(w, "Invalid parameter begin", http.StatusBadRequest)
		return
	}

	count, err := strconv.Atoi(r.URL.Query().Get("count"))
	if err != nil || count <= 0 {
		http.Error(w, "Invalid parameter count", http.StatusBadRequest)
		return
	}

	incidents, err := s.u.GetGroupIncidents(r.Context())
	if err != nil {
		http.Error(w, "Failed to retrieve incidents", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(incidents); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
