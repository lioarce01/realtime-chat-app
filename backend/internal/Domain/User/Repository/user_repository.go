package repository

import (
	"backend/config"
	chatDomain "backend/internal/Domain/Chat/Domain"
	domain "backend/internal/Domain/User/Domain"
	ports "backend/internal/Domain/User/Ports"
	"context"
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ ports.UserPort = &UserRepository{
    
}

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

    _, err = collection.InsertOne(context.TODO(), user)
    return err
}

func (r *UserRepository) GetAllUsers(filter bson.M) ([]domain.User, error) {
    collection := config.DB.Collection("users")
    var users []domain.User

    cursor, err := collection.Find(context.TODO(), filter)
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

func (r *UserRepository) GetUserBySubOrID(identifier string) (*domain.User, error) {
    userCollection := config.DB.Collection("users")
    chatCollection := config.DB.Collection("chats")

    var user domain.User

    objectID, err := primitive.ObjectIDFromHex(identifier)
    var filter bson.M

    if err == nil {
        filter = bson.M{"_id": objectID}
    } else {
        filter = bson.M{"sub": identifier}
    }

    err = userCollection.FindOne(context.TODO(), filter).Decode(&user)
    if err != nil {
        return nil, fmt.Errorf("user not found")
    }

    chatFilter := bson.M{
        "$or": []bson.M{
            {"user1_id": user.ID},
            {"user2_id": user.ID},
        },
    }

    cursor, err := chatCollection.Find(context.TODO(), chatFilter)
    if err != nil {
        return nil, fmt.Errorf("error fetching chats: %v", err)
    }
    defer cursor.Close(context.TODO())

    var chatIDs []chatDomain.Chat
    for cursor.Next(context.TODO()) {
        var chat chatDomain.Chat
        if err := cursor.Decode(&chat); err != nil {
            return nil, fmt.Errorf("error decoding chat: %v", err)
        }
        chatIDs = append(chatIDs, chat)
    }

    user.Chats = chatIDs

    return &user, nil
}
