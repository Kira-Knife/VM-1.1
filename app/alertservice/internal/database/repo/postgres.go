package repo

import (
	"alertservice/pkg/logger"
	"alertservice/pkg/postgres"
	"context"
	"log"
)

const _defaultEntityCap = 64

// TranslationRepo -.
type PostgresRepo struct {
	db     *postgres.Postgres
	logger *logger.Logger
}

// New -.
func New(pg *postgres.Postgres, l *logger.Logger) *PostgresRepo {
	return &PostgresRepo{
		db:     pg,
		logger: l,
	}
}

func (r *PostgresRepo) AutoMigrate() {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS alerts (
			alert_id SERIAL PRIMARY KEY,
			alert_name VARCHAR(255) NOT NULL, 
			severity VARCHAR(50) NOT NULL,
			description TEXT,
			timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			generator_url TEXT,
			status VARCHAR(50)
		);`,

		`CREATE TABLE IF NOT EXISTS incidents (
			incident_id SERIAL PRIMARY KEY,
			alert_id INTEGER,
			severity VARCHAR(50) NOT NULL,
			description TEXT,
			status VARCHAR(50),
			timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
			generator_url TEXT,
			FOREIGN KEY (alert_id) REFERENCES alerts(alert_id) ON DELETE CASCADE
		);`,

		`CREATE TABLE IF NOT EXISTS notifications (
			id SERIAL PRIMARY KEY,
			recipient VARCHAR(255) NOT NULL,  
			message TEXT NOT NULL
		);`,
	}

	for _, query := range queries {
		_, err := r.db.Pool.Exec(context.Background(), query)
		if err != nil {
			log.Fatalf("TranslationRepo - AutoMigrate - Exec: %v", err)
		}
	}
}
