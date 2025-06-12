package app

import (
	"github.com/IvanMonichev/void-market-gin/user-svc/internal/handler"
	"github.com/IvanMonichev/void-market-gin/user-svc/internal/service"
	"github.com/gin-gonic/gin"
)

func SetupRouter(service service.UserService) *gin.Engine {
	router := gin.Default()

	userHandler := handler.New(service)
	api := router.Group("/api/users")
	api.POST("/", userHandler.Create)
	api.GET("/:id", userHandler.Find)
	api.PUT("/:id", userHandler.Update)
	api.DELETE("/:id", userHandler.Delete)

	return router
}
