package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Sport struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name" binding:"required"`
}

type UpdateSport struct {
	Name string `json:"name" bson:"name"`
}
