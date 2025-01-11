-- Alerts table
CREATE TABLE alerts (
  alert_id BIGSERIAL PRIMARY KEY,  -- Unique identifier for each alert
  alert_name VARCHAR(255) NOT NULL,  -- Name of the alert
  severity VARCHAR(50) NOT NULL,  -- Severity level of the alert
  description TEXT,  -- Detailed description of the alert
  create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Timestamp when the alert was created
  generator_url TEXT,  -- URL of the alert generator
  status VARCHAR(50),  -- Current status of the alert
  job VARCHAR(255),  -- Job name, e.g., node_exporter, kafka_exporter
  service VARCHAR(255),  -- Arbitrary label for service, can be null
  instance VARCHAR(255),  -- IP and port of the source, e.g., 84.252.132.204:9100
  starts_at TIMESTAMP,  -- Start time of the alert
  ends_at TIMESTAMP  -- End time of the alert
);

-- Incidents table
CREATE TABLE incidents (
  incident_id BIGSERIAL PRIMARY KEY,  -- Unique identifier for each incident
  severity_id INT REFERENCES severities(id),  -- Level of criticality
  description TEXT,  -- Detailed description of the incident
  status_id INT NOT NULL REFERENCES incident_states(id),  -- Current status of the incident
  create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Timestamp when the incident was created
  generator_url TEXT  -- URL of the incident generator
);

-- Alert states table
CREATE TABLE incident_states (
  id SERIAL PRIMARY KEY,  -- Unique identifier for each state
  name VARCHAR(50) NOT NULL,  -- Name of the state
  description TEXT  -- Detailed description of the state
);

-- Severities table
CREATE TABLE severities (
  id BIGSERIAL PRIMARY KEY,  -- Unique identifier for each severity level
  name VARCHAR(50) NOT NULL UNIQUE,  -- Name of the severity level
  description TEXT  -- Detailed description of the severity level
);

-- Table to track alerts per incident
CREATE TABLE incident_alerts (
  id BIGSERIAL PRIMARY KEY,  -- Unique identifier for each record
  incident_id BIGINT NOT NULL REFERENCES incidents(incident_id),  -- Reference to the incident
  alert_id BIGINT NOT NULL REFERENCES alerts(alert_id)  -- Reference to the alert
);

-- Notifications table
CREATE TABLE notifications (
  id BIGSERIAL PRIMARY KEY,  -- Unique identifier for each notification
  recipient VARCHAR(255) NOT NULL,  -- Recipient of the notification
  message TEXT NOT NULL  -- Message content of the notification
);





INSERT INTO incident_states (name, description) VALUES
('Created', ''),
('In progress', ''),
('Resolved', ''),
('Rejected', '');







