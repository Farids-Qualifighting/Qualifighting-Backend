package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ParentData struct {
	ID        primitive.ObjectID   `json:"_id" bson:"_id"`
	FirstName string               `json:"first_name"`
	LastName  string               `json:"last_name"`
	Children  []primitive.ObjectID `json:"children" bson:"children"`
	Address   Address              `json:"address" bson:"address"`
	Phone     string               `json:"phone_numer" bson:"phone_numer"`
	Email     string               `json:"email" bson:"email"`
}
