package main

import (
	"fmt"
	"github.com/IvanMonichev/void-market-gin/payment-svc/internal/handler"
	broker "github.com/IvanMonichev/void-market-gin/payment-svc/internal/rabbitmq"
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func getAMQPURLFromEnv() (string, error) {
	user := os.Getenv("RABBITMQ_USER")
	password := os.Getenv("RABBITMQ_PASSWORD")
	host := os.Getenv("RABBITMQ_HOST")
	port := os.Getenv("RABBITMQ_PORT")

	if user == "" || password == "" || host == "" || port == "" {
		return "", fmt.Errorf("missing RabbitMQ environment variables")
	}

	return fmt.Sprintf("amqp://%s:%s@%s:%s/", user, password, host, port), nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found or couldn't be loaded")
	}

	r := gin.Default()

	amqpURL, err := getAMQPURLFromEnv()
	fmt.Println(amqpURL)
	if err != nil {
		log.Fatal("Failed to get RabbitMQ config:", err)
	}
	publisher, err := broker.NewPublisher(amqpURL, "order_status_changed")
	if err != nil {
		log.Fatal("Failed to init publisher:", err)
	}
	defer publisher.Close()

	h := handler.NewPaymentHandler(publisher)
	r.POST("/payment/orders/:id/status", h.UpdateOrderStatus)

	r.Run(":4012")
}
