package main

import (
	"alertservice/config"
	"alertservice/database/repo"
	"alertservice/entity"
	"alertservice/pkg/postgres"
	"context"
	"fmt"
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

	tr := repo.New(pg)
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
	err = tr.StoreAlert(ctx, alert)
	if err != nil {
		fmt.Println(err)
	}
	ctx = context.Background()
	alerts, err := tr.GetAlerts(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(alerts)
}
