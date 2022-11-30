package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Trainer struct {
	ID        primitive.ObjectID   `json:"id" bson:"id" binding:"required"`
	FirstName string               `json:"first_name" bson:"first_name" binding:"required"`
	LastName  string               `json:"last_name" bson:"last_name" binding:"required"`
	StudentID []primitive.ObjectID `json:"student_ids" bson:"student_id" binding:"required"`
	Sport     []primitive.ObjectID `json:"sport" bson:"sport" binding:"required"`
}
