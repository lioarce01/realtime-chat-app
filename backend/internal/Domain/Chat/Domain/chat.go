package chat

import (
	"time"

	messageDomain "backend/internal/Domain/Message/Domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chat struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	User1ID   primitive.ObjectID `bson:"user1_id" json:"user1_id"` 
	User2ID   primitive.ObjectID `bson:"user2_id" json:"user2_id"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	Messages []messageDomain.Message          `bson:"messages" json:"messages"`
}