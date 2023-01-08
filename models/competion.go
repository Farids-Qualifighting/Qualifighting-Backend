package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Competition struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Name           string             `json:"name" bson:"name" binding:"required" encryption:"true"`
	Date           time.Time          `json:"date" bson:"date" binding:"required"`
	Rank           uint8              `json:"rank" bson:"rank" binding:"required"`
	WonCompetition bool               `json:"won_competition" bson:"won_competition" binding:"required"`
	StudentID      primitive.ObjectID `json:"student_id" bson:"student_id" binding:"required"`
}

type UpdateCompetition struct {
	Name           string             `json:"name" bson:"name,omitempty"`
	Date           time.Time          `json:"date" bson:"date,omitempty"`
	Rank           uint8              `json:"rank" bson:"rank,omitempty"`
	WonCompetition bool               `json:"won_competition" bson:"won_competition,omitempty"`
	StudentID      primitive.ObjectID `json:"student_id" bson:"student_id,omitempty"`
}
