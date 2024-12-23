package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Define Address struct here
type Address struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	AddressLine1 string             `json:"addressLine1" bson:"addressLine1"`
	AddressLine2 string             `json:"addressLine2" bson:"addressLine2"`
	City         string             `json:"city" bson:"city"`
	State        string             `json:"state" bson:"state"`
	ZipCode      string             `json:"zipCode" bson:"zipCode"`
	Country      string             `json:"country" bson:"country"`
}
