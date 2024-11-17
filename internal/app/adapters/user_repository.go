package adapters

import (
	"backend/internal/config"
	"backend/internal/models"
	"backend/internal/utils"
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {}

func (r *UserRepository) Register(user *models.User) error {
	collection := config.DB.Collection("users")

	indexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	}
	if _, err := collection.Indexes().CreateOne(context.TODO(), indexModel); err != nil {
		log.Printf("Failed to create unique index on email: %v", err)
		return errors.New("failed to create unique index on email")
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return errors.New("password hashing failed")
	}
	user.Password = hashedPassword

	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			log.Printf("Duplicate email error: %v", err)
			return errors.New("email already registered")
		}
		log.Printf("Error inserting new user: %v", err)
		return errors.New("failed to insert user")
	}

	log.Println("User registered successfully:", user.Email)
	return nil
}

func (r *UserRepository) Login(email, password string) (string, error) {
    collection := config.DB.Collection("users")
    var user models.User

    err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return "", errors.New("user not found")
        }
        return "", err
    }

    if !utils.CheckPasswordHash(password, user.Password) {
        return "", errors.New("invalid password")
    }

    token, err := utils.GenerateJWT(user.Email)
    if err != nil {
        return "", err
    }

    return token, nil
}