package model

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	OrderID   uuid.UUID `gorm:"type:uuid;not null;index" json:"orderId"`
	Name      string    `gorm:"not null" json:"name"`
	Quantity  int       `gorm:"not null" json:"quantity"`
	UnitPrice float64   `gorm:"type:numeric(10,2);not null" json:"unitPrice"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}
