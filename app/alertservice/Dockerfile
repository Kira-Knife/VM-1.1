# Этап сборки
FROM golang:1.23 AS builder

# Установка рабочей директории
WORKDIR /app

# Копируем go.mod и go.sum для кэширования зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем все файлы приложения
COPY . .

# Компилируем приложение
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o alertservice main.go

# Этап выполнения
FROM alpine:latest

# Установка необходимых библиотек
RUN apk --no-cache add ca-certificates

# Копируем скомпилированное приложение из этапа сборaки
WORKDIR /app
COPY --from=builder /app/config ./config
COPY --from=builder /app/docs ./docs
COPY --from=builder /app/alertservice alertservice

# Указываем команду для запуска приложения
CMD ["./alertservice"]

# Указываем порт, который будет использовать приложение
EXPOSE 8787
