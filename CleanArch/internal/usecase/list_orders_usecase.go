package usecase

import (
	"log"

	"github.com/devfullcycle/20-CleanArch/internal/entity"
)

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(OrderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (uc *ListOrdersUseCase) Execute() ([]OrderOutputDTO, error) {
	log.Println("Executing ListOrdersUseCase")
	orders, err := uc.OrderRepository.List()
	if err != nil {
		return nil, err
	}

	var result []OrderOutputDTO
	for _, order := range orders {
		dto := OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.Price + order.Tax,
		}
		result = append(result, dto)
	}

	return result, nil
}
