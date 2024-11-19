package usecase

import (
	domain "backend/internal/User/Domain"
	ports "backend/internal/User/Ports"
)

type UserService struct {
    UserRepo ports.UserPort
}

func NewUserService(userRepo ports.UserPort) *UserService {
    return &UserService{UserRepo: userRepo}
}

func (s *UserService) Register(user *domain.User) error {
    return s.UserRepo.Register(user)
}

func (s *UserService) Login(email, password string) (string, error) {
    return s.UserRepo.Login(email, password)
}

func (s *UserService) GetAllUsers() ([]domain.User, error) {
    return s.UserRepo.GetAllUsers()
}

func (s *UserService) GetUserByID(id string) (*domain.User, error) {
    return s.UserRepo.GetUserByID(id)
}
