package broker

import (
	"context"
	"encoding/json"
	"github.com/IvanMonichev/void-market-gin/order-svc/internal/repository"
	"github.com/streadway/amqp"
	"log"
	"shared/enum"
)

type StatusEvent struct {
	OrderID uint             `json:"orderId"`
	Status  enum.OrderStatus `json:"status"`
}

func StartStatusConsumer(amqpURL, queue string, repo repository.OrderRepository) error {
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return err
	}

	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	q, err := ch.QueueDeclare(queue, false, false, false, false, nil)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			var event StatusEvent
			if err := json.Unmarshal(d.Body, &event); err != nil {
				log.Println("Invalid event:", err)
				continue
			}

			log.Printf("Received event: %+v", event)

			order, err := repo.FindById(context.Background(), event.OrderID)
			if err != nil {
				log.Printf("Order not found: %d", event.OrderID)
				continue
			}

			order.Status = event.Status

			if _, err := repo.Update(context.Background(), order, order.ID); err != nil {
				log.Printf("Failed to update order: %v", err)
			}

		}
	}()

	log.Println("Listening for order status events on queue:", queue)
	return nil
}
