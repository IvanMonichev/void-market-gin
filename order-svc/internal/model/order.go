package model

import (
	"github.com/google/uuid"
	"time"
)

type Order struct {
	ID        uuid.UUID   `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	UserID    uuid.UUID   `gorm:"type:uuid;not null" json:"userId"`
	Products  []Product   `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE" json:"products"`
	Total     float64     `gorm:"type:numeric(10,2);not null" json:"total"`
	Status    OrderStatus `gorm:"not null" json:"status"`
	CreatedAt time.Time   `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time   `gorm:"autoUpdateTime" json:"updatedAt"`
}
