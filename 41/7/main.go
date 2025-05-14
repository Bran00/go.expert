package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	id  int64
	Msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var i int64 = 0

	// RabbitMQ

	go func() {
		for {
			atomic.AddInt64(&i, 1)
			time.Sleep(1 * time.Second)
			msg := Message{i + 1, "Hello RabbitMQ"}
			c1 <- msg
		}
	}()

	// Kafka
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			time.Sleep(1 * time.Second)
			msg := Message{i + 1, "Hello Kafka"}
			c2 <- msg
		}
	}()

	//for i := 0; i < 2; i++ {
	for {
		select {
		case msg := <-c1: //rabbitmq
			fmt.Printf("Received from RabbitMQ: %v - ID: %d\n", msg.Msg, msg.id)

		case msg := <-c2: //kafka
			fmt.Printf("Received from Kafka: %v - ID: %d\n", msg.Msg, msg.id)

		case <-time.After(3 * time.Second):
			println("Timeout")

			/* default:
			println("Default") */
		}
	}
}
