package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Exam struct {
	ID       primitive.ObjectID   `bson:"_id,omitempty"`
	Topics   []string             `json:"topics" bson:"topics" binding:"required"`
	Date     time.Time            `json:"date" bson:"date" binding:"required"`
	Students []primitive.ObjectID `json:"students" bson:"students" binding:"required"`
}

type UpdateExam struct {
	Topics   []string             `json:"topics" bson:"topics,omitempty"`
	Date     time.Time            `json:"date" bson:"date,omitempty"`
	Students []primitive.ObjectID `json:"students" bson:"students,omitempty"`
}
