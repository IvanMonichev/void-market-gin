package handler

import (
	"context"
	"github.com/IvanMonichev/void-market-gin/payment-svc/internal/internal/model"
	"github.com/IvanMonichev/void-market-gin/payment-svc/internal/mq"
	"github.com/IvanMonichev/void-market-gin/payment-svc/internal/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type PaymentHandler struct {
	repo      repository.PaymentRepository
	publisher *mq.Publisher
}

func NewPaymentHandler(repo repository.PaymentRepository, publisher *mq.Publisher) *PaymentHandler {
	return &PaymentHandler{
		repo:      repo,
		publisher: publisher,
	}
}

func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	var payment model.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payment payload"})
		return
	}

	payment.Status = model.StatusPending
	payment.CreatedAt = time.Now()

	if err := h.repo.Save(c.Request.Context(), &payment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save payment"})
		return
	}

	// Публикуем сообщение в очередь RabbitMQ
	if err := h.publisher.Publish(payment); err != nil {
		// Логируем, но не прерываем выполнение
		c.JSON(http.StatusAccepted, gin.H{"warning": "payment saved but not published", "payment": payment})
		return
	}

	c.JSON(http.StatusCreated, payment)
}

func (h *PaymentHandler) GetPaymentByOrderID(c *gin.Context) {
	orderID := c.Param("orderId")

	payment, err := h.repo.FindByOrderID(context.Background(), orderID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "payment not found"})
		return
	}

	c.JSON(http.StatusOK, payment)
}
