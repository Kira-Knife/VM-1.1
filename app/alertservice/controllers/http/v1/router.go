package v1

import (
	"alertservice/entity"
	"context"
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) alertHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var alerts []entity.Alert
	if err := json.NewDecoder(r.Body).Decode(&alerts); err != nil {
		http.Error(w, "Failed to decode alert", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	for _, alert := range alerts {
		if err := s.u.StoreAlert(ctx, alert); err != nil {
			log.Printf("Failed to store alert: %v", err)
			http.Error(w, "Failed to store alert", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusAccepted)
}
