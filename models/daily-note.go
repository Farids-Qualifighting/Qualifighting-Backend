package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DailyNote struct {
	ID        primitive.ObjectID `json:"id" bson:"id" binding:"required"`
	StudentID primitive.ObjectID `json:"student_id" bson:"student_id" binding:"required"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at" binding:"required"`
	Subject   primitive.ObjectID `json:"subject" bson:"subject" binding:"required"`
	Rating    uint8              `json:"rating" bson:"rating" binding:"required"`
	Note      string             `json:"note" bson:"note" binding:"required"`
	CreatorID primitive.ObjectID `json:"creator_id" bson:"creator_id" binding:"required"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at" binding:"required"`
}
