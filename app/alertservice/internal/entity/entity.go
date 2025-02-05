package entity

import (
	"encoding/json"
	"time"
)

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

// Alert represents an alert entity
type Alert struct {
	AlertID      int64     `db:"alert_id" json:"alert_id"`           // Unique identifier for each alert
	AlertName    string    `db:"alert_name" json:"alert_name"`       // Name of the alert
	Severity     string    `db:"severity" json:"severity"`           // Severity level of the alert
	Description  string    `db:"description" json:"description"`     // Detailed description of the alert
	GeneratorURL string    `db:"generator_url" json:"generator_url"` // URL of the alert generator
	Status       string    `db:"status" json:"status"`               // Current status of the alert
	Job          string    `db:"job" json:"job"`                     // Job name, e.g., node_exporter, kafka_exporter
	Service      string    `db:"service" json:"service"`             // Arbitrary label for service, can be null
	Instance     string    `db:"instance" json:"instance"`           // IP and port of the source
	StartsAt     time.Time `db:"starts_at" json:"starts_at"`         // Start time of the alert
	EndsAt       time.Time `db:"ends_at" json:"ends_at"`             // End time of the alert
	CreateAt     time.Time `db:"create_at" json:"create_at"`         // Timestamp when the alert was created
	UpdateAt     time.Time `db:"update_at" json:"update_at"`
}

// Incident represents an incident entity
type Incident struct {
	IncidentID   int64     `db:"incident_id" json:"incident_id"`     // Unique identifier for each incident
	SeverityID   int       `db:"severity_id" json:"severity_id"`     // Level of criticality
	Description  string    `db:"description" json:"description"`     // Detailed description of the incident
	Assigned     string    `json:"assigned"`                         // Assigned
	StatusID     int       `db:"status_id" json:"status_id"`         // Current status of the incident
	GeneratorURL string    `db:"generator_url" json:"generator_url"` // URL of the incident generator
	CreateAt     time.Time `db:"create_at" json:"create_at"`         // Timestamp when the incident was created
	UpdateAt     time.Time `db:"update_at" json:"update_at"`
}

// В таком виде инцидент получает фронт
type IncidentResponse struct {
	IncidentID   int64     `json:"incident_id"`                      // Unique identifier for each incident
	Severity     string    `json:"severity"`                         // Level of criticality
	Priority     string    `json:"priority"`                         // Level of criticality
	Assigned     string    `json:"assigned"`                         // Assigned
	Description  string    `db:"description" json:"description"`     // Detailed description of the incident
	Status       string    `json:"status"`                           // Current status of the incident
	CreateAt     time.Time `db:"create_at" json:"create_at"`         // Timestamp when the incident was created
	GeneratorURL string    `db:"generator_url" json:"generator_url"` // URL of the incident generator
	UpdateAt     time.Time `db:"update_at" json:"update_at"`
}

// IncidentState represents the state of an incident
type IncidentState struct {
	ID   int    `db:"id" json:"id"` // Unique identifier for each state
	Name string `db:"name" json:"name"`
}

// Severity represents a severity level
type Severity struct {
	ID       int64  `db:"id" json:"id"`             // Unique identifier for each severity level
	Name     string `db:"name" json:"name"`         // Name of the severity level
	Priority string `db:"priority" json:"priority"` // priority
}

// IncidentAlert represents the relationship between incidents and alerts
type IncidentAlert struct {
	ID         int64 `db:"id" json:"id"`                   // Unique identifier for each record
	IncidentID int64 `db:"incident_id" json:"incident_id"` // Reference to the incident
	AlertID    int64 `db:"alert_id" json:"alert_id"`       // Reference to the alert
}

// Notification represents a notification entity
type Notification struct {
	ID        int64  `db:"id" json:"id"`               // Unique identifier for each notification
	Recipient string `db:"recipient" json:"recipient"` // Recipient of the notification
	Message   string `db:"message" json:"message"`     // Message content of the notification
}

type IntervalSetting struct {
}

// VMAlert представляет собой структуру для таблицы vm_alert.
type VMAlertJson struct {
	ID    int64           `json:"id"`    // Уникальный идентификатор
	Alert json.RawMessage `json:"alert"` // Данные в формате JSON
}

// VMAlertServAlert представляет собой структуру для таблицы vm_alert_serv_alert.
type VMAlertServAlert struct {
	AlertID       int64 `json:"alert_id"`    // Идентификатор оповещения
	VMAlertJsonID int64 `json:"vm_alert_id"` // Идентификатор VM оповещения
}

type GroupIncident struct {
	IncidentID   int64     `json:"incident_id"`
	Description  string    `json:"description"` // sql.NullString
	CreateAt     time.Time `json:"create_at"`
	GeneratorURL string    `json:"generator_url"` // sql.NullString
	Severity     string    `json:"severity"`
	Priority     string    `json:"priority"`
	Job          string    `json:"job"`      // Job name, e.g., node_exporter, kafka_exporter (экспортер)
	Service      string    `json:"service"`  // Arbitrary label for service, can be null
	Instance     string    `json:"instance"` // хост (Grafana) нужно ещё посмотреть на это
	Status       string    `json:"status"`
	AlertCount   int       `json:"alert_count"`
	FirstStartAt time.Time `json:"first_start_at"`
	LastStartAt  time.Time `json:"last_start_at"`
}
