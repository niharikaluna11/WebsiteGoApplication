// models/user.go
package models

type UserRole string

const (
	Admin    UserRole = "ADMIN"
	Customer UserRole = "CUSTOMER"
)

type User struct {
	ID       string   `gorm:"type:char(36);primaryKey" json:"id"`
	Name     string   `gorm:"type:varchar(100);not null" json:"name"`
	Email    string   `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	Password string   `gorm:"type:varchar(255);not null" json:"-"`
	Role     UserRole `gorm:"type:varchar(20);default:'CUSTOMER'" json:"role"`
	CreatedAt string  `gorm:"autoCreateTime" json:"createdAt"`
}
