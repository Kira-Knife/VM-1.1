package config

import (
	"log"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	YDB struct {
		ConnectionString string `yaml:"connection_string"`
		PoolSize int `yaml:"pool_size"`
	} `yaml:"ydb"`

/*Структура AlertStatuses содержит Statuses - Список статусов алертов.
Значения ("Создано", "В работе" и т.д.) загружаются из YAML-файла*/
	AlertStatuses struct {
		Statuses []string `yaml:"statuses"`
	} `yaml:"alert_statuses"`
}

/*Функция LoadConfig:
- читает файл конфигурации с указанным путём;
- парсит содержимое файла и преобразует его в структуру Config.*/
func LoadConfig(path string) *Config {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Открытие файла неуспешно: %v", err)
	}
	defer file.Close()

// дописать декодер
// дописать логгирование ошибок при загрузке

	return cfg
}