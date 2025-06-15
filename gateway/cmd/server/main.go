package main

import (
	"gateway/internal/client"
	"gateway/internal/config"
	"gateway/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	clients := client.NewClients(config.UserServiceBaseURL, config.OrderServiceBaseURL)

	router.RegisterUserRoutes(r, clients.User)
	router.RegisterOrderRouter(r, clients)

	r.Run(":4000")
}
