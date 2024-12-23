package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Delivery struct {
	ID              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Receiver        string             `json:"receiver" bson:"receiver"`
	Address         Address            `json:"address" bson:"address"`
	Status          string             `json:"status" bson:"status"`
	ProofOfDelivery string             `json:"proofOfDelivery" bson:"proofOfDelivery"`
	CourierId       string             `json:"courierId" bson:"courierId"`
}

func (d *Delivery) GetIDAsString() string {
	return d.ID.Hex() // Mengubah ID ke format string
}
