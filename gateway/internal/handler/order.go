package handler

import (
	"encoding/json"
	"fmt"
	"gateway/internal/client"
	"gateway/internal/transport"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type OrderHandler struct {
	clients *client.Clients
}

func NewOrderHandler(clients *client.Clients) *OrderHandler {
	return &OrderHandler{
		clients: clients,
	}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	var input struct {
		UserID string `json:"userId"`
	}
	if err := json.Unmarshal(body, &input); err != nil || input.UserID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userId is required"})
		return
	}

	userResp, err := h.clients.User.R().
		Get(fmt.Sprintf("/users/%s", input.UserID))
	if err != nil || userResp.StatusCode() != http.StatusOK {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid or missing user"})
		return
	}

	// Создание заказа
	orderResp, err := h.clients.Order.R().
		SetBody(body).
		Post("/orders")
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "order service unavailable"})
		return
	}

	c.Data(orderResp.StatusCode(), orderResp.Header().Get("Content-Type"), orderResp.Body())
}

func (h *OrderHandler) GetAllOrders(c *gin.Context) {
	offset := c.DefaultQuery("offset", "0")
	limit := c.DefaultQuery("limit", "10")

	resp, err := h.clients.Order.R().
		SetQueryParams(map[string]string{
			"offset": offset,
			"limit":  limit,
		}).
		Get("/orders/all")

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadGateway, gin.H{"error": "failed to fetch orders"})
		return
	}

	var parsed struct {
		Total  int                  `json:"total"`
		Orders []transport.OrderRDO `json:"orders"`
	}

	if err := json.Unmarshal(resp.Body(), &parsed); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse order data"})
		return
	}

	var fullOrders []transport.OrderWithUserRDO

	for _, order := range parsed.Orders {
		userResp, err := h.clients.User.R().
			SetResult(&transport.User{}).
			Get(fmt.Sprintf("/users/%s", order.UserID))

		var user transport.User
		if err == nil && userResp.StatusCode() == http.StatusOK {
			user = *userResp.Result().(*transport.User)
		}

		fullOrders = append(fullOrders, transport.OrderWithUserRDO{
			ID:        order.ID,
			User:      user,
			Status:    order.Status,
			Total:     order.Total,
			Items:     order.Items,
			CreatedAt: order.CreatedAt,
			UpdatedAt: order.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"total":  parsed.Total,
		"orders": fullOrders,
	})
}

func (h *OrderHandler) UpdateOrder(c *gin.Context) {
	id := c.Param("id")

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	resp, err := h.clients.Order.R().
		SetBody(body).
		Put(fmt.Sprintf("/orders/%s", id))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "order service unavailable"})
		return
	}

	c.Data(resp.StatusCode(), resp.Header().Get("Content-Type"), resp.Body())
}

func (h *OrderHandler) DeleteOrder(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.clients.Order.R().
		Delete(fmt.Sprintf("/orders/%s", id))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "order service unavailable"})
		return
	}

	c.Status(resp.StatusCode())
}

func (h *OrderHandler) GetOrder(c *gin.Context) {
	id := c.Param("id")

	order := &transport.OrderRDO{}
	orderResp, err := h.clients.Order.R().
		SetResult(&order).
		Get(fmt.Sprintf("/orders/%s", id))
	if err != nil || orderResp.StatusCode() != http.StatusOK {
		c.JSON(http.StatusBadGateway, gin.H{"error": "failed to fetch order"})
		return
	}

	userResp, err := h.clients.User.R().
		SetResult(&transport.User{}).
		Get(fmt.Sprintf("/users/%s", order.UserID))
	if err != nil || userResp.StatusCode() != http.StatusOK {
		c.JSON(http.StatusBadGateway, gin.H{"error": "failed to fetch user"})
		return
	}
	user := userResp.Result().(*transport.User)

	fullOrder := transport.OrderWithUserRDO{
		ID:        order.ID,
		User:      *user,
		Status:    order.Status,
		Total:     order.Total,
		Items:     order.Items,
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
	}
	c.JSON(http.StatusOK, fullOrder)
}
