/* Схема таблицы для хранения статусов событий алертов*/

CREATE TABLE alert_statuses (
    alert_id Utf8,
    status Utf8,
    updated_at Datetime,
    PRIMARY KEY (alert_id)
);
