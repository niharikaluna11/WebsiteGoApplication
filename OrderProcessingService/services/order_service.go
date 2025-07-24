package services

import (
	"OrderProcessingService/models"
	repository "OrderProcessingService/respository"
)

type OrderServiceInterface interface {
	GenerateOrder(*models.OrderCreateDTO) (*models.Order, error)
	GetOrderById(string) (*models.Order, error)
	GetOrders() ([]models.Order, error)
	UpdateOrderStatus(string, string) (*models.Order, error)
}

type OrderService struct {
	Repo repository.OrderRepository
}

func (s *OrderService) GenerateOrder(dto *models.OrderCreateDTO) (*models.Order, error) {
	return s.Repo.GenerateOrder(dto)
}

func (s *OrderService) GetOrderById(id string) (*models.Order, error) {
	return s.Repo.GetOrderById(id)
}

func (s *OrderService) GetOrders() ([]models.Order, error) {
	return s.Repo.GetOrders()
}

func (s *OrderService) UpdateOrderStatus(id, status string) (*models.Order, error) {
	return s.Repo.UpdateOrderStatus(id, status)
}
