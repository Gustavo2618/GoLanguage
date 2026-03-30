package main

import (
	"kafkaTeste/Adpter/kafka"
	"kafkaTeste/Adpter/repository"
	"kafkaTeste/service"
)

func main() {

	repo := repository.NewInMemoryRepository()
	service := service.NewOrderService(repo)

	consumer := kafka.NewKafkaConsumer(
		[]string{"localhost:9092"},
		"orders.created",
		service,
	)
	consumer.Start()
}
