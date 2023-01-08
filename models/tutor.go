package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tutor struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty"`
	FirstName string               `json:"first_name" bson:"first_name" binding:"required" encryption:"true"`
	LastName  string               `json:"last_name" bson:"last_name" binding:"required" encryption:"true"`
	Email     string               `json:"email" bson:"email" binding:"required" encryption:"true"`
	Phone     string               `json:"phone" bson:"phone" binding:"required" encryption:"true"`
	Subjects  []primitive.ObjectID `json:"subjects" bson:"subjects" binding:"required"`
	Students  []primitive.ObjectID `json:"students" bson:"students" binding:"required"`
}

type UpdateTutor struct {
	FirstName string               `json:"first_name" bson:"first_name,omitempty" encryption:"true"`
	LastName  string               `json:"last_name" bson:"last_name,omitempty" encryption:"true"`
	Email     string               `json:"email" bson:"email,omitempty" encryption:"true"`
	Phone     string               `json:"phone" bson:"phone,omitempty" encryption:"true"`
	Subjects  []primitive.ObjectID `json:"subjects" bson:"subjects,omitempty"`
	Students  []primitive.ObjectID `json:"students" bson:"students,omitempty"`
}
