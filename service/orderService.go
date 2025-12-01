package service

import (
	"fmt"

	"github.com/davidpogosian/pickupee/platform/models"
	"github.com/davidpogosian/pickupee/platform/repository"
)

type OrderService struct {
	repository *repository.OrderRepository
}

func NewOrderService(repository *repository.OrderRepository) *OrderService {
	return &OrderService{repository: repository}
}

func (s *OrderService) PlaceOrder(userID int, itemIDs []int) (int, error) {
	orderID, err := s.repository.CreateOrder(userID, itemIDs)
	if err != nil {
		return -1, fmt.Errorf("Couldn't place order: %w", err)
	}
	return orderID, nil
}

func (s *OrderService) ListOrdersForUser(userID int) ([]models.Order, error) {
	orders, err := s.repository.ListByUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("Couldn't list orders for user %d: %w", userID, err)
	}
	return orders, nil
}
