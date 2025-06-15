package router

import (
	"gateway/internal/client"
	"gateway/internal/handler"
	"github.com/gin-gonic/gin"
)

func RegisterOrderRouter(r *gin.Engine, clients *client.Clients) {
	h := handler.NewOrderHandler(clients)

	order := r.Group("/api/orders")
	{
		order.POST("/", h.CreateOrder)
		order.GET("/:id", h.GetOrder)
		order.PUT("/:id", h.UpdateOrder)
		order.DELETE("/:id", h.DeleteOrder)
		order.GET("/all", h.GetAllOrders)
	}
}
