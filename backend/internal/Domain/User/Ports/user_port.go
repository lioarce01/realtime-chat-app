package ports

import (
	domain "backend/internal/Domain/User/Domain"

	"go.mongodb.org/mongo-driver/bson"
)

type UserPort interface {
    Register(user *domain.User) error
    GetUserBySubOrID(sub string) (*domain.User, error)
    GetAllUsers(filter bson.M) ([]domain.User, error)
}
