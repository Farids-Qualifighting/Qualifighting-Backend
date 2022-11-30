package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tutor struct {
	ID        primitive.ObjectID   `json:"id" bson:"id" binding:"required"`
	FirstName string               `json:"first_name" bson:"first_name" binding:"required"`
	LastName  string               `json:"last_name" bson:"last_name" binding:"required"`
	Email     string               `json:"email" bson:"email" binding:"required"`
	Phone     string               `json:"phone" bson:"phone" binding:"required"`
	Subjects  []primitive.ObjectID `json:"subjects" bson:"subjects" binding:"required"`
	Students  []primitive.ObjectID `json:"students" bson:"students" binding:"required"`
}
