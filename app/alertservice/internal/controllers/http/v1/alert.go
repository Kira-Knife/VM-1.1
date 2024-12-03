package v1

import (
	"alertservice/internal/entity"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// incomingAlerts хендлер получения входящих от VM алертов
// @Summary принимает алерты от VM
// @Description ппринимает алерты от VM и сохраняет их в базу с дальнейшем уведомлением
// @Tags alerts
// @Accept json
// @Produce json
// @Param notification body entity.VMAlertNotification true "уведомление об алертах"
// @Success 200 {int} http.StatusOK
// @Router /api/v1/alerts [post]
func (s *Server) incomingAlerts(w http.ResponseWriter, r *http.Request) {
	var notification entity.VMAlertNotification
	if err := json.NewDecoder(r.Body).Decode(&notification); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Обрабатываем полученные алерты
	if err := s.u.IncomingAlerts(notification); err != nil {
		http.Error(w, "Alert processing error", http.StatusInternalServerError)
		return
	}

	// Возвращаем успешный ответ
	w.WriteHeader(http.StatusOK)
}

// @Summary Get all alerts
// @Description Retrieve a list of alerts
// @Tags alerts
// @Produce json
// @Success 200 {array} entity.Alert
// @Router /api/v1/alerts [get]
func (s *Server) getAlerts(w http.ResponseWriter, r *http.Request) {
	// Handler logic
	alerts, err := s.u.GetAlerts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	} else {
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(alerts)
	}

}

// @Summary Get alert by ID
// @Description Retrieve an alert by its ID
// @Tags alerts
// @Produce json
// @Param alert_id path int true "Alert ID"
// @Success 200 {object} entity.Alert
// @Router /api/v1/alerts/{alert_id:int64} [get]
func (s *Server) getAlertByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	alertIDStr := vars["alert_id"]                       // Получаем alert_id из параметров
	alertID, err := strconv.ParseInt(alertIDStr, 10, 64) // Преобразуем в int64
	if err != nil {
		http.Error(w, "Invalid alert ID", http.StatusBadRequest)
		return
	}

	// Вызов метода GetAlert
	alert, err := s.u.GetAlert(alertID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Установка заголовка Content-Type
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Возврат JSON
	if err := json.NewEncoder(w).Encode(alert); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
