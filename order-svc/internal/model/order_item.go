package model

import (
	"time"
)

type OrderItem struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	OrderID   uint `gorm:"not null;index"`
	Name      string
	Quantity  int
	UnitPrice float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
