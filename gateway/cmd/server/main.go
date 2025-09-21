package main

import (
	"fmt"
	"gateway/internal/client"
	"gateway/internal/config"
	"gateway/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	cfg := config.MustLoad()

	fmt.Println(cfg)

	urls := client.URLs{
		UserURL:    cfg.Services.User,
		OrderURL:   cfg.Services.Order,
		PaymentURL: cfg.Services.Payment,
	}
	clients := client.NewClients(urls)

	router.RegisterUserRoutes(r, clients.User)
	router.RegisterOrderRouter(r, clients)
	router.RegisterPaymentRouter(r, clients)

	r.Run(":4010")
}
