package model

import (
	"github.com/google/uuid"
	"time"
)

type OrderItem struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	OrderID   uuid.UUID `gorm:"type:uuid;not null;index"`
	Name      string
	Quantity  int
	UnitPrice float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
