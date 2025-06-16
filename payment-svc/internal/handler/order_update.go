package handler

import (
	broker "github.com/IvanMonichev/void-market-gin/payment-svc/internal/rabbitmq"
	"github.com/IvanMonichev/void-market-gin/payment-svc/pkg/enum"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderStatusUpdateRequest struct {
	Status enum.OrderStatus `json:"status" binding:"required"`
}

type PaymentHandler struct {
	publisher *broker.Publisher
}

func NewPaymentHandler(publisher *broker.Publisher) *PaymentHandler {
	return &PaymentHandler{
		publisher: publisher,
	}
}

func (h *PaymentHandler) UpdateOrderStatus(c *gin.Context) {
	orderIDParam := c.Param("id")
	orderID, err := strconv.ParseUint(orderIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order ID"})
		return
	}

	var req OrderStatusUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if !req.Status.IsValid() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid status value"})
		return
	}

	message := map[string]interface{}{
		"orderId": orderID,
		"status":  req.Status,
	}

	if err := h.publisher.Publish(message); err != nil {
		log.Println("failed to publish status update:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to publish event"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "status update published"})
}
