package models

type Address struct {
	Street string `json:"street" bson:"street"`
	Zip    uint8  `json:"zip" bson:"zip"`
	City   string `json:"city" bson:"city"`
}
