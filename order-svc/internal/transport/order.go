package transport

import (
	"github.com/IvanMonichev/void-market-gin/order-svc/internal/domain"
	"time"
)

type OrderDTO struct {
	UserID string             `json:"userId" binding:"required"`
	Status domain.OrderStatus `json:"status" binding:"required,oneof=pending paid shipped delivery cancelled"`
	Items  []OrderItemDTO     `json:"items" binding:"required,dive"`
}

type OrderRDO struct {
	ID        uint           `json:"id"`
	UserID    uint           `json:"userId"`
	Status    string         `json:"status"`
	Total     float64        `json:"total"`
	Items     []OrderItemRDO `json:"items"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
}
