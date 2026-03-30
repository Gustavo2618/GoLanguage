package repository

import (
	"fmt"
	"kafkaTeste/domain"
)

type inMemoryRepository struct{}

func NewInMemoryRepository() *inMemoryRepository {
	return &inMemoryRepository{}
}

func (r *inMemoryRepository) Save(order domain.Order) error {
	fmt.Println("Saving order: ", order)
	return nil
}
