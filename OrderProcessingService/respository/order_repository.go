package repository

import (
	"OrderProcessingService/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderRepoImpl struct {
	DB *gorm.DB
}

func NewOrderRepoImpl(db *gorm.DB) OrderRepository {
	return &OrderRepoImpl{DB: db}
}

// GenerateOrder creates a new order record in the database
func (r *OrderRepoImpl) GenerateOrder(dto *models.OrderCreateDTO) (*models.Order, error) {
	orderID := "ORD-" + uuid.New().String()
	newOrder := models.Order{
		OrderId:     orderID,
		CustomerId:  dto.CustomerId,
		ProductId:   dto.ProductId,
		Quantity:    dto.Quantity,
		Status:      models.Pending,
		TotalAmount: dto.TotalAmount,
		OrderDate:   time.Now().Format("2006-01-02 15:04:05"),
	}

	if err := r.DB.Create(&newOrder).Error; err != nil {
		return nil, err
	}
	return &newOrder, nil
}

// GetOrderById fetches an order using the order ID
func (r *OrderRepoImpl) GetOrderById(id string) (*models.Order, error) {
	var order models.Order
	err := r.DB.Where("order_id = ?", id).First(&order).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // not an error, just not found
		}
		return nil, err
	}
	return &order, nil
}

// GetOrders retrieves all orders from the database
func (r *OrderRepoImpl) GetOrders() ([]models.Order, error) {
	var orders []models.Order
	err := r.DB.Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

// UpdateOrderStatus updates the status of a specific order
func (r *OrderRepoImpl) UpdateOrderStatus(id string, status string, paymentStatus string) (*models.Order, error) {
	order, err := r.GetOrderById(id)
	if order == nil {
		return nil, gorm.ErrRecordNotFound
	}

	if paymentStatus == "" {
		paymentStatus = string(order.PaymentStatus)
	}

	err = r.DB.Model(order).Updates(map[string]interface{}{
		"status":         status,
		"payment_status": paymentStatus,
	}).Error

	if err != nil {
		return nil, err
	}

	if err := r.DB.Where("order_id = ?", id).First(order).Error; err != nil {
		return nil, err
	}

	return order, nil
}
