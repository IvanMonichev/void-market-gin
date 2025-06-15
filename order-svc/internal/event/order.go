package event

import (
	"github.com/IvanMonichev/void-market-gin/order-svc/internal/broker"
	"github.com/IvanMonichev/void-market-gin/order-svc/internal/domain"
	"github.com/IvanMonichev/void-market-gin/order-svc/internal/model"
)

type OrderCreatedEvent struct {
	OrderID uint               `json:"orderId"`
	UserID  string             `json:"userId"`
	Amount  float64            `json:"amount"`
	Status  domain.OrderStatus `json:"status"`
}

func PublishOrderCreated(publisher *broker.Publisher, order *model.Order) error {
	event := OrderCreatedEvent{
		OrderID: order.ID,
		UserID:  order.UserID,
		Amount:  order.Total,
		Status:  order.Status,
	}

	return publisher.Publish(event)
}
