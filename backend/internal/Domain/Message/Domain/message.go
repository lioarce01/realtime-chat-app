package message

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ChatID    primitive.ObjectID `bson:"chat_id" json:"chat_id"`
	Content   string             `bson:"content" json:"content"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	Sender    UserDetail         `json:"sender"`   
	Receiver  UserDetail         `json:"receiver"` 
}

type UserDetail struct {
	ID          primitive.ObjectID `json:"id"`
	Username    string             `json:"username"`
	Profile_Pic  string             `json:"profile_pic"`
}
