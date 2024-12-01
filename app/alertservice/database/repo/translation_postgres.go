package repo

import (
	"alertservice/pkg/postgres"
	"context"
	"log"
)

const _defaultEntityCap = 64

// TranslationRepo -.
type TranslationRepo struct {
	*postgres.Postgres
}

// New -.
func New(pg *postgres.Postgres) *TranslationRepo {
	return &TranslationRepo{pg}
}

func (r *TranslationRepo) AutoMigrate() {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS alerts (
            alert_name VARCHAR(255) PRIMARY KEY,
            severity VARCHAR(50),
            description TEXT,
            timestamp TIMESTAMP,
            generator_url TEXT,
            status VARCHAR(50)
        )`,

		`CREATE TABLE IF NOT EXISTS incidents (
            incident_id SERIAL PRIMARY KEY,
            alert_name VARCHAR(255),
            severity VARCHAR(50),
            description TEXT,
            status VARCHAR(50),
            timestamp TIMESTAMP,
            generator_url TEXT
        )`,

		`CREATE TABLE IF NOT EXISTS notifications (
            id SERIAL PRIMARY KEY,
            recipient VARCHAR(255),
            message TEXT
        )`,
	}

	for _, query := range queries {
		_, err := r.Pool.Exec(context.Background(), query)
		if err != nil {
			log.Fatalf("TranslationRepo - AutoMigrate - Exec: %v", err)
		}
	}
}
