// models/order_dto.go
package models

type OrderCreateDTO struct {
	CustomerId  string  `json:"customerId" validate:"required"`
	ProductId   string  `json:"productId" validate:"required"`
	Quantity    int     `json:"quantity" validate:"required,min=1"`
	TotalAmount float64 `json:"totalAmount" validate:"required,gt=0"`
}

type OrderStatusUpdateDTO struct {
	OrderId string      `json:"orderId" validate:"required"`
	Status  OrderStatus `json:"status" validate:"required,oneof=PENDING PROCESSING SUCCESS CANCELED"`
}
