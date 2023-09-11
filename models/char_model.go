package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Character is the model for a character in the game

type Character struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `json:"name,omitempty" validate:"required"`
	Class string             `json:"class,omitempty" validate:"required"`
	Level int                `json:"level,omitempty" validate:"required"`
}
