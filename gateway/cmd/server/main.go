package main

import (
	"gateway/internal/client"
	"gateway/internal/config"
	"gateway/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	urls := client.URLs{
		UserURL:    config.UserServiceBaseURL,
		OrderURL:   config.OrderServiceBaseURL,
		PaymentURL: config.PaymentServiceBaseURL,
	}
	clients := client.NewClients(urls)

	router.RegisterUserRoutes(r, clients.User)
	router.RegisterOrderRouter(r, clients)

	r.Run(":4000")
}
