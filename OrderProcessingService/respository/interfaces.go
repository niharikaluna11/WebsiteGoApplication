package repository

import "OrderProcessingService/models"

type OrderRepository interface {
	GenerateOrder(o *models.OrderCreateDTO) (*models.Order, error)
	GetOrderById(id string) (*models.Order, error)
	GetOrders() ([]models.Order, error)
	UpdateOrderStatus(id string, status string) (*models.Order, error)
}

type UserRepo interface {
	CreateUser(u *models.UserRegisterDTO) error
	GetUserByEmail(email string) (*models.User, error)
}
