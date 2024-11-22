package chat

import (
	"time"

	messageDomain "backend/internal/Domain/Message/Domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chat struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	User1ID     primitive.ObjectID   `bson:"user1_id" json:"user1_id"`
	User2ID     primitive.ObjectID   `bson:"user2_id" json:"user2_id"`
	User1     *UserDetail        `bson:"-" json:"user1"` 
	User2     *UserDetail        `bson:"-" json:"user2"` 
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	Messages  []messageDomain.Message `bson:"messages" json:"messages"`
	LastMessage *messageDomain.Message `bson:"-" json:"last_message"`
}

type UserDetail struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username string              `bson:"username" json:"username"`
}