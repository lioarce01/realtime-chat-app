package repository

import (
	"backend/config"
	domain "backend/internal/Domain/Message/Domain"
	ports "backend/internal/Domain/Message/Ports"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ ports.MessagePort = &MessageRepository{}

type MessageRepository struct{}

func NewMessageRepository() *MessageRepository{
	return &MessageRepository{}
}

func (r *MessageRepository) SendMessage(message *domain.Message) error {
	collection := config.DB.Collection("messages")
	message.CreatedAt = time.Now()
	_, err := collection.InsertOne(context.TODO(), message)
	return err
}

func (r *MessageRepository) GetMessagesByChatID(chatID primitive.ObjectID) ([]domain.Message, error) {
    messagesCollection := config.DB.Collection("messages")
    usersCollection := config.DB.Collection("users")

    var messages []domain.Message

    cursor, err := messagesCollection.Find(context.TODO(), bson.M{"chat_id": chatID})
    if err != nil {
        log.Printf("Error querying messages: %v", err)
        return nil, err
    }
    defer cursor.Close(context.TODO())

    if err := cursor.All(context.TODO(), &messages); err != nil {
        log.Printf("Error retrieving messages from cursor: %v", err)
        return nil, err
    }

    var enrichedMessages []domain.Message

    for _, message := range messages {
        log.Printf("Fetching Sender and Receiver for Message ID %v: SenderID %v, ReceiverID %v", 
                    message.ID, message.Sender.ID, message.Receiver.ID)

        var sender domain.UserDetail
        err := usersCollection.FindOne(context.TODO(), bson.M{"_id": message.Sender.ID}).Decode(&sender)
        if err != nil {
            if err.Error() == "mongo: no documents in result" {
                log.Printf("Sender not found for Message ID %v", message.ID)
            } else {
                log.Printf("Error retrieving sender user: %v", err)
            }
            return nil, err
        }

        var receiver domain.UserDetail
        err = usersCollection.FindOne(context.TODO(), bson.M{"_id": message.Receiver.ID}).Decode(&receiver)
        if err != nil {
            if err.Error() == "mongo: no documents in result" {
                log.Printf("Receiver not found for Message ID %v", message.ID)
            } else {
                log.Printf("Error retrieving receiver user: %v", err)
            }
            return nil, err
        }

        log.Printf("Sender: %+v", sender)
        log.Printf("Receiver: %+v", receiver)

        message.Sender = sender
        message.Receiver = receiver

        enrichedMessages = append(enrichedMessages, message)
    }

    return enrichedMessages, nil
}
