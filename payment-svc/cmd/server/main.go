package main

import (
	"github.com/IvanMonichev/void-market-gin/payment-svc/internal/handler"
	broker "github.com/IvanMonichev/void-market-gin/payment-svc/internal/rabbitmq"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	publisher, err := broker.NewPublisher("amqp://guest:guest@localhost:5672/", "order_status_changed")
	if err != nil {
		log.Fatal("Failed to init publisher:", err)
	}
	defer publisher.Close()

	h := handler.NewPaymentHandler(publisher)
	r.POST("/payment/orders/:id/status", h.UpdateOrderStatus)

	r.Run(":4030")
}
