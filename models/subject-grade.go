package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type SubjectGrade struct {
	Subject           primitive.ObjectID   `json:"subject_id" bson:"subject_id" binding:"required"`
	Grade             uint8                `json:"grade" bson:"grade" binding:"required"`
	IntermediateGrade bool                 `json:"intermediate_grade" bson:"intermediate_id" binding:"required"`
	StudentID         []primitive.ObjectID `json:"student_id" bson:"student_id" binding:"required"`
}
