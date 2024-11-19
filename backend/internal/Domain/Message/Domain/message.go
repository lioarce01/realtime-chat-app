package message

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ChatID      primitive.ObjectID `bson:"chat_id" json:"chat_id"`
	SenderID    primitive.ObjectID `bson:"sender_id" json:"sender_id"`
	ReceiverID  primitive.ObjectID `bson:"receiver_id" json:"receiver_id"`
	Content     string             `bson:"content" json:"content"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
}
