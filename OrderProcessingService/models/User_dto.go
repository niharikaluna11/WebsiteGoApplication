// models/user_dto.go
package models

type UserRegisterDTO struct {
	Name     string   `json:"name" validate:"required"`
	Email    string   `json:"email" validate:"required,email"`
	Password string   `json:"password" validate:"required,min=6"`
	Role     UserRole `json:"role" validate:"required,oneof=admin user"`
}

type UserLoginDTO struct {
	Email    string   `json:"email" validate:"required,email"`
	Password string   `json:"password" validate:"required"`
}

type UserResponseDTO struct {
	ID    string   `json:"id"`
	Name  string   `json:"name"`
	Email string   `json:"email"`
	Role  UserRole `json:"role"`
}
