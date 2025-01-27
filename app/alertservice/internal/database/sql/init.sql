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

-- Alert states table
CREATE TABLE incident_states (
  id SERIAL PRIMARY KEY,  -- Unique identifier for each state
  name VARCHAR(50) NOT NULL,  -- Name of the state
  description TEXT  -- Detailed description of the state
);

INSERT INTO incident_states (id, name, description) VALUES
(1, 'Открыт', ''),
(2, 'В работе', ''),
(3, 'Решенный', ''),
(4, 'Отклоненный', '');

-- Severities table
CREATE TABLE severities (
  id BIGSERIAL PRIMARY KEY,  -- Unique identifier for each severity level
  name VARCHAR(50) NOT NULL UNIQUE,  -- Name of the severity level
  description TEXT  -- Detailed description of the severity level
);


INSERT INTO severities (name, description) VALUES
('critical', ''),
('warning', ''),
('info', '');

-- Incidents table
CREATE TABLE incidents (
  incident_id BIGSERIAL PRIMARY KEY,  -- Unique identifier for each incident
  severity_id INT REFERENCES severities(id),  -- Level of criticality
  description TEXT,  -- Detailed description of the incident
  status_id INT NOT NULL REFERENCES incident_states(id),  -- Current status of the incident
  create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Timestamp when the incident was created
  generator_url TEXT  -- URL of the incident generator
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

CREATE TABLE vm_alert (
  id BIGSERIAL PRIMARY KEY,
  alert JSON NOT NULL
);

CREATE TABLE vm_alert_serv_alert (
  alert_id BIGINT NOT NULL REFERENCES alerts(alert_id),
  vm_alert_id BIGINT NOT NULL REFERENCES vm_alert(id)
);

-- Добавление заглушки для Jira
CREATE TABLE task_status (
  id SERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL UNIQUE
);

INSERT INTO task_status (id, name) VALUES
(1, 'To Do'),
(2, 'In Progress'),
(3, 'Done'),
(4, 'Canceled');

CREATE TABLE incident_states_task_status (
  task_ststus_id INT NOT NULL REFERENCES task_status(id),
  incident_states_id INT NOT NULL REFERENCES incident_states(id),
);

INSERT INTO incident_states_task_status (task_ststus_id, incident_states_id) VALUES
(1, 1),
(2, 2),
(3, 3),
(4, 4);

CREATE TABLE tasks (
    task_id BIGSERIAL PRIMARY KEY, -- Уникальный идентификатор задачи в нашей системе
    task_main_id UUID UNIQUE, -- Уникальный идентификатор задачи в системе источнике (Jira)
    task_title TEXT, -- Сочетание акронима проекта/пространства и номера задачи в системе источнике
    task_status_id INT NOT NULL REFERENCES task_status(id), -- Статус задачи (To Do, In Progress, Done, Canceled)
    assigned VARCHAR(250), -- Назначенный исполнитель (логин)
    owner VARCHAR(250), -- Наблюдатель  (логин)
    incident_id BIGINT NOT NULL REFERENCES incidents(incident_id)-- Идентификатор инцидента
);