package controllers

import (
	"net/http"

	"basic-server/database"
	"basic-server/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreatePerson handles creating a new person
func CreatePerson(c *gin.Context) {
	var person models.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Jika ID kosong, buat ID baru
	if person.Id == "" {
		person.Id = primitive.NewObjectID().Hex() // Generate new ObjectId
	}

	// Save the person record to the database
	SavePersonToDB(c, person)
}

// Retrieve Person: /person/<person-id>
func GetPerson(c *gin.Context) {
	id := c.Param("id")
	db := database.New()
	collection := db.Client.Database("admin").Collection("PersonForDemo")

	var result models.Person
	filter := bson.D{{"_id", id}}
	err := collection.FindOne(c, filter).Decode(&result)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}
	c.JSON(http.StatusOK, result)
}

// Delete Person: /person/<person-id>
func DeletePerson(c *gin.Context) {
	id := c.Param("id")
	db := database.New()
	collection := db.Client.Database("admin").Collection("PersonForDemo")

	filter := bson.D{{"_id", id}}

	_, err := collection.DeleteOne(c, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete person"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Person successfully deleted"})
}

// Update Person: /person/<person-id>
func UpdatePerson(c *gin.Context) {
	id := c.Param("id")
	db := database.New()
	collection := db.Client.Database("admin").Collection("PersonForDemo")

	// Get updated person data from the request
	var person models.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Prepare the filter for finding the person
	filter := bson.M{"_id": id}

	// Prepare update data
	update := bson.M{
		"$set": bson.M{
			"firstName":    person.FirstName,
			"lastName":     person.LastName,
			"phoneNumber":  person.PhoneNumber,
			"address":      person.Address,
			"emailAddress": person.EmailAddress,
		},
	}

	// Perform the update
	_, err := collection.UpdateOne(c, filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update person"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Person updated successfully"})
}

// Helper function to save person to the database
func SavePersonToDB(c *gin.Context, person models.Person) {
	db := database.New()
	collection := db.Client.Database("admin").Collection("PersonForDemo")

	// Insert the person into the database
	result, err := collection.InsertOne(c, person)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create person"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": result.InsertedID})
}
