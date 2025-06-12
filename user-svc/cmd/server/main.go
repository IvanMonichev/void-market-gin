package main

import (
	"fmt"
	"github.com/IvanMonichev/void-market-gin/user-svc/internal/app"
	"github.com/IvanMonichev/void-market-gin/user-svc/internal/config"
	"github.com/IvanMonichev/void-market-gin/user-svc/internal/repository"
	"github.com/IvanMonichev/void-market-gin/user-svc/internal/service"
	"github.com/IvanMonichev/void-market-gin/user-svc/internal/storage"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg.Mongo)

	db := storage.MustConnect(storage.MongoConfig{
		URI:      cfg.Mongo.URI,
		Database: cfg.Mongo.Database,
		Timeout:  cfg.Mongo.Timeout,
	})

	repo := repository.NewMongoUserRepository(db.Collection("users"))
	svc := service.New(repo)
	router := app.SetupRouter(svc)

	router.Run(":4010")
}
