package main

import (
	"log"
	"strconv"
	"time"

	"github.com/streadway/amqp"
)

type User struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	defer ch.Close()

	exchangeName := "my-exchange"
	ch.ExchangeDeclare(exchangeName, "fanout", true, false, false, false, nil)

	for i := 0; i < 2; i++ {
		queueName := "my-queue" + strconv.Itoa(i)
		q, err := ch.QueueDeclare(
			queueName, // queue name
			false,     // durable
			false,     // delete when unused
			false,     // exclusive
			false,     // no-wait
			nil,       // arguments
		)
		if err != nil {
			log.Fatalf("Failed to declare a queue: %s", err)
		}
		ch.QueueBind(q.Name, "", exchangeName, false, nil)
	}

	var body string
	for i := 0; i < 10; i++ {
		body = "Hello, RabbitMQ!" + strconv.Itoa(i)
		err = ch.Publish(
			exchangeName, // exchange name
			"",           // routing key
			false,        // mandatory
			false,        // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			},
		)

		if err != nil {
			log.Fatalf("Failed to publish a message: %s", err)
		}

		log.Printf("Message sent: %s", body)
		<-time.After(2 * time.Second)
	}
}
