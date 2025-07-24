package repository

import (
	"OrderProcessingService/models"
	"time"
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type UserRepoImpl struct {
	DB *gorm.DB
}

func NewUserRepoImpl(db *gorm.DB) UserRepo {
	return &UserRepoImpl{DB: db}
}

// GetUserByEmail retrieves a user by their email address
func (r *UserRepoImpl) GetUserByEmail(email string) (*models.User,
	error) {
	var user models.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil 
		}
		return nil, err
	}
	return &user, nil
}

// CreateUser creates a new user record in the database
func (r *UserRepoImpl) CreateUser(u *models.UserRegisterDTO) error {
	newUser := models.User{
		ID:       uuid.New().String(),
		Email:    u.Email,
		Password: u.Password, 
		Role:     u.Role,
		Name:     u.Name,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	if err := r.DB.Create(&newUser).Error; err != nil {
		return err
	}

	return nil
}
