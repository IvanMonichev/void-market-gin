package broker

import (
	"encoding/json"
	"github.com/streadway/amqp"
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
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
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

func (p *Publisher) Publish(message interface{}) error {
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = p.channel.Publish(
		"",           // exchange
		p.queue.Name, // routing key
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	return err
}

func (p *Publisher) Close() {
	p.channel.Close()
	p.conn.Close()
}
