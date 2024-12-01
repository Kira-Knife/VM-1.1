package entity

import "time"

// Alert represents the alerts table with its fields.
type Alert struct {
	AlertID      int       `json:"alert_id"`
	AlertName    string    `json:"alert_name"`
	Severity     string    `json:"severity"`
	Description  string    `json:"description,omitempty"`
	Timestamp    time.Time `json:"timestamp"`
	GeneratorURL string    `json:"generator_url,omitempty"`
	Status       string    `json:"status,omitempty"`
}

// Incident represents the incidents table with its fields.
type Incident struct {
	IncidentID   int       `json:"incident_id"`
	AlertID      int       `json:"alert_id,omitempty"`
	Severity     string    `json:"severity"`
	Description  string    `json:"description,omitempty"`
	Status       string    `json:"status,omitempty"`
	Timestamp    time.Time `json:"timestamp"`
	GeneratorURL string    `json:"generator_url,omitempty"`
}

// Notification represents the notifications table with its fields.
type Notification struct {
	ID        int    `json:"id"`
	Recipient string `json:"recipient"`
	Message   string `json:"message"`
}
