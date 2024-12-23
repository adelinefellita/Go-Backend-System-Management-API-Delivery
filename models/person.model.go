package models

// Define Person struct here
type Person struct {
	Id           string  `json:"id" bson:"_id"`
	FirstName    string  `json:"firstName" bson:"firstName"`
	LastName     string  `json:"lastName" bson:"lastName"`
	PhoneNumber  string  `json:"phoneNumber" bson:"phoneNumber"`
	Address      Address `json:"address" bson:"address"` // Tidak perlu impor models.Address
	EmailAddress string  `json:"emailAddress" bson:"emailAddress"`
}
