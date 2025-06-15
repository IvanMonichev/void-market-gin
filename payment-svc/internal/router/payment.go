package router

import (
	"github.com/IvanMonichev/void-market-gin/payment-svc/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(paymentHandler *handler.PaymentHandler) *gin.Engine {
	router := gin.Default()

	payments := router.Group("/payments")
	{
		payments.POST("/", paymentHandler.CreatePayment)
		payments.GET("/order/:orderId", paymentHandler.GetPaymentByOrderID)
	}

	return router
}
