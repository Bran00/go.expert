package main

import (
	"fmt"

	rabbitmq "github.com/Bran00/go.expert/42/pkg/rabbitMQ"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgs := make(chan amqp.Delivery)

	go rabbitmq.Consume(ch, msgs)
	
	for msg := range msgs {
		fmt.Println(string(msg.Body))
		msg.Ack(false)
	}
}