package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
    ID          primitive.ObjectID `bson:"_id,omitempty"`
    ChatID      primitive.ObjectID `bson:"chat_id"`
    SenderID    primitive.ObjectID `bson:"sender_id"`
    ReceiverID  primitive.ObjectID `bson:"receiver_id"`
    Content     string             `bson:"content"`
    CreatedAt   time.Time             `bson:"created_at"`
}