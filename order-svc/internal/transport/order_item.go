package transport

import "github.com/google/uuid"

type OrderItemDTO struct {
	Name      string  `json:"name" binding:"required,min=2"`
	Quantity  int     `json:"quantity" binding:"required,gte=1"`
	UnitPrice float64 `json:"unitPrice" binding:"required,gte=0"`
}

type OrderItemRDO struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Quantity  int       `json:"quantity"`
	UnitPrice float64   `json:"unitPrice"`
}
