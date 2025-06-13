package transport

import "time"

type ItemRDO struct {
	ID       string  `json:"productId"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type OrderRDO struct {
	OrderID   string    `json:"orderId"`
	UserID    string    `json:"userId"`
	Items     []ItemRDO `json:"items"`
	Total     float64   `json:"total"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}
