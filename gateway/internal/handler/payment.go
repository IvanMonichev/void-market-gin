package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"net/http"
	"shared/enum"
)

type OrderStatusUpdateRequest struct {
	Status enum.OrderStatus `json:"status" binding:"required"`
}

type PaymentHandler struct {
	Client *resty.Client
}

func NewPaymentHandler(client *resty.Client) *PaymentHandler {
	return &PaymentHandler{
		Client: client,
	}
}

func (h *PaymentHandler) UpdateOrderStatus(c *gin.Context) {
	orderID := c.Param("id")

	var req OrderStatusUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	resp, err := h.Client.R().
		SetBody(req).
		SetResult(map[string]interface{}{}).
		Post(fmt.Sprintf("/orders/%s/status", orderID))

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "payment service unreachable"})
		return
	}

	c.JSON(resp.StatusCode(), resp.Result())
}
