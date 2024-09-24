package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name" binding:"required"`
	Password string             `json:"password" bson:"password" binding:"required"`
	Role     string             `json:"role" bson:"role"`
}

type SignInInput struct {
	Name     string `json:"name" bson:"name" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
	Device   string `json:"device" bson:"device"`
}
