package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Parent struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty"`
	FirstName string               `json:"first_name" bson:"first_name" binding:"required"`
	LastName  string               `json:"last_name" bson:"last_name" binding:"required"`
	Children  []primitive.ObjectID `json:"children" bson:"children" binding:"required"`
	Address   Address              `json:"address" bson:"address,inline" binding:"required"`
	Phone     string               `json:"phone" bson:"phone" binding:"required"`
	Email     string               `json:"email" bson:"email"`
}

type UpdateParent struct {
	FirstName string               `json:"first_name" bson:"first_name,omitempty"`
	LastName  string               `json:"last_name" bson:"last_name,omitempty"`
	Children  []primitive.ObjectID `json:"children" bson:"children,omitempty"`
	Address   UpdateAddress        `json:"address" bson:"address,inline,omitempty"`
	Phone     string               `json:"phone" bson:"phone,omitempty"`
	Email     string               `json:"email" bson:"email,omitempty"`
}
