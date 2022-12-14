package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Trainer struct {
	ID         primitive.ObjectID   `bson:"_id,omitempty"`
	FirstName  string               `json:"first_name" bson:"first_name" binding:"required" encryption:"true"`
	LastName   string               `json:"last_name" bson:"last_name" binding:"required" encryption:"true"`
	StudentIDs []primitive.ObjectID `json:"student_ids" bson:"student_id" binding:"required"`
	Sport      []primitive.ObjectID `json:"sport" bson:"sport" binding:"required"`
}

type UpdateTrainer struct {
	FirstName string               `json:"first_name" bson:"first_name,omitempty" encryption:"true"`
	LastName  string               `json:"last_name" bson:"last_name,omitempty" encryption:"true"`
	StudentID []primitive.ObjectID `json:"student_ids" bson:"student_id,omitempty"`
	Sport     []primitive.ObjectID `json:"sport" bson:"sport,omitempty"`
}
