package v1

import (
	_ "alertservice/internal/entity"
	"encoding/json"
	"net/http"
)

// @Summary Get all alerts
// @Description Retrieve a list of alerts states
// @Tags state
// @Produce json
// @Success 200 {array} entity.IncidentState
// @Router /api/v1/incident_states [get]
func (s *Server) getIncidentStates(w http.ResponseWriter, r *http.Request) {
	// Handler logic
	in_states, err := s.u.GetIncidentStates()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	} else {
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(in_states)
	}
}
