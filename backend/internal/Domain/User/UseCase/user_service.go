package usecase

import (
	domain "backend/internal/Domain/User/Domain"
	ports "backend/internal/Domain/User/Ports"
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

func (s *UserService) GetAllUsers() ([]domain.User, error) {
    return s.UserRepo.GetAllUsers()
}

func (s *UserService) GetUserBySubOrID(sub string) (*domain.User, error) {
    return s.UserRepo.GetUserBySubOrID(sub)
}
