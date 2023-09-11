package models

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `json:"name,omitempty" validate:"required"`
	Type  string             `json:"type,omitempty" validate:"required"`
	Price string             `json:"price,omitempty" validate:"required"`
	Level int                `json:"level,omitempty" validate:"required"`
}

// create a method to check the type of item

func (i *Item) CType(itemType string) error {
	if itemType != "weapon" && itemType != "armor" && itemType != "consumable" {
		return errors.New("item type must be weapon, armor, or consumable")
	}
	return nil
}
