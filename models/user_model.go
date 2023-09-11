package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Name   string             `json:"name,omitempty" validate:"required"`
	Email  string             `json:"email,omitempty" validate:"required,email"`
	Image  string             `json:"image,omitempty"`
	CharID primitive.ObjectID `bson:"char_id,omitempty"`
}
