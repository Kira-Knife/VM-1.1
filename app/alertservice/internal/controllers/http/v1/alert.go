package v1

import (
	"alertservice/internal/entity"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// @Summary Обработка входящих от VM алертов
// @Description Принимает уведомления об алертах (алерты) от VM с их дальнейшей обработкой
// @Tags alerts
// @Accept json
// @Produce json
// @Param notification body entity.VMAlertNotification true "Json структура уведомления об алертах системы VM"
// @Success 200
// @Failure 400 {string} string "Ошибка в теле запроса"
// @Failure 500 {string} string "Ошибка обработки алерта"
// @Router /api/v1/alerts [post]
func (s *Server) incomingAlerts(w http.ResponseWriter, r *http.Request) {
	s.logger.Debug("incomingAlerts")

	var notification entity.VMAlertNotification
	if err := json.NewDecoder(r.Body).Decode(&notification); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Ошибка в теле запроса: %v", err)))
		return
	}

	// Обрабатываем полученные алерты
	if err := s.u.IncomingAlerts(notification); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Ошибка обработки алертов: %v", err)))
		return
	}

	// Возвращаем успешный ответ
	w.WriteHeader(http.StatusOK)
}

// @Summary Получение всех алертов
// @Description Вернет массив всех алертов, отсортированных по дате (по невозрастанию)
// @Tags alerts
// @Produce json
// @Success 200 {array} entity.Alert
// @Failure 500 {string} string "Ошибка получения данных"
// @Router /api/v1/alerts [get]
func (s *Server) getAlerts(w http.ResponseWriter, r *http.Request) {
	alerts, err := s.u.GetAlerts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Ошибка получения данных: %v.", err)))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(alerts)
}

// @Summary Получение списка алертов
// @Description Вернет список алертов отсортированных по времени начиная с begin в размере count
// @Tags alerts
// @Produce json
// @Param begin query int true "Начальный индекс"
// @Param count query int true "Колличество получаемых алертов"
// @Success 200 {array} entity.Alert
// @Failure 400 {string} string "Недопустимые параметры"
// @Failure 500 {string} string "Ошибка получения данных"
// @Router /api/v1/alerts/list [get]
func (s *Server) getListAlerts(w http.ResponseWriter, r *http.Request) {
	begin, err := strconv.Atoi(r.URL.Query().Get("begin"))
	if err != nil || begin < 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Недопустимый параметр `begin` (begin >= 0): %v.", err)))
		return
	}

	count, err := strconv.Atoi(r.URL.Query().Get("count"))
	if err != nil || count <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Недопустимый параметр `count` (count > 0): %v.", err)))
		return
	}

	alerts, err := s.u.GetListAlerts(begin, count)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Ошибка получения списка алертов: %v.", err)))
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(alerts)
}

// @Summary Получение алерта по ID
// @Description Вернет алерт по указанному ID
// @Tags alerts
// @Produce json
// @Param alert_id path int true "Alert ID"
// @Success 200 {object} entity.Alert
// @Failure 400 {string} string "Недопустимые параметры"
// @Failure 500 {string} string "Ошибка получения данных"
// @Router /api/v1/alerts/{alert_id} [get]
func (s *Server) getAlertByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	alertIDStr := vars["alert_id"]                       // Получаем alert_id из параметров
	alertID, err := strconv.ParseInt(alertIDStr, 10, 64) // Преобразуем в int64
	if err != nil || alertID < 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Некорректный параметр ID (int64 > 0): %v.", err)))
		return
	}

	// Вызов метода GetAlert
	alert, err := s.u.GetAlert(alertID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Ошибка получения алерта: %v.", err)))
		return
	}

	// Установка заголовка Content-Type
	w.Header().Set("Content-Type", "application/json")

	// Возврат JSON
	if err := json.NewEncoder(w).Encode(alert); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Failed to encode response: %v.", err)))
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
