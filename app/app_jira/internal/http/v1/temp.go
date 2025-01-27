package v1

import (
	"net/http"
)

// @Summary Temp
// @Description Temp
// @Tags temp
// @Produce json
// @Success 200
// @Router /api/v1/temp [post]
func (s *Server) tempHandle(w http.ResponseWriter, r *http.Request) {
	// Handler logic
	w.WriteHeader(http.StatusOK)
}
