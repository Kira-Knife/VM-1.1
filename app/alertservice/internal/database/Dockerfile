# Используем официальный образ PostgreSQL
FROM postgres:latest

# Устанавливаем рабочую директорию
WORKDIR /docker-entrypoint-initdb.d

# Копируем скрипт инициализации в контейнер
COPY ./sql/init.sql .

# Устанавливаем переменные окружения для базы данных
# Эти переменные можно переопределить при запуске контейнера
ENV POSTGRES_DB=vmalertservice
ENV POSTGRES_USER=admin
ENV POSTGRES_PASSWORD=admin

# Контейнер будет автоматически запускать PostgreSQL
# и выполнять скрипты из /docker-entrypoint-initdb.d
