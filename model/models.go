package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Todo       string             `json:"todo,omitempty"`
	IsResolved bool               `json:"isResolved,omitempty"`
}
