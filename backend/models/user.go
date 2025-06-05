package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
    ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Name         string             `bson:"name" json:"name"`
    Email        string             `bson:"email" json:"email"`
    PasswordHash string             `bson:"password_hash" json:"-"`
    CreatedAt    primitive.DateTime `bson:"created_at,omitempty" json:"created_at"`
}
