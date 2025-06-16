package model

import (
	"github.com/IvanMonichev/void-market-gin/order-svc/pkg/enum"
	"time"
)

type Order struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	UserID    string
	Status    enum.OrderStatus
	Total     float64
	Items     []OrderItem `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
