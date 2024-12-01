package v1

import (
	"alertservice/entity"
	"encoding/json"
	"net/http"
)

func (s *Server) incomingAlerts(w http.ResponseWriter, r *http.Request) {
	var alerts []entity.VMAlert

	// Декодируем тело запроса в срез VMAlert
	if err := json.NewDecoder(r.Body).Decode(&alerts); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// Обрабатываем полученные алерты
	if err := s.u.IncomingAlerts(alerts); err != nil {
		http.Error(w, "Alert processing error", http.StatusInternalServerError)
		return
	}

	// Возвращаем успешный ответ
	w.WriteHeader(http.StatusAccepted)
	_, _ = w.Write([]byte("Alerts received and processed successfully."))
}
