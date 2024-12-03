package repo

/*
import (
	"context"
	"testing"
	"time"

	"alertservice/internal/database/repo"
	"alertservice/internal/entity"

	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/assert"
)

func TestGetAlerts(t *testing.T) {
	db, mock, err := pgxmock.New()
	if err != nil {
		t.Fatalf("failed to create mock: %v", err)
	}
	defer db.Close()

	r := &repo.PostgresRepo{Pool: db}

	// Настройка ожиданий
	mock.ExpectQuery("SELECT alert_id, alert_name, severity, description, timestamp, generator_url, status, starts_at, ends_at").
		WillReturnRows(pgxmock.NewRows([]string{"alert_id", "alert_name", "severity", "description", "timestamp", "generator_url", "status", "starts_at", "ends_at"}).
			AddRow(1, "Test Alert", "high", "Test description", time.Now(), "http://example.com", "active", time.Now(), time.Now()))

	alerts, err := r.GetAlerts(context.Background())
	assert.NoError(t, err)
	assert.Len(t, alerts, 1)
	assert.Equal(t, "Test Alert", alerts[0].AlertName)

	// Проверка ожиданий
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetAlert(t *testing.T) {
	db, mock, err := pgxmock.New()
	if err != nil {
		t.Fatalf("failed to create mock: %v", err)
	}
	defer db.Close()

	r := &repo.PostgresRepo{Pool: db}

	alertID := int64(1)
	mock.ExpectQuery("SELECT alert_id, alert_name, severity, description, timestamp, generator_url, status, starts_at, ends_at").
		WithArgs(alertID).
		WillReturnRows(pgxmock.NewRows([]string{"alert_id", "alert_name", "severity", "description", "timestamp", "generator_url", "status", "starts_at", "ends_at"}).
			AddRow(alertID, "Test Alert", "high", "Test description", time.Now(), "http://example.com", "active", time.Now(), time.Now()))

	alert, err := r.GetAlert(context.Background(), alertID)
	assert.NoError(t, err)
	assert.Equal(t, "Test Alert", alert.AlertName)

	// Проверка ожиданий
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestStoreAlert(t *testing.T) {
	db, mock, err := pgxmock.New()
	if err != nil {
		t.Fatalf("failed to create mock: %v", err)
	}
	defer db.Close()

	r := &repo.PostgresRepo{Pool: db}

	alert := entity.Alert{
		AlertName:    "Test Alert",
		Severity:     "high",
		Description:  "Test description",
		CreateAt:     time.Now(),
		GeneratorURL: "http://example.com",
		Status:       "active",
		StartsAt:     time.Now(),
		EndsAt:       time.Now(),
	}

	mock.ExpectQuery("INSERT INTO alerts").
		WithArgs(alert.AlertName, alert.Severity, alert.Description, alert.CreateAt, alert.GeneratorURL, alert.Status, alert.StartsAt, alert.EndsAt).
		WillReturnRows(pgxmock.NewRows([]string{"alert_id"}).AddRow(1))

	alertID, err := r.StoreAlert(context.Background(), alert)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), alertID)

	// Проверка ожиданий
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteAlert(t *testing.T) {
	db, mock, err := pgxmock.New()
	if err != nil {
		t.Fatalf("failed to create mock: %v", err)
	}
	defer db.Close()

	r := &repo.PostgresRepo{Pool: db}

	alertID := int64(1)
	mock.ExpectExec("DELETE FROM alerts").
		WithArgs(alertID).
		WillReturnResult(pgxmock.NewResult(1, 1)) // Успешное удаление одной записи

	err = r.DeleteAlert(context.Background(), alertID)
	assert.NoError(t, err)

	// Проверка ожиданий
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
*/
