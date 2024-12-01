package main

import (
	"alertservice/config"
	v1 "alertservice/controllers/http/v1"
	"alertservice/database/repo"
	"alertservice/entity"
	"alertservice/pkg/postgres"
	"alertservice/usecase"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		//log.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
		fmt.Println(err)
		return
	}
	defer pg.Close()

	db := repo.New(pg)
	u := usecase.New(cfg, db)
	server := v1.New(cfg, u)

	go func() {
		if err := server.Run(); err != nil {
			log.Fatalf("Failed to run server: %v", err)
		}
	}()

	// Создаем канал для получения сигналов
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt) // Подписываемся на сигнал прерывания (Ctrl+C)

	// Ожидаем сигнала
	<-signalChan
	log.Println("Received shutdown signal, stopping server...")

	// Останавливаем сервер
	if err := server.Stop(context.Background()); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server stopped gracefully")

	alert := entity.Alert{
		AlertID:      1,
		AlertName:    "High CPU Usage",
		Severity:     "High",
		Description:  "The CPU usage has exceeded 90% for the last 5 minutes.",
		Timestamp:    time.Now(),
		GeneratorURL: "http://monitoring.example.com",
		Status:       "active",
	}
	ctx := context.Background()
	alertId, err := db.StoreAlert(ctx, alert)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("AlertId: %v.\n", alertId)
	ctx = context.Background()
	alerts, err := db.GetAlerts(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(alerts)
}
