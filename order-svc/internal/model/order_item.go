package model

import (
	"time"
)

type OrderItem struct {
	ID        uint      `gorm:"type:serial;primaryKey;autoIncrement"`
	OrderID   uint      `gorm:"type:integer;not null;index"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Quantity  int       `gorm:"type:integer;not null"`
	UnitPrice float64   `gorm:"type:numeric(10,2);not null"`
	CreatedAt time.Time `gorm:"type:timestamp;not null;default:now()"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null;default:now()"`
}
