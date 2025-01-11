package v1

import (
	_ "alertservice/internal/entity"
	"encoding/json"
	"net/http"
)

// @Summary Get all alerts
// @Description Retrieve a list of alerts states
// @Tags alerts
// @Produce json
// @Success 200 {array} entity.AlertState
// @Router /api/v1/alert/states [get]
func (s *Server) getAlertStates(w http.ResponseWriter, r *http.Request) {
	// Handler logic
	alerts, err := s.u.GetIncidentStates()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	} else {
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(alerts)
	}
}
