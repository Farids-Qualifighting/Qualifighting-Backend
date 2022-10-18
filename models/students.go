package models

import (
	"time"
)

type CreateStudentData struct {
	Email              string    `json:"email"`
	Password           string    `json:"password"`
	FirstName          string    `json:"first_name"`
	LastName           string    `json:"last_name"`
	SchoolID           string    `json:"school_id"`
	ClassTeacher       string    `json:"class_teacher"`
	BirthDate          time.Time `json:"birth_date"`
	Gender             string    `json:"gender"`
	Address            Address   `json:"address"`
	Phone              string    `json:"phone_numer"`
	SocialMediaContact string    `json:"social_media_contact"`
}

type StudentData struct {
	ID                 string    `bson:"_id"`
	FirstName          string    `json:"first_name" bson:"first_name"`
	LastName           string    `json:"last_name" bson:"last_name"`
	ClassTeacher       string    `bson:"class_teacher,omitempty"`
	BirthDate          time.Time `json:"birth_date" bson:"birth_date"`
	Gender             string    `json:"gender" bson:"gender"`
	Address            Address   `json:"address" bson:"address"`
	Phone              string    `json:"phone_numer" bson:"phone_numer"`
	Email              string    `json:"email" bson:"email"`
	SocialMediaContact string    `json:"social_media_contact" bson:"social_media_contact"`
	Certificates       []string  `json:"certificates" bson:"certificates,omitempty"`
	Level              uint8     `json:"level" bson:"level"`
	SchoolID           string    `json:"school_id" bson:"school_id"`
}
