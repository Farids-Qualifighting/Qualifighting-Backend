package models

type Address struct {
	Street  string `json:"street" bson:"street" binding:"required"`
	Zip     uint8  `json:"zip" bson:"zip" binding:"required"`
	HouseNo uint8  `json:"house_no" bson:"house_no" binding:"required"`
	City    string `json:"city" bson:"city" binding:"required"`
}

type UpdateAddress struct {
	Street  string `json:"street" bson:"street,omitempty"`
	Zip     uint8  `json:"zip" bson:"zip,omitempty"`
	HouseNo uint8  `json:"house_no" bson:"house_no,omitempty"`
	City    string `json:"city" bson:"city,omitempty"`
}
