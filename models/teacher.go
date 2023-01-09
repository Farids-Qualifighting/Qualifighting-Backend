package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Teacher struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty"`
	FirstName string               `json:"first_name" bson:"first_name" binding:"required" encryption:"true"`
	LastName  string               `json:"last_name" bson:"last_name" binding:"required" encryption:"true"`
	SchoolID  []primitive.ObjectID `json:"school" bson:"school" binding:"required"`
	Email     string               `json:"email" bson:"email" binding:"required" encryption:"true"`
	Phone     string               `json:"phone" bson:"phone" binding:"required" encryption:"true"`
}

type UpdateTeacher struct {
	FirstName string               `json:"first_name" bson:"first_name,omitempty" encryption:"true"`
	LastName  string               `json:"last_name" bson:"last_name,omitempty" encryption:"true"`
	SchoolID  []primitive.ObjectID `json:"school" bson:"school,omitempty"`
	Email     string               `json:"email" bson:"email,omitempty" encryption:"true"`
	Phone     string               `json:"phone" bson:"phone,omitempty" encryption:"true"`
}
