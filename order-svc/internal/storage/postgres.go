package storage

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MustConnect(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect to postgres: %v", err))
	}

	return db

}
