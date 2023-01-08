package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Parent struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty"`
	FirstName string               `json:"first_name" bson:"first_name" binding:"required" encryption:"true"`
	LastName  string               `json:"last_name" bson:"last_name" binding:"required" encryption:"true"`
	Children  []primitive.ObjectID `json:"children" bson:"children" binding:"required"`
	Address   Address              `json:"address" bson:"address,inline" binding:"required"`
	Phone     string               `json:"phone" bson:"phone" binding:"required" encryption:"true"`
	Email     string               `json:"email" bson:"email" encryption:"true"`
}

type UpdateParent struct {
	FirstName string               `json:"first_name" bson:"first_name,omitempty" encryption:"true"`
	LastName  string               `json:"last_name" bson:"last_name,omitempty" encryption:"true"`
	Children  []primitive.ObjectID `json:"children" bson:"children,omitempty"`
	Address   UpdateAddress        `json:"address" bson:"address,inline,omitempty"`
	Phone     string               `json:"phone" bson:"phone,omitempty" encryption:"true"`
	Email     string               `json:"email" bson:"email,omitempty" encryption:"true"`
}
