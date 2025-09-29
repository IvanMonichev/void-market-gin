package transport

import "time"

type OrderItemRDO struct {
	Name      string    `json:"name"`
	Quantity  int       `json:"quantity"`
	UnitPrice float64   `json:"unitPrice"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
