package config

import (
	"log"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	YDB struct {
		ConnectionString string `yaml:"connection_string"`
		PoolSize         int    `yaml:"pool_size"`
	} `yaml:"ydb"`

	AlertStatuses struct {
		Statuses []string `yaml:"statuses"`
	} `yaml:"alert_statuses"`
}

func LoadConfig(path string) *Config {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}
	defer file.Close()

// дописать декодер

	return cfg
}