package repository

import (
	"backend/config"
	domain "backend/internal/Domain/User/Domain"
	"backend/internal/utils"
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
    return &UserRepository{}
}

func (r *UserRepository) Register(user *domain.User) error {
    collection := config.DB.Collection("users")
    var existingUser domain.User
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
    var user domain.User

    err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
    if err != nil {
        return "", errors.New("invalid credentials")
    }

    if !utils.CheckPasswordHash(password, user.Password) {
        return "", errors.New("invalid credentials")
    }

    token, err := utils.GenerateJWT(user.ID.Hex(), user.Email)
    if err != nil {
        return "", err
    }

    return token, nil
}

func (r *UserRepository) GetAllUsers() ([]domain.User, error) {
    collection := config.DB.Collection("users")
    var users []domain.User

    cursor, err := collection.Find(context.TODO(), bson.M{}, options.Find())
    if err != nil {
        log.Println("Error fetching users:", err)
        return nil, err
    }
    defer cursor.Close(context.TODO())

    for cursor.Next(context.TODO()) {
        var user domain.User
        if err := cursor.Decode(&user); err != nil {
            log.Println("Error decoding user:", err)
            continue
        }
        users = append(users, user)
    }

    return users, cursor.Err()
}

func (r *UserRepository) GetUserByID(id string) (*domain.User, error) {
    collection := config.DB.Collection("users")
    var user domain.User

    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, errors.New("invalid user ID format")
    }

    err = collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&user)
    if err != nil {
        return nil, errors.New("user not found")
    }

    return &user, nil
}
