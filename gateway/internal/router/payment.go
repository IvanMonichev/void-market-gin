package router

import (
	"gateway/internal/client"
	"gateway/internal/handler"
	"github.com/gin-gonic/gin"
)

func RegisterPaymentRouter(r *gin.Engine, clients *client.Clients) {
	h := handler.NewPaymentHandler(clients.Payment)

	payment := r.Group("/api/payment")
	{
		payment.POST("/orders/:id/status", h.UpdateOrderStatus)
	}
}
