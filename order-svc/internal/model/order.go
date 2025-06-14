package model

import (
	"github.com/IvanMonichev/void-market-gin/order-svc/internal/domain"
	"time"
)

type Order struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	UserID    string
	Status    domain.OrderStatus
	Total     float64
	Items     []OrderItem `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
