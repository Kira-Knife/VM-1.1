package v1

import (
	"encoding/json"
	"net/http"
)

// GetGroupedIncidents хендлер для получения сгруппированных инцидентов.
// @Summary Получить сгруппированные инциденты
// @Description Возвращает список инцидентов с детализированной информацией, отсортированных по времени начала
// @Tags groups
// @Accept json
// @Produce json
// @Success 200 {array} entity.Incident
// @Router /api/v1/alerts/groups [get]
func (s *Server) GetGroupedIncidents(w http.ResponseWriter, r *http.Request) {
	s.logger.Debug("GetGroupedIncidents")

	incidents, err := s.u.GetSortedIncidents(r.Context())
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
