package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func ConnectConsumer(brokersUrl []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	conn, err := sarama.NewConsumer(brokersUrl, config)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func ReadMessageFromQueue(topic string) (string, error) {
	brokersUrl := []string{"kafkahost1:9092", "kafkahost2:9092"}

	worker, err := ConnectConsumer(brokersUrl)
	if err != nil {
		return "", err
	}

	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		return "", err
	}

	msg := <-consumer.Messages()

	fmt.Printf("Received message Topic(%s) | Message(%s) \n", string(msg.Topic), string(msg.Value))

	return string(msg.Value), nil
}
