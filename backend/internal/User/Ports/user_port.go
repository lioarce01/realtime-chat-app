package ports

import domain "backend/internal/User/Domain"

type UserPort interface {
    Register(user *domain.User) error
    Login(email, password string) (string, error)
    GetUserByID(id string) (*domain.User, error)
    GetAllUsers() ([]domain.User, error)
}
