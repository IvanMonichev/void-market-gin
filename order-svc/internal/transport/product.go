package transport

import "github.com/google/uuid"

type ProductDTO struct {
	ProductID uuid.UUID `json:"productId" binding:"required"`
	Name      string    `json:"name" binding:"required"`
	Quantity  int       `json:"quantity" binding:"required,gte=1"`
	UnitPrice float64   `json:"unitPrice" binding:"required,gte=0"`
}

type ProductRDO struct {
	ID        uuid.UUID `json:"id"`
	OrderID   uuid.UUID `json:"orderId"`
	Name      string    `json:"name"`
	Quantity  int       `json:"quantity"`
	UnitPrice float64   `json:"unitPrice"`
}
