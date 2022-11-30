package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Exam struct {
	ID       primitive.ObjectID   `json:"id" bson:"id" binding:"required"`
	Topics   []string             `json:"topics" bson:"topics" binding:"required"`
	Date     time.Time            `json:"date" bson:"date" binding:"required"`
	Students []primitive.ObjectID `json:"students" bson:"students" binding:"required"`
}
