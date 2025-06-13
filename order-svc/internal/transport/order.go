package transport

import (
	"github.com/IvanMonichev/void-market-gin/order-svc/internal/model"
	"github.com/google/uuid"
	"time"
)

type OrderDTO struct {
	UserID     uuid.UUID         `json:"userId" binding:"required"`
	ProductIDs []uuid.UUID       `json:"productIds" binding:"required,dive,required"`
	Status     model.OrderStatus `json:"status" binding:"required"`
}
type OrderRDO struct {
	ID        uuid.UUID    `json:"id"`
	UserID    uuid.UUID    `json:"userId"`
	Status    string       `json:"status"`
	Total     float64      `json:"total"`
	Products  []ProductRDO `json:"products"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
}
