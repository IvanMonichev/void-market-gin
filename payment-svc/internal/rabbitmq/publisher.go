package rabbitmq

import (
	"encoding/json"
	"github.com/IvanMonichev/void-market-gin/payment-svc/internal/internal/model"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

func NewPublisher(amqpURL, queueName string) (*Publisher, error) {
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &Publisher{
		conn:    conn,
		channel: ch,
		queue:   q,
	}, nil
}

func (p *Publisher) Publish(payment model.Payment) error {
	body, err := json.Marshal(payment)
	if err != nil {
		return err
	}

	err = p.channel.Publish(
		"",
		p.queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return err
	}

	log.Printf("[MQ OUT] published payment for orderID=%s", payment.OrderID)
	return nil
}

func (p *Publisher) Close() {
	_ = p.channel.Close()
	_ = p.conn.Close()
}
