package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type School struct {
	ID       primitive.ObjectID   `bson:"_id"`
	Name     string               `json:"name" bson:"name" binding:"required"`
	Address  Address              `json:"address" bson:"address" binding:"required"`
	Email    string               `json:"email" bson:"email" binding:"required"`
	Phone    string               `json:"phone" bson:"phone" binding:"required"`
	Teachers []primitive.ObjectID `json:"teachers" bson:"teachers" binding:"required"`
}

type UpdateSchool struct {
	Name     string               `json:"name" bson:"name" binding:"required"`
	Address  Address              `json:"address" bson:"address" binding:"required"`
	Email    string               `json:"email" bson:"email" binding:"required"`
	Phone    string               `json:"phone" bson:"phone" binding:"required"`
	Teachers []primitive.ObjectID `json:"teachers" bson:"teachers" binding:"required"`
}
