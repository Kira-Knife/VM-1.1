# Название приложения
APP_NAME := RequestTimeStats

# Определение операционной системы
OS := $(shell uname -s)

# Флаги компиляции
GO_FLAGS := -ldflags="-s -w"

# Установка исполняемого файла в зависимости от операционной системы
ifeq ($(OS), Linux)
	BINARY := $(APP_NAME)
else
	BINARY := $(APP_NAME).exe
endif

# Цель по умолчанию
all: build

# Сборка приложения
build:
	@echo Building $(APP_NAME) for $(OS)...
	go build $(GO_FLAGS) -o $(BINARY) main.go


# Запуск приложения
run: build
	@echo Running $(APP_NAME)...
	./$(BINARY)

# Очистка файлов сборки
clean:
	@echo Cleaning up...
	rm -f $(BINARY)

# Помощь
help:
	@echo "Makefile for building and launching the application $(APP_NAME)"
	@echo "Available commands:"
	@echo "	make build - Build the application"
	@echo "	make run - Build and launch the application"
	@echo "	make clean - Delete compiled files"
	@echo "	make help - Show this message"

.PHONY: all build run clean help
