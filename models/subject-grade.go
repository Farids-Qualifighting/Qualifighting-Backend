package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type SubjectGrade struct {
	ID                primitive.ObjectID   `bson:"_id"`
	Grade             uint8                `json:"grade" bson:"grade" binding:"required"`
	IntermediateGrade bool                 `json:"intermediate_grade" bson:"intermediate_id" binding:"required"`
	StudentID         []primitive.ObjectID `json:"student_id" bson:"student_id" binding:"required"`
}

type UpdateSubjectGrade struct {
	Grade             uint8                `json:"grade" bson:"grade"`
	IntermediateGrade bool                 `json:"intermediate_grade" bson:"intermediate_id"`
	StudentID         []primitive.ObjectID `json:"student_id" bson:"student_id"`
}
