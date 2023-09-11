package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Character struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Name    string             `json:"name,omitempty" validate:"required"`
	Level   int                `json:"level,omitempty" validate:"required"`
	ClassID primitive.ObjectID `bson:"class_id,omitempty"`
	ItemID  primitive.ObjectID `bson:"item_id,omitempty"`
}
