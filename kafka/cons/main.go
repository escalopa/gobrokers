package main

import (
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama consumer:", err)
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln("Failed to close Sarama consumer:", err)
		}
	}()

	partitionConsumer, err := consumer.ConsumePartition("my-topic", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalln("Failed to start consuming:", err)
	}
	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalln("Failed to close partition consumer:", err)
		}
	}()

	for message := range partitionConsumer.Messages() {
		fmt.Printf("Received message with value %s\n", message.Value)
	}
}
