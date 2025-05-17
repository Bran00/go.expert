package main

import rabbitmq "github.com/Bran00/go.expert/42/pkg/rabbitMQ"

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	rabbitmq.Publisher(ch, "Hello World", "amq.direct")
}