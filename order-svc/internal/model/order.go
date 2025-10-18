package model

import (
	"github.com/IvanMonichev/void-market-gin/order-svc/pkg/enum"
	"time"
)

type Order struct {
	ID        uint             `gorm:"primaryKey;autoIncrement"`
	UserID    string           `gorm:"type:varchar(255);not null"`
	Status    enum.OrderStatus `gorm:"type:varchar(32);not null"`
	Total     float64          `gorm:"type:numeric(10,2);not null"`
	Items     []OrderItem      `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time        `gorm:"type:timestamp;not null;default:now()"`
	UpdatedAt time.Time        `gorm:"type:timestamp;not null;default:now()"`
}
