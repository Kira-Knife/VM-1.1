package repo_test

/*
import (
	"context"
	"testing"
	"time"

	"alertservice/database/repo"
	"alertservice/entity"
	"alertservice/pkg/postgres"

	"github.com/Masterminds/squirrel"
	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/assert"
)

// setupRepo initializes a TranslationRepo for testing with a mock database pool.
func setupRepo(mockPool pgxmock.PgxPoolIface) *repo.TranslationRepo {
	return &repo.TranslationRepo{
		Postgres: &postgres.Postgres{
			Builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
			Pool:    mockPool,
		},
	}
}

// TestGetIncidents verifies that GetIncidents retrieves all incidents correctly.
func TestGetIncidents(t *testing.T) {
	ctx := context.Background()
	mockPool, err := pgxmock.NewPool()
	assert.NoError(t, err)
	defer mockPool.Close()

	r := setupRepo(mockPool)

	expectedIncidents := []entity.Incident{
		{
			IncidentID:   1,
			AlertID:      1,
			Severity:     "High",
			Description:  "Test description",
			Status:       "Open",
			Timestamp:    time.Now(),
			GeneratorURL: "http://example.com",
		},
	}

	// Define the columns expected in the result set.
	columns := []string{"incident_id", "alert_id", "severity", "description", "status", "timestamp", "generator_url"}

	// Mock the expected query and result.
	mockPool.ExpectQuery("SELECT (.+) FROM incidents").
		WillReturnRows(pgxmock.NewRows(columns).
			AddRow(expectedIncidents[0].IncidentID, expectedIncidents[0].AlertID, expectedIncidents[0].Severity, expectedIncidents[0].Description, expectedIncidents[0].Status, expectedIncidents[0].Timestamp, expectedIncidents[0].GeneratorURL))

	// Execute the function and verify results.
	incidents, err := r.GetIncidents(ctx)
	assert.NoError(t, err)
	assert.Equal(t, expectedIncidents, incidents)
}

// TestStoreIncident verifies that StoreIncident correctly inserts a new incident.
func TestStoreIncident(t *testing.T) {
	ctx := context.Background()
	mockPool, err := pgxmock.NewPool()
	assert.NoError(t, err)
	defer mockPool.Close()

	r := setupRepo(mockPool)

	newIncident := entity.Incident{
		AlertID:      1,
		Severity:     "Medium",
		Description:  "New incident description",
		Status:       "New",
		Timestamp:    time.Now(),
		GeneratorURL: "http://example.org",
	}

	// Mock the expected insert command and its result.
	mockPool.ExpectExec("INSERT INTO incidents \\(alert_id, severity, description, status, timestamp, generator_url\\)").
		WithArgs(newIncident.AlertID, newIncident.Severity, newIncident.Description, newIncident.Status, newIncident.Timestamp, newIncident.GeneratorURL).
		WillReturnResult(pgxmock.NewResult("INSERT", 1))

	// Execute the function and verify no error occurred.
	err = r.StoreIncident(ctx, newIncident)
	assert.NoError(t, err)
}

// TestGetIncident verifies that GetIncident retrieves a specific incident by its ID.
func TestGetIncident(t *testing.T) {
	ctx := context.Background()
	mockPool, err := pgxmock.NewPool()
	assert.NoError(t, err)
	defer mockPool.Close()

	r := setupRepo(mockPool)

	expectedIncident := entity.Incident{
		IncidentID:   1,
		AlertID:      1,
		Severity:     "High",
		Description:  "Test description",
		Status:       "Open",
		Timestamp:    time.Now(),
		GeneratorURL: "http://example.com",
	}

	// Define the columns expected in the result set.
	columns := []string{"incident_id", "alert_id", "severity", "description", "status", "timestamp", "generator_url"}

	// Mock the expected query and result.
	mockPool.ExpectQuery("SELECT (.+) FROM incidents WHERE incident_id = \\$1").
		WithArgs(expectedIncident.IncidentID).
		WillReturnRows(pgxmock.NewRows(columns).
			AddRow(expectedIncident.IncidentID, expectedIncident.AlertID, expectedIncident.Severity, expectedIncident.Description, expectedIncident.Status, expectedIncident.Timestamp, expectedIncident.GeneratorURL))

	// Execute the function and verify results.
	incident, err := r.GetIncident(ctx, expectedIncident.IncidentID)
	assert.NoError(t, err)
	assert.Equal(t, expectedIncident, incident)
}

// TestUpdateIncidentStatus verifies that UpdateIncidentStatus correctly updates the status of an incident.
func TestUpdateIncidentStatus(t *testing.T) {
	ctx := context.Background()
	mockPool, err := pgxmock.NewPool()
	assert.NoError(t, err)
	defer mockPool.Close()

	r := setupRepo(mockPool)

	incidentID := 1
	newStatus := "Closed"

	// Mock the expected update command and its result.
	mockPool.ExpectExec("UPDATE incidents SET status = \\$1 WHERE incident_id = \\$2").
		WithArgs(newStatus, incidentID).
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	// Execute the function and verify no error occurred.
	err = r.UpdateIncidentStatus(ctx, incidentID, newStatus)
	assert.NoError(t, err)
}
*/
