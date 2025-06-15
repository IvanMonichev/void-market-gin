package main

import (
	"github.com/IvanMonichev/void-market-gin/payment-svc/internal/handler"
	"github.com/IvanMonichev/void-market-gin/payment-svc/internal/rabbitmq"
	"github.com/IvanMonichev/void-market-gin/payment-svc/internal/repository"
	"github.com/IvanMonichev/void-market-gin/payment-svc/internal/router"
	"github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
)

func main() {
	// MongoDB
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("void_market")
	repo := repository.NewMongoPaymentRepository(db)

	// RabbitMQ
	rmqConn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("RabbitMQ connection error:", err)
	}
	defer rmqConn.Close()

	publisher, err := rabbitmq.NewPublisher(rmqConn, "payment_created")
	if err != nil {
		log.Fatal("RabbitMQ publisher error:", err)
	}

	// Handler
	h := handler.NewPaymentHandler(repo, publisher)

	// Router
	r := router.SetupRouter(h)
	if err := r.Run(":4030"); err != nil {
		log.Fatal(err)
	}
}
