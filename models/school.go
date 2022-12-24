package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type School struct {
	ID       primitive.ObjectID   `bson:"_id,omitempty"`
	Name     string               `json:"name" bson:"name" binding:"required"`
	Address  Address              `json:"address" bson:"address,inline" binding:"required"`
	Email    string               `json:"email" bson:"email" binding:"required"`
	Phone    string               `json:"phone" bson:"phone" binding:"required"`
	Teachers []primitive.ObjectID `json:"teachers" bson:"teachers" binding:"required"`
}

type UpdateSchool struct {
	Name     string               `json:"name" bson:"name,omitempty"`
	Address  UpdateAddress        `json:"address" bson:"address,inline,omitempty"`
	Email    string               `json:"email" bson:"email,omitempty"`
	Phone    string               `json:"phone" bson:"phone,omitempty"`
	Teachers []primitive.ObjectID `json:"teachers" bson:"teachers,omitempty"`
}
