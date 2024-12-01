package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failedOnError(err, "failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failedOnError(err, "Failed to open a channel")
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	failedOnError(err, "Failed to open a channel")

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	failedOnError(err, "Failed to register a consumer")
	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Recieved a message: %s", d.Body)
		}
	}()

	log.Printf("......Waiting for messages. To exit press CTRL + C")
	<-forever
}

func failedOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s %s", msg, err)
	}
}
