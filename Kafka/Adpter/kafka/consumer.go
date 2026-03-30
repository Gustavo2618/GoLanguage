package kafka

import (
	"context"
	"encoding/json"
	"kafkaTeste/domain"
	"kafkaTeste/service"
	"log"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	reader  *kafka.Reader
	service *service.OrderService
}

func NewKafkaConsumer(brokers []string, topic string, svc *service.OrderService) *KafkaConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: "order-group",
	})

	return &KafkaConsumer{
		reader:  reader,
		service: svc,
	}
}

func (c *KafkaConsumer) Start() {
	for {
		msg, err := c.reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("error reading message: ", err)
			continue
		}

		var order domain.Order
		err = json.Unmarshal(msg.Value, &msg.Value)
		if err != nil {
			log.Println("error parsing message: ", err)
			continue
		}

		err = c.service.Process(order)
		if err != nil {
			log.Panicln("error processing order: ", err)
		}
	}
}
