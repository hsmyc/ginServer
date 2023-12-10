package models

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Effect_Type struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" validate:"required"`
	IsAttack  bool               `json:"isAttack,omitempty" validate:"required"`
	IsDefense bool               `json:"isDefense,omitempty" validate:"required"`
	IsSupport bool               `json:"isSupport,omitempty" validate:"required"`
}

func (e *Effect_Type) CType(effectType string) error {
	if effectType != "attack" && effectType != "defense" && effectType != "support" {
		return errors.New("effect type must be attack, defense, or support")
	}
	return nil
}

type Effect_Detail struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Effect_Type  Effect_Type        `bson:"effect_type,omitempty"`
	Effect_Value int                `json:"effect_value,omitempty" validate:"required"`
	Effect_Tile  int                `json:"effect_tile,omitempty" validate:"required"`
}
