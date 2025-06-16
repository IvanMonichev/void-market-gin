package transport

import (
	"github.com/IvanMonichev/void-market-gin/order-svc/pkg/enum"
	"time"
)

type OrderDTO struct {
	UserID string           `json:"userId" binding:"required"`
	Status enum.OrderStatus `json:"status" binding:"required,oneof=pending paid shipped delivery cancelled"`
	Items  []OrderItemDTO   `json:"items" binding:"required,dive"`
}

type OrderRDO struct {
	ID        uint           `json:"id"`
	UserID    string         `json:"userId"`
	Status    string         `json:"status"`
	Total     float64        `json:"total"`
	Items     []OrderItemRDO `json:"items"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
}
