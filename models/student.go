package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	ID           primitive.ObjectID `bson:"_id"`
	FirstName    string             `json:"first_name" bson:"first_name" binding:"required"`
	LastName     string             `json:"last_name" bson:"last_name" binding:"required"`
	ClassTeacher string             `json:"class_teacher" bson:"class_teacher" binding:"required"`
	Birthday     time.Time          `json:"birthday" bson:"birthday" binding:"required"`
	Gender       string             `json:"gender" bson:"gender" binding:"required"`
	Address      Address            `json:"address" bson:"address" binding:"required"`
	Phone        string             `json:"phone" bson:"phone" binding:"required"`
	Email        string             `json:"email" bson:"email"`
	SocialMedia  string             `json:"social_media" bson:"social_media"`
	Certificate  []string           `json:"certificate" bson:"certificate"`
}

type UpdateStudent struct {
	FirstName    string        `json:"first_name" bson:"first_name,omitempty"`
	LastName     string        `json:"last_name" bson:"last_name,omitempty"`
	ClassTeacher string        `json:"class_teacher" bson:"class_teacher,omitempty"`
	Birthday     time.Time     `json:"birthday" bson:"birthday,omitempty"`
	Gender       string        `json:"gender" bson:"gender,omitempty"`
	Address      UpdateAddress `json:"address" bson:"address,omitempty"`
	Phone        string        `json:"phone" bson:"phone,omitempty"`
	Email        string        `json:"email" bson:"email,omitempty"`
	SocialMedia  string        `json:"social_media" bson:"social_media,omitempty"`
	Certificate  []string      `json:"certificate" bson:"certificate,omitempty"`
}
