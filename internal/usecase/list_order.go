package usecase

import (
	"github.com/tiagocosta/goexpert-desafio-cleanarch/internal/entity"
)

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (uc *ListOrderUseCase) Execute() ([]OrderOutputDTO, error) {
	orders, err := uc.OrderRepository.List()

	if err != nil {
		return []OrderOutputDTO{}, err
	}

	var ordersOutput []OrderOutputDTO = []OrderOutputDTO{}

	for _, order := range orders {
		dto := OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.Price + order.Tax,
		}

		ordersOutput = append(ordersOutput, dto)
	}

	return ordersOutput, nil
}
