package model

import "time"

type PaymentStatus string

const (
	StatusPending PaymentStatus = "pending"
	StatusPaid    PaymentStatus = "paid"
	StatusFailed  PaymentStatus = "failed"
)

type Payment struct {
	ID        string        `bson:"_id,omitempty" json:"id"`
	OrderID   string        `bson:"orderId" json:"orderId"`
	UserID    string        `bson:"userId" json:"userId"`
	Amount    float64       `bson:"amount" json:"amount"`
	Status    PaymentStatus `bson:"status" json:"status"`
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time     `bson:"updatedAt" json:"updatedAt"`
}
