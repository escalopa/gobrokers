package main

import (
	"flag"
	"fmt"

	"github.com/streadway/amqp"
)

var queueName string

func main() {
	flag.StringVar(&queueName, "queue", "my-queue", "queue name")
	flag.Parse()

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	q, er := ch.QueueDeclare(
		queueName, // queue name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	if er != nil {
		panic(er)
	}

	msgs, err := ch.Consume(
		q.Name, // queue name
		"",     // consumer name
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {
		panic(err)
	}

	for msg := range msgs {
		fmt.Println(string(msg.Body))
	}
}
