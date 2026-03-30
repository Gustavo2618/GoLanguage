package ports

import "kafkaTeste/domain"

type OrderRepository interface {
	Save(order domain.Order) error
}

type OrderConsumer interface {
	Consumer(order domain.Order) error
}
