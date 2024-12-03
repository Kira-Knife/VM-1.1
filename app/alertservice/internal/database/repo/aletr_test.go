package repo

import (
	"context"
	"testing"
	"time"

	"alertservice/internal/entity"
	"alertservice/pkg/logger"
	"alertservice/pkg/postgres"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
)

var testPool *pgxpool.Pool
var testRepo *TranslationRepo

func setup() {
	var err error
	testPool, err = pgxpool.Connect(context.Background(), "postgres://admin:admin@localhost:5432/vmalertservice?sslmode=disable")
	if err != nil {
		panic(err)
	}

	// Инициализация репозитория
	testRepo = New(&postgres.Postgres{Pool: testPool}, logger.New("debug"))

	// Автоматическая миграция для создания таблиц
	testRepo.AutoMigrate()
}

func teardown() {
	// Удаление таблиц после тестов
	_, _ = testPool.Exec(context.Background(), "DROP TABLE IF EXISTS notifications;")
	_, _ = testPool.Exec(context.Background(), "DROP TABLE IF EXISTS incidents;")
	_, _ = testPool.Exec(context.Background(), "DROP TABLE IF EXISTS alerts;")
	testPool.Close()
}

func TestStoreAlert(t *testing.T) {
	setup()
	defer teardown()

	alert := entity.Alert{
		AlertName:    "Test Alert",
		Severity:     "High",
		Description:  "This is a test alert",
		Timestamp:    time.Now(),
		GeneratorURL: "http://example.com",
		Status:       "active",
	}

	_, err := testRepo.StoreAlert(context.Background(), alert)
	assert.NoError(t, err)

	alerts, err := testRepo.GetAlerts(context.Background())
	assert.NoError(t, err)
	assert.Len(t, alerts, 1)
	assert.Equal(t, alerts[0].AlertName, alert.AlertName)
}

func TestGetAlert(t *testing.T) {
	setup()
	defer teardown()

	alert := entity.Alert{
		AlertName:    "Test Alert",
		Severity:     "High",
		Description:  "This is a test alert",
		Timestamp:    time.Now(),
		GeneratorURL: "http://example.com",
		Status:       "active",
	}

	_, err := testRepo.StoreAlert(context.Background(), alert)
	assert.NoError(t, err)

	fetchedAlert, err := testRepo.GetAlert(context.Background(), alert.AlertID)
	assert.NoError(t, err)
	assert.Equal(t, fetchedAlert.AlertName, alert.AlertName)
}

func TestGetAlerts(t *testing.T) {
	setup()
	defer teardown()

	alert1 := entity.Alert{
		AlertName:    "Test Alert 1",
		Severity:     "High",
		Description:  "This is a test alert 1",
		Timestamp:    time.Now(),
		GeneratorURL: "http://example.com",
		Status:       "active",
	}

	alert2 := entity.Alert{
		AlertName:    "Test Alert 2",
		Severity:     "Medium",
		Description:  "This is a test alert 2",
		Timestamp:    time.Now(),
		GeneratorURL: "http://example.com",
		Status:       "active",
	}

	_, err := testRepo.StoreAlert(context.Background(), alert1)
	assert.NoError(t, err)

	_, err = testRepo.StoreAlert(context.Background(), alert2)
	assert.NoError(t, err)

	alerts, err := testRepo.GetAlerts(context.Background())
	assert.NoError(t, err)
	assert.Len(t, alerts, 2)
}
