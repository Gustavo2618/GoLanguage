package service

import (
	"fmt"
	"kafkaTeste/domain"
	"kafkaTeste/ports"
)

type OrderService struct {
	repo ports.OrderRepository
}

// construtor
func NewOrderService(repo ports.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) Process(order domain.Order) error {
	if order.Quantity <= 0 {
		return fmt.Errorf("invalid Quantity")
	}

	fmt.Println("Processing Order: ", order.ID)

	return s.repo.Save(order)
}


