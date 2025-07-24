package services

import (
	"OrderProcessingService/auth"
	"OrderProcessingService/models"
	repository "OrderProcessingService/respository"
	"errors"
)

type UserService struct {
	repo repository.UserRepo
}

func NewUserService(r repository.UserRepo) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) Register(dto *models.UserRegisterDTO) error {

	existingUser, err := s.repo.GetUserByEmail(dto.Email)
	if err != nil {
		return err
	}

	if existingUser != nil {
		return errors.New("user already exists")
	}

	user := &models.UserRegisterDTO{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
		Role:     models.Customer,
	}

	return s.repo.CreateUser(user)
}

func (s *UserService) Login(dto *models.UserLoginDTO) (string, error) {

	user, err := s.repo.GetUserByEmail(dto.Email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := auth.GenerateJWT(user.ID, string(user.Role))
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil

}
