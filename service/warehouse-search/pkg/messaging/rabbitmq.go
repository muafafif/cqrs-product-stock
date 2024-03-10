package messaging

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type publisher interface {
	Publish() error
}

type subscriber interface {
	Subscribe() error
}

type RabbitMQ struct {
	source     string
	channel    *amqp.Channel
	connection *amqp.Connection
	publish    publisher
	subscribe  subscriber
}

func New(source string, publish publisher, subscribe subscriber) *RabbitMQ {
	return &RabbitMQ{
		source:    source,
		publish:   publish,
		subscribe: subscribe,
	}
}

func (r *RabbitMQ) Connect() error {
	connection, err := amqp.Dial(r.source)
	if err != nil {
		return err
	}
	defer connection.Close()
	r.connection = connection
	return nil
}

func (r *RabbitMQ) Channel() error {
	channel, err := r.connection.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()
	r.channel = channel
	return nil
}

func (r *RabbitMQ) Publish() error {
	return r.publish.Publish()
}

func (r *RabbitMQ) Subscribe() error {
	return r.subscribe.Subscribe()
}
