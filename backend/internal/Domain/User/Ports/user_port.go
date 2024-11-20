package ports

import domain "backend/internal/Domain/User/Domain"

type UserPort interface {
    Register(user *domain.User) error
    GetUserBySubOrID(sub string) (*domain.User, error)
    GetAllUsers() ([]domain.User, error)
}
