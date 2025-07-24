// models/order.go
package models

type OrderStatus string

const (
	Pending    OrderStatus = "PENDING"
	Success    OrderStatus = "SUCCESS"
	Processing OrderStatus = "PROCESSING"
	Canceled   OrderStatus = "CANCELED"
)

type Order struct {
	OrderId     string      `gorm:"type:char(40);primaryKey"`
	CustomerId  string      `json:"customerId" gorm:"not null;index"`
	ProductId   string      `json:"productId" gorm:"not null"`
	Quantity    int         `json:"quantity" gorm:"not null"`
	TotalAmount float64     `json:"totalAmount" gorm:"not null"`
	OrderDate   string      `json:"orderDate" gorm:"autoCreateTime"`
	Status      OrderStatus `json:"status" gorm:"type:varchar(20);default:'PROCESSING'"`
	PaymentStatus  PaymentStatus `json:"paymentStatus" gorm:"type:varchar(20);default:'PENDING'"`
}
