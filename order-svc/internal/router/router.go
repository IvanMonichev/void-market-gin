package router

import (
	"github.com/IvanMonichev/void-market-gin/order-svc/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(h *handler.OrderHandler) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")

	orders := api.Group("/orders")
	{
		orders.POST("", h.Create)
		orders.GET("", h.GetAll)
		orders.GET("/:id", h.Find)
		orders.PUT("/:id", h.Update)
		orders.DELETE("/:id", h.Delete)
	}

	return r
}
