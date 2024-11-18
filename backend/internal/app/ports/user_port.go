package ports

import "backend/internal/models"

type UserPort interface {
    Register(user *models.User) error
    Login(email, password string) (string, error) 
}
