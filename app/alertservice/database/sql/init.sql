CREATE TABLE IF NOT EXISTS alerts (
    alert_id SERIAL PRIMARY KEY,
    alert_name VARCHAR(255) NOT NULL,  -- Добавлено NOT NULL для обязательных полей
    severity VARCHAR(50) NOT NULL,
    description TEXT,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Установка значения по умолчанию
    generator_url TEXT,
    status VARCHAR(50)
);

CREATE TABLE IF NOT EXISTS incidents (
    incident_id SERIAL PRIMARY KEY,
    alert_id INTEGER,
    severity VARCHAR(50) NOT NULL,
    description TEXT,
    status VARCHAR(50),
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Установка значения по умолчанию
    generator_url TEXT,
    FOREIGN KEY (alert_id) REFERENCES alerts(alert_id) ON DELETE CASCADE  -- Устанавливаем внешний ключ с каскадным удалением
);

CREATE TABLE IF NOT EXISTS notifications (
    id SERIAL PRIMARY KEY,
    recipient VARCHAR(255) NOT NULL,  -- Добавлено NOT NULL для обязательных полей
    message TEXT NOT NULL  -- Добавлено NOT NULL для обязательных полей
);
