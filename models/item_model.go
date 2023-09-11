package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ItemType string

const (
	Weapon     ItemType = "Weapon"
	Armor      ItemType = "Armor"
	Consumable ItemType = "Consumable"
)

type Item struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `json:"name,omitempty" validate:"required"`
	Type  ItemType           `json:"type,omitempty" validate:"required"`
	Price string             `json:"price,omitempty" validate:"required"`
	Level int                `json:"level,omitempty" validate:"required"`
}

