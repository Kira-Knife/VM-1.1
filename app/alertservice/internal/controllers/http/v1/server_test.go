package v1

import (
	"fmt"
	"log"
	"testing"

	"alertservice/config"
	"alertservice/internal/database/repo"
	"alertservice/internal/usecase"
	"alertservice/pkg/logger"
	"alertservice/pkg/postgres"
)

func setupTestServer() *Server {
	cfg, _ := config.NewConfig()
	cfg.Port = "8080"
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
		fmt.Println(err)
		return nil
	}
	defer pg.Close()
	l := logger.New(cfg.Log.Level)
	db := repo.New(pg, l)
	u := usecase.New(cfg, db, l)
	return New(cfg, u, l)
}

func TestIncomingAlerts(t *testing.T) {
	/*server := setupTestServer()

	req, err := http.NewRequest(http.MethodPost, "/api/v1/alerts", bytes.NewBuffer([]byte(`{"some": "json"}`)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.incomingAlerts)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}*/
}

/*
func TestGetListAlerts(t *testing.T) {
	server := setupTestServer()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/alerts/list?begin=0&count=10", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.getListAlerts)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestGetAlertByID(t *testing.T) {
	server := setupTestServer()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/alerts/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/alerts/{alert_id:[0-9]+}", server.getAlertByID)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestHandlerPass(t *testing.T) {
	server := setupTestServer()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/settings/interval", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.handlerPass)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusServiceUnavailable {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusServiceUnavailable)
	}
}
*/
