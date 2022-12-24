package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type SubjectGrade struct {
	ID                primitive.ObjectID   `bson:"_id,omitempty"`
	Grade             uint8                `json:"grade" bson:"grade" binding:"required"`
	IntermediateGrade bool                 `json:"intermediate_grade" bson:"intermediate_grade"`
	StudentID         []primitive.ObjectID `json:"student_id" bson:"student_id" binding:"required"`
}

type UpdateSubjectGrade struct {
	Grade             uint8                `json:"grade" bson:"grade,omitempty"`
	IntermediateGrade bool                 `json:"intermediate_grade" bson:"intermediate_grade,omitempty"`
	StudentID         []primitive.ObjectID `json:"student_id" bson:"student_id,omitempty"`
}
