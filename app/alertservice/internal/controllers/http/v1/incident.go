package v1

import (
	"encoding/json"
	"fmt"
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
// @Success 200
// @Failure 400 {object} map[string]string "Invalid incident ID or request body"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/v1/incidents/{incident_id} [patch]
func (s *Server) updateIncidentStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	s.logger.Debug("Run updateIncidentStatus")
	// Извлечение incident_id из параметров маршрута
	vars := mux.Vars(r)
	incidentIDStr := vars["incident_id"]                       // Получаем incident_id из параметров
	incidentID, err := strconv.ParseInt(incidentIDStr, 10, 64) // Преобразуем в int64
	if err != nil {
		err = fmt.Errorf("Invalid incident ID: %w", err)
		s.logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// Извлечение нового статуса из тела запроса
	var requestBody Status

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		err = fmt.Errorf("Invalid request body: %w", err)
		s.logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	s.logger.Debug("Вызов метода UpdateIncidentStatus")
	// Вызов метода UpdateIncidentStatus
	err = s.u.UpdateIncidentStatus(incidentID, requestBody.NewStatus)
	if err != nil {
		s.logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// Успешное обновление статуса
	w.WriteHeader(http.StatusOK) // 200 No Content
}

// @Summary Получение списка инцидентов
// @Description Вернет список инцидентов отсортированных по времени начиная с begin в размере count
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

// @Summary Получение количества алертов для данного инцидента
// @Description Вернет количество алертов по данногому инциденту (на данный момент у одного инцидента алерты с одним именем)
// @Tags incidents
// @Produce json
// @Param incident_id path int true "Incident ID"
// @Success 200
// @Router /api/v1/incidents/{incident_id}/alerts/count [get]
func (s *Server) getCountOfAlertsForIncident(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	incidentIDStr := vars["incident_id"]                       // Получаем incident_id из параметров
	incidentID, err := strconv.ParseInt(incidentIDStr, 10, 64) // Преобразуем в int64
	if err != nil {
		http.Error(w, "Invalid `incident_id`", http.StatusBadRequest)
		return
	}

	alerts_count, err := s.u.GetCountOfAlertsForIncident(incidentID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(alerts_count)
}
