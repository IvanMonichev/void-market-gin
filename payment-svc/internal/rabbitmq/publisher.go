package broker

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"time"
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

	q, err := ch.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	return &Publisher{conn: conn, channel: ch, queue: q}, nil
}

func (p *Publisher) Publish(msg interface{}) error {
	// Включаем режим подтверждений (делается один раз на канал)
	if err := p.channel.Confirm(false); err != nil {
		return fmt.Errorf("failed to enable confirm mode: %w", err)
	}

	// Канал для получения подтверждения публикации
	confirmChan := p.channel.NotifyPublish(make(chan amqp.Confirmation, 1))

	// Сериализуем сообщение
	body, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	// Отправляем сообщение
	err = p.channel.Publish(
		"",           // exchange
		p.queue.Name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         body,
			DeliveryMode: amqp.Persistent, // делаем сообщение устойчивым
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	// Ожидаем подтверждения от брокера
	select {
	case confirm := <-confirmChan:
		if confirm.Ack {
			return nil // всё ок
		}
		return fmt.Errorf("message not acknowledged by broker (nack received)")
	case <-time.After(5 * time.Second):
		return fmt.Errorf("publish confirmation timeout")
	}
}

func (p *Publisher) Close() {
	p.channel.Close()
	p.conn.Close()
}
