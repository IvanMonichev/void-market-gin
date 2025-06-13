package model

import "time"

type Order struct {
	ID         uint        `gorm:"primary_key;auto_increment" json:"id"`
	UserID     uint        `json:"user_id"`
	Items      []string    `gorm:"type:text[]" json:"items"`
	TotalPrice float32     `json:"total_price"`
	Status     OrderStatus `gorm:"type:text" json:"status"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}
