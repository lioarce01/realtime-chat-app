package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
    ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Username string             `bson:"username,omitempty" json:"username"`
    Email    string             `bson:"email,omitempty" json:"email"`
    Password string             `bson:"password,omitempty" json:"password"`
}
