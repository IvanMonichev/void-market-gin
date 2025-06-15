package router

import (
	"gateway/internal/handler"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func RegisterUserRoutes(r *gin.Engine, client *resty.Client) {
	h := handler.NewUserHandler(client)

	userGroup := r.Group("/api/users")
	{
		userGroup.POST("/", h.CreateUser)
		userGroup.GET("/:id", h.GetUser)
		userGroup.PUT("/:id", h.UpdateUser)
		userGroup.DELETE("/:id", h.DeleteUser)
		userGroup.GET("/all", h.GetAllUsers)
	}
}
