package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type TrainerData struct {
	ID        primitive.ObjectID   `json:"_id" bson:"_id"`
	FirstName string               `json:"first_name" bson:"first_name"`
	LastName  string               `json:"last_name" bson:"last_name"`
	Sports    []primitive.ObjectID `json:"sports" bson:"sports"`
	Students  []primitive.ObjectID `json:"students" bson:"students"`
}
