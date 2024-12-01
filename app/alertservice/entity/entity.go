package entity

import "time"

type VMAlert struct {
	Labels       map[string]string `json:"labels"`
	Annotations  map[string]string `json:"annotations"`
	StartsAt     time.Time         `json:"startsAt"`
	EndsAt       time.Time         `json:"endsAt"`
	GeneratorURL string            `json:"generatorURL"`
	Fingerprint  string            `json:"fingerprint"`
}

type Alert struct {
	AlertID      int       `json:"alert_id"`
	AlertName    string    `json:"alert_name"`
	Severity     string    `json:"severity"`
	Description  string    `json:"description,omitempty"`
	Timestamp    time.Time `json:"timestamp"`
	GeneratorURL string    `json:"generator_url,omitempty"`
	Status       string    `json:"status,omitempty"`
}

type Incident struct {
	IncidentID   int       `json:"incident_id"`
	AlertID      int       `json:"alert_id,omitempty"`
	Severity     string    `json:"severity"`
	Description  string    `json:"description,omitempty"`
	Status       string    `json:"status,omitempty"`
	Timestamp    time.Time `json:"timestamp"`
	GeneratorURL string    `json:"generator_url,omitempty"`
}

type Notification struct {
	ID        int    `json:"id"`
	Recipient string `json:"recipient"`
	Message   string `json:"message"`
}
