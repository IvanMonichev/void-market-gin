package model

import (
	"github.com/IvanMonichev/void-market-gin/order-svc/internal/domain"
	"github.com/google/uuid"
	"time"
)

type Order struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID    uuid.UUID
	Status    domain.OrderStatus
	Total     float64
	Items     []OrderItem `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
