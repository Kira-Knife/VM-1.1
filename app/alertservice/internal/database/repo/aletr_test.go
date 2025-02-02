package repo

import (
	"alertservice/config"
	"alertservice/pkg/logger"
	"alertservice/pkg/postgres"
	"fmt"
	"log"
	"testing"
)

func TestMain(m *testing.M) {
	var err error

	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	// docs.SwaggerInfo.Host = cfg.HTTP.host

	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
		fmt.Println(err)
		return
	}
	defer pg.Close()

	l := logger.New(cfg.Log.Level)
	db := New(pg, l)
	// Создание репозитория

	// Запуск тестов
	db.db.Close()
}

func TestStoreAlert(t *testing.T) {
	// alert := entity.Alert{
	// 	AlertName:    "Test Alert",
	// 	Severity:     "High",
	// 	Description:  "Test Description",
	// 	CreateAt:     time.Now(),
	// 	GeneratorURL: "http://example.com",
	// 	Status:       "Active",
	// 	Job:          "TestJob",
	// 	Service:      "TestService",
	// 	Instance:     "TestInstance",
	// 	StartsAt:     time.Now(),
	// 	EndsAt:       time.Now(),
	// }

	// // Вставка алерта
	// alertID, err := repo.StoreAlert(testCtx, alert)
	// if err != nil {
	// 	t.Fatalf("Failed to store alert: %v", err)
	// }

	// // Проверка, что alertID больше 0
	// if alertID <= 0 {
	// 	t.Errorf("Expected alertID to be greater than 0, got %d", alertID)
	// }
}

func TestDeleteAlert(t *testing.T) {
	// Удаление тестового алерта
	// err := repo.DeleteAlert(testCtx, 1) // Используйте корректный alertID
	// if err != nil {
	// 	t.Fatalf("Failed to delete alert: %v", err)
	// }
}
