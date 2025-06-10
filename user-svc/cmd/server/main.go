package main

import (
	"github.com/IvanMonichev/void-market-gin/user-svc/internal/app"
	"github.com/IvanMonichev/void-market-gin/user-svc/internal/repository"
	"github.com/IvanMonichev/void-market-gin/user-svc/internal/service"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
)

func main() {

	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://root:password@localhost:27019"))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	db := client.Database("void_market_users")

	repo := repository.NewMongoUserRepository(db.Collection("users"))
	svc := service.New(repo)
	router := app.SetupRouter(svc)

	router.Run(":4010")
}
