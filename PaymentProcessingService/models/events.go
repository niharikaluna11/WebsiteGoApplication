package models

type OrderCreatedEvent struct {
	OrderID     string  `json:"orderId"`
	TotalAmount float64 `json:"totalAmount"`
}

type PaymentEvent struct {
	OrderID string      `json:"orderId"`
	Status  OrderStatus `json:"status"`
}

type OrderStatus string

const (
	Pending    OrderStatus = "PENDING"
	Processing OrderStatus = "PROCESSING"
	Shipped    OrderStatus = "SHIPPED"
	Delivered  OrderStatus = "DELIVERED"
	Canceled   OrderStatus = "CANCELED"
)