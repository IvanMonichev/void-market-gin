package router

import (
	"github.com/IvanMonichev/void-market-gin/order-svc/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(h *handler.OrderHandler) *gin.Engine {
	r := gin.Default()

	api := r.Group("/orders")

	{
		api.POST("", h.Create)
		api.GET("/all", h.GetAll)
		api.GET("/:id", h.Find)
		api.PUT("/:id", h.Update)
		api.DELETE("/:id", h.Delete)
	}

	return r
}
