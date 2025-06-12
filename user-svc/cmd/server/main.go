package main

import (
	"fmt"
	"github.com/IvanMonichev/void-market-gin/user-svc/internal/app"
	"github.com/IvanMonichev/void-market-gin/user-svc/internal/config"
	"github.com/IvanMonichev/void-market-gin/user-svc/internal/repository"
	"github.com/IvanMonichev/void-market-gin/user-svc/internal/service"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg.Mongo)

	client, err := mongo.Connect(options.Client().ApplyURI(cfg.Mongo.URI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	db := client.Database("void_market_users")

	repo := repository.NewMongoUserRepository(db.Collection("users"))
	svc := service.New(repo)
	router := app.SetupRouter(svc)

	router.Run(":4010")
}
