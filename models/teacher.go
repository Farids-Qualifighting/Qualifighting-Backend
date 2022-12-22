package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Teacher struct {
	ID        primitive.ObjectID   `bson:"_id"`
	FirstName string               `json:"first_name" bson:"first_name" binding:"required"`
	LastName  string               `json:"last_name" bson:"last_name" binding:"required"`
	SchoolID  []primitive.ObjectID `json:"won_competition" bson:"won_competition" binding:"required"`
	Email     string               `json:"email" bson:"email" binding:"required"`
	Phone     string               `json:"string" bson:"string" binding:"required"`
}

type UpdateTeacher struct {
	FirstName string               `json:"first_name" bson:"first_name"`
	LastName  string               `json:"last_name" bson:"last_name"`
	SchoolID  []primitive.ObjectID `json:"won_competition" bson:"won_competition"`
	Email     string               `json:"email" bson:"email"`
	Phone     string               `json:"string" bson:"string"`
}
