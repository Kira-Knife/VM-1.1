package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Status struct {
	NewStatus string `json:"status"`
}

// @Summary Get all incidents
// @Description Retrieve a list of incidents
// @Tags incidents
// @Produce json
// @Success 200 {array} entity.IncidentResponse
// @Router /api/v1/incidents [get]
func (s *Server) getIncidents(w http.ResponseWriter, r *http.Request) {
	// Handler logic
	ins, err := s.u.GetIncidents()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	} else {
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(ins)
	}
}

// @Summary Get incident by ID
// @Description Retrieve an incident by its ID
// @Tags incidents
// @Produce json
// @Param incident_id path int true "Incident ID"
// @Success 200 {object} entity.IncidentResponse
// @Router /api/v1/incidents/{incident_id} [get]
func (s *Server) getIncidentByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["incidents_id"]                        // Получаем alert_id из параметров
	incidents_id, err := strconv.ParseInt(id, 10, 64) // Преобразуем в int64
	if err != nil {
		http.Error(w, "Invalid alert ID", http.StatusBadRequest)
		return
	}

	// Вызов метода GetAlert
	incident, err := s.u.GetIncident(incidents_id)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Установка заголовка Content-Type
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Возврат JSON
	if err := json.NewEncoder(w).Encode(incident); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// @Summary Update an incident
// @Description Update an existing incident
// @Tags incidents
// @Param incident_id path int true "Incident ID"
// @Param incident_status body Status true "New status"
// @Success 204
// @Router /api/v1/incidents/{incident_id} [patch]
func (s *Server) updateIncidentStatus(w http.ResponseWriter, r *http.Request) {
	// Извлечение incident_id из параметров маршрута
	vars := mux.Vars(r)
	incidentIDStr := vars["incidents_id"]                      // Получаем incident_id из параметров
	incidentID, err := strconv.ParseInt(incidentIDStr, 10, 64) // Преобразуем в int64
	if err != nil {
		http.Error(w, "Invalid incident ID", http.StatusBadRequest)
		return
	}

	// Извлечение нового статуса из тела запроса
	var requestBody Status

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Вызов метода UpdateIncidentStatus
	err = s.u.UpdateIncidentStatus(incidentID, requestBody.NewStatus)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Успешное обновление статуса
	w.WriteHeader(http.StatusNoContent) // 204 No Content
}

// @Summary Получение списка инцидентов
// @Description Вернет список инцидентов отсортированных по времени начиная с begin в колличестве count
// @Tags incidents
// @Produce json
// @Param begin query int true "Starting incident index"
// @Param count query int true "Number of incidents"
// @Success 200 {array} entity.IncidentResponse
// @Router /api/v1/incidents/list [get]
func (s *Server) getListIncidents(w http.ResponseWriter, r *http.Request) {
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

	incidents, err := s.u.GetListIncidents(begin, count)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(incidents)
}
