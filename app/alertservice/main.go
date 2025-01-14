package main

import (
	"alertservice/config"
	v1 "alertservice/internal/controllers/http/v1"
	"alertservice/internal/database/repo"
	"alertservice/internal/usecase"
	"alertservice/pkg/logger"
	"alertservice/pkg/postgres"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	_ "alertservice/docs"
)

// @title           AlertService
// @version         1.0
// @description     AlertService
// @termsOfService  http://swagger.io/terms/

// @host      localhost:8787
// @BasePath  /

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

	l := logger.New(cfg.Log.Level)
	db := repo.New(pg, l)
	u := usecase.New(cfg, db, l)
	server := v1.New(cfg, u, l)

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
}
