package main

import (
	"fmt"
	"github.com/IvanMonichev/void-market-gin/order-svc/internal/config"
	"github.com/IvanMonichev/void-market-gin/order-svc/internal/handler"
	"github.com/IvanMonichev/void-market-gin/order-svc/internal/repository"
	"github.com/IvanMonichev/void-market-gin/order-svc/internal/router"
	"github.com/IvanMonichev/void-market-gin/order-svc/internal/storage"
	"log"
)

func main() {
	cfg := config.MustLoad()

	db := storage.MustConnect(cfg.Postgres.DSN)
	storage.AutoMigrate(db)

	orderRepo := repository.NewGormOrderRepository(db)
	orderHandler := handler.NewOrderHandler(orderRepo)

	r := router.SetupRouter(orderHandler)

	address := fmt.Sprintf("%s:%s", cfg.Server.Address, cfg.Server.Port)
	if err := r.Run(address); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
