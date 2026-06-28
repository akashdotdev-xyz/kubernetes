package service

import "github.com/akashdotdev-xyz/order-service/internal/model"

type OrderService struct{}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) GetOrders() []model.Order {
	return []model.Order{
		{
			ID:    1,
			Item:  "MacBook Pro",
			Price: 250000,
		},
		{
			ID:    2,
			Item:  "Mechanical Keyboard",
			Price: 12000,
		},
		{
			ID:    3,
			Item:  "Magic Mouse",
			Price: 8000,
		},
	}
}
