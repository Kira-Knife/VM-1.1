package v1

import "net/http"

// @Summary Get interval setting
// @Description Retrieve the current interval setting
// @Tags settings
// @Produce json
// @Success 200 {object} entity.IntervalSetting
// @Router /api/v1/settings/interval [get]
func (s *Server) getInterval(w http.ResponseWriter, r *http.Request) {
	// Handler logic
}

// @Summary Set interval setting
// @Description Set a new interval setting
// @Tags settings
// @Accept json
// @Produce json
// @Param interval body entity.IntervalSetting true "Interval Setting"
// @Success 204
// @Router /api/v1/settings/interval [post]
func (s *Server) setInterval(w http.ResponseWriter, r *http.Request) {
	// Handler logic
}
