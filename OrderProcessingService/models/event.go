// models/order_event.go
package models

type OrderEvent struct {
	OrderID     string  `json:"orderId"`
	TotalAmount float64 `json:"totalAmount"`
}

// models/payment_event.go

type PaymentEvent struct {
	OrderID       string        `json:"orderId"`
	Status        OrderStatus   `json:"status"`
	PaymentStatus PaymentStatus `json:"paymentStatus"`
}
