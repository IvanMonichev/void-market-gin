package transport

import "time"

type OrderRDO struct {
	ID        uint           `json:"id"`
	UserID    string         `json:"userId"`
	Status    string         `json:"status"`
	Total     float64        `json:"total"`
	Items     []OrderItemRDO `json:"items"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
}

type OrderWithUserRDO struct {
	ID        uint           `json:"id"`
	User      *User          `json:"user,omitempty"`
	Status    string         `json:"status"`
	Total     float64        `json:"total"`
	Items     []OrderItemRDO `json:"items"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
}
