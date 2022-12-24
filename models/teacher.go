package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Teacher struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty"`
	FirstName string               `json:"first_name" bson:"first_name" binding:"required"`
	LastName  string               `json:"last_name" bson:"last_name" binding:"required"`
	SchoolID  []primitive.ObjectID `json:"school" bson:"school" binding:"required"`
	Email     string               `json:"email" bson:"email" binding:"required"`
	Phone     string               `json:"phone" bson:"phone" binding:"required"`
}

type UpdateTeacher struct {
	FirstName string               `json:"first_name" bson:"first_name,omitempty"`
	LastName  string               `json:"last_name" bson:"last_name,omitempty"`
	SchoolID  []primitive.ObjectID `json:"school" bson:"school,omitempty"`
	Email     string               `json:"email" bson:"email,omitempty"`
	Phone     string               `json:"phone" bson:"phone,omitempty"`
}
