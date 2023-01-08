package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	FirstName    string             `json:"first_name" bson:"first_name" binding:"required" encryption:"true"`
	LastName     string             `json:"last_name" bson:"last_name" binding:"required" encryption:"true"`
	ClassTeacher string             `json:"class_teacher" bson:"class_teacher" encryption:"true"`
	Birthday     time.Time          `json:"birthday" bson:"birthday" binding:"required"`
	Gender       string             `json:"gender" bson:"gender" binding:"required" encryption:"true"`
	Address      Address            `json:"address" bson:"address,inline" binding:"required"`
	Phone        string             `json:"phone" bson:"phone" binding:"required" encryption:"true"`
	Email        string             `json:"email" bson:"email" encryption:"true"`
	SocialMedia  string             `json:"social_media" bson:"social_media" encryption:"true"`
	Certificate  []string           `json:"certificate" bson:"certificate"`
}

type UpdateStudent struct {
	FirstName    string        `json:"first_name" bson:"first_name,omitempty" encryption:"true"`
	LastName     string        `json:"last_name" bson:"last_name,omitempty" encryption:"true"`
	ClassTeacher string        `json:"class_teacher" bson:"class_teacher,omitempty"`
	Birthday     time.Time     `json:"birthday" bson:"birthday,omitempty"`
	Gender       string        `json:"gender" bson:"gender,omitempty" encryption:"true"`
	Address      UpdateAddress `json:"address" bson:"address,inline,omitempty"`
	Phone        string        `json:"phone" bson:"phone,omitempty" encryption:"true"`
	Email        string        `json:"email" bson:"email,omitempty" encryption:"true"`
	SocialMedia  string        `json:"social_media" bson:"social_media,omitempty" encryption:"true"`
	Certificate  []string      `json:"certificate" bson:"certificate,omitempty"`
}
