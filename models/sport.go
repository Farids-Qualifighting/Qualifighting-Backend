package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Sport struct {
	ID   primitive.ObjectID `json:"id" bson:"id" binding:"required"`
	Name string             `json:"name" bson:"name" binding:"required"`
}
