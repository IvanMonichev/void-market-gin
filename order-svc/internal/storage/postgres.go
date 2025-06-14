package storage

import (
	"fmt"
	"github.com/IvanMonichev/void-market-gin/order-svc/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func MustConnect(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect to postgres: %v", err))
	}

	return db

}

func AutoMigrate(db *gorm.DB) {
	if err := db.AutoMigrate(&model.Order{}, &model.OrderItem{}); err != nil {
		log.Fatalf("auto migration failed: %v", err)
	}
}
