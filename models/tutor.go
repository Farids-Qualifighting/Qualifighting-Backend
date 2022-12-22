package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tutor struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty"`
	FirstName string               `json:"first_name" bson:"first_name" binding:"required"`
	LastName  string               `json:"last_name" bson:"last_name" binding:"required"`
	Email     string               `json:"email" bson:"email" binding:"required"`
	Phone     string               `json:"phone" bson:"phone" binding:"required"`
	Subjects  []primitive.ObjectID `json:"subjects" bson:"subjects" binding:"required"`
	Students  []primitive.ObjectID `json:"students" bson:"students" binding:"required"`
}

type UpdateTutor struct {
	FirstName string               `json:"first_name" bson:"first_name,omitempty"`
	LastName  string               `json:"last_name" bson:"last_name,omitempty"`
	Email     string               `json:"email" bson:"email,omitempty"`
	Phone     string               `json:"phone" bson:"phone,omitempty"`
	Subjects  []primitive.ObjectID `json:"subjects" bson:"subjects,omitempty"`
	Students  []primitive.ObjectID `json:"students" bson:"students,omitempty"`
}
