package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Parent struct {
	ID        primitive.ObjectID   `json:"id" bson:"id" binding:"required"`
	FirstName string               `json:"first_name" bson:"first_name" binding:"required"`
	LastName  string               `json:"last_name" bson:"last_name" binding:"required"`
	Children  []primitive.ObjectID `json:"children" bson:"children" binding:"required"`
	Address   Address              `json:"address" bson:"address" binding:"required"`
	Phone     string               `json:"phone" bson:"phone" binding:"required"`
	Email     string               `json:"email" bson:"email"`
}
