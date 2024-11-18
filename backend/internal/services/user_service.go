package services

import (
	"backend/internal/app/ports"
	"backend/internal/models"
)

type UserService struct {
	UserRepo ports.UserPort
}

func NewUserService(userRepo ports.UserPort) *UserService {
	return &UserService{UserRepo: userRepo}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.UserRepo.GetAllUsers()
}

func (s *UserService) GetUserByID(id string) (*models.User, error) {
	return s.UserRepo.GetUserByID(id)
}
