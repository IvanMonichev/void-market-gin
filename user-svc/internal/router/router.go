package router

import (
	"github.com/IvanMonichev/void-market-gin/user-svc/internal/handler"
	"github.com/IvanMonichev/void-market-gin/user-svc/internal/repository"
	"github.com/gin-gonic/gin"
)

func SetupRouter(repository repository.UserRepository) *gin.Engine {
	router := gin.Default()

	userHandler := handler.New(repository)
	api := router.Group("/users")
	api.POST("/", userHandler.Create)
	api.GET("/:id", userHandler.Find)
	api.PUT("/:id", userHandler.Update)
	api.DELETE("/:id", userHandler.Delete)
	api.GET("/all", userHandler.GetAll)

	return router
}
