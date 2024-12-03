package entity

import "time"

// Alert represents an individual alert within the JSON payload.
type VMAlert struct {
	Status       string            `json:"status"`
	Labels       map[string]string `json:"labels"`
	Annotations  map[string]string `json:"annotations"`
	StartsAt     time.Time         `json:"startsAt"`
	EndsAt       time.Time         `json:"endsAt"`
	GeneratorURL string            `json:"generatorURL"`
	Fingerprint  string            `json:"fingerprint"`
	SilenceURL   string            `json:"silenceURL"`
	DashboardURL string            `json:"dashboardURL,omitempty"`
	PanelURL     string            `json:"panelURL,omitempty"`
	Values       interface{}       `json:"values"`
	ValueString  string            `json:"valueString"`
}

// AlertNotification represents the full JSON payload.
type VMAlertNotification struct {
	Receiver          string            `json:"receiver"`
	Status            string            `json:"status"`
	VMAlerts          []VMAlert         `json:"alerts"`
	GroupLabels       map[string]string `json:"groupLabels"`
	CommonLabels      map[string]string `json:"commonLabels"`
	CommonAnnotations map[string]string `json:"commonAnnotations"`
	ExternalURL       string            `json:"externalURL"`
	Version           string            `json:"version"`
	GroupKey          string            `json:"groupKey"`
	TruncatedAlerts   int               `json:"truncatedAlerts"`
	OrgID             int               `json:"orgId"`
	Title             string            `json:"title"`
	State             string            `json:"state"`
	Message           string            `json:"message"`
}

type Alert struct {
	AlertID      int64     `json:"alert_id"`
	AlertName    string    `json:"alert_name"`
	Severity     string    `json:"severity"`
	Description  string    `json:"description,omitempty"`
	CreateAt     time.Time `json:"timestamp"`
	GeneratorURL string    `json:"generator_url,omitempty"`
	Status       string    `json:"status,omitempty"`
	StartsAt     time.Time `json:"startsAt"`
	EndsAt       time.Time `json:"endsAt"`
}

type Incident struct {
	IncidentID   int       `json:"incident_id"`
	AlertID      int64     `json:"alert_id,omitempty"`
	Severity     string    `json:"severity"`
	Description  string    `json:"description,omitempty"`
	Status       string    `json:"status,omitempty"`
	CreateAt     time.Time `json:"timestamp"`
	GeneratorURL string    `json:"generator_url,omitempty"`
}

type Notification struct {
	ID        int    `json:"id"`
	Recipient string `json:"recipient"`
	Message   string `json:"message"`
}

type IntervalSetting struct {
}
