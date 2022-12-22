package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Subject struct {
	ID         primitive.ObjectID   `bson:"_id,omitempty"`
	Name       string               `json:"name" bson:"name" binding:"required"`
	Teacher    primitive.ObjectID   `json:"teacher" bson:"teacher" binding:"required"`
	SchoolID   primitive.ObjectID   `json:"school_id" bson:"school_id" binding:"required"`
	StudentIDs []primitive.ObjectID `json:"student_ids" bson:"student_ids" binding:"required"`
}

type UpdateSubject struct {
	Name       string               `json:"name" bson:"name"`
	Teacher    primitive.ObjectID   `json:"teacher" bson:"teacher"`
	SchoolID   primitive.ObjectID   `json:"school_id" bson:"school_id"`
	StudentIDs []primitive.ObjectID `json:"student_ids" bson:"student_ids"`
}
