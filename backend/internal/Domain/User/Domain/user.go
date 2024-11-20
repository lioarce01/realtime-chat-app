package domain

import (
	chatDomain "backend/internal/Domain/Chat/Domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
    ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Username string             `bson:"username,omitempty" json:"username"`
    Email    string             `bson:"email,omitempty" json:"email"`
    Profile_Pic string `bson:"profile_pic,omitempty" json:"profile_pic"`
    Sub string `bson:"sub,omitempty" json:"sub"`
    Chats []chatDomain.Chat `bson:"chats,omitempty" json:"chats"`
}
