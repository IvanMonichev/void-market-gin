package main

import (
	"fmt"
	"gateway/internal/config"
	"gateway/internal/router"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func main() {
	r := gin.Default()

	// Инициализируем resty клиента
	restyClient := resty.New().
		SetBaseURL(config.UserServiceBaseURL).
		SetHeader("Content-Type", "application/json").
		OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
			fmt.Printf("[OUT] %s %s\n", req.Method, req.URL)
			return nil
		}).
		OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
			fmt.Printf("[IN ] %d %s\n", resp.StatusCode(), resp.Request.URL)
			return nil
		})

	// Регистрируем маршруты
	router.RegisterUserRoutes(r, restyClient)

	// Запускаем сервер
	r.Run(":4000")
}
