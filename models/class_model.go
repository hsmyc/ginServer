package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Class struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `json:"name,omitempty" validate:"required"`
	Base_Attack  int                `json:"attack,omitempty" validate:"required"`
	Base_Defense int                `json:"defense,omitempty" validate:"required"`
	Base_Health  int                `json:"health,omitempty" validate:"required"`
}
