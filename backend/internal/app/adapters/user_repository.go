package adapters

import (
	"backend/internal/config"
	"backend/internal/models"
	"backend/internal/utils"
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

type UserRepository struct {}

func (r *UserRepository) Register(user *models.User) error {
	collection := config.DB.Collection("users")
	var existingUser models.User
	err := collection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		return errors.New("email already registered")
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Println("Error hashing password:", err)
		return err
	}
	user.Password = hashedPassword

	_, err = collection.InsertOne(context.TODO(), user)
	return err
}

func (r *UserRepository) Login(email, password string) (string, error) {
	collection := config.DB.Collection("users")
	var user models.User

	err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateJWT(user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
