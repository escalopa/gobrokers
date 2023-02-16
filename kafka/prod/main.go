package main

import (
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln("Failed to close Sarama producer:", err)
		}
	}()

	message := &sarama.ProducerMessage{
		Topic: "my-topic",
		Value: sarama.StringEncoder("Hello, Kafka!"),
	}

	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		log.Fatalln("Failed to send message:", err)
	}

	fmt.Printf("Message sent to partition %d at offset %d\n", partition, offset)
}
