package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DailyNote struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	StudentID primitive.ObjectID `json:"student_id" bson:"student_id" binding:"required"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at" binding:"required"`
	Subject   primitive.ObjectID `json:"subject" bson:"subject" binding:"required"`
	Rating    uint8              `json:"rating" bson:"rating" binding:"required"`
	Note      string             `json:"note" bson:"note" binding:"required"`
	CreatorID primitive.ObjectID `json:"creator_id" bson:"creator_id" binding:"required"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at" binding:"required"`
}

type UpdateDailyNote struct {
	StudentID primitive.ObjectID `json:"student_id" bson:"student_id,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at,omitempty"`
	Subject   primitive.ObjectID `json:"subject" bson:"subject,omitempty"`
	Rating    uint8              `json:"rating" bson:"rating,omitempty"`
	Note      string             `json:"note" bson:"note,omitempty"`
	CreatorID primitive.ObjectID `json:"creator_id" bson:"creator_id,omitempty"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
}
