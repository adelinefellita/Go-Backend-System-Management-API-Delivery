package controllers

import (
	"basic-server/database"
	"basic-server/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateDeliveryAddress handles the creation of delivery address
func CreateDeliveryAddress(c *gin.Context) {
	var delivery models.Delivery
	// Bind the incoming JSON to the Delivery struct
	if err := c.ShouldBindJSON(&delivery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Insert the address into the Addresses collection
	db := database.New()
	addressCollection := db.Client.Database("admin").Collection("Addresses")

	// Insert the new address into MongoDB and get the inserted ID
	addressResult, err := addressCollection.InsertOne(c, delivery.Address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create address", "details": err.Error()})
		return
	}

	// Assign the address ID to the delivery
	delivery.Address.ID = addressResult.InsertedID.(primitive.ObjectID)

	// Insert the new delivery into the Deliveries collection
	deliveryCollection := db.Client.Database("admin").Collection("Deliveries")
	_, err = deliveryCollection.InsertOne(c, delivery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create delivery"})
		return
	}

	// Return success message with the delivery ID
	c.JSON(http.StatusCreated, gin.H{"message": "Address created successfully", "id": delivery.ID.Hex()})
}

// Define other functions for Update, Delete, etc.
func UpdateDeliveryAddress(c *gin.Context) {
	// Ambil ID dari URL parameter
	id := c.Param("id")

	// Validasi ID format ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Bind input JSON ke struct Delivery
	var updatedDelivery models.Delivery
	if err := c.ShouldBindJSON(&updatedDelivery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Akses database
	db := database.New()
	deliveryCollection := db.Client.Database("admin").Collection("Deliveries")

	// Update delivery berdasarkan ID
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": updatedDelivery}

	result, err := deliveryCollection.UpdateOne(c, filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update delivery"})
		return
	}

	// Cek apakah ada dokumen yang diupdate
	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Delivery not found"})
		return
	}

	// Return respons sukses
	c.JSON(http.StatusOK, gin.H{"message": "Delivery updated successfully", "id": id})
}

// DeleteDeliveryAddress handles deleting delivery address
func DeleteDeliveryAddress(c *gin.Context) {
	// Ambil ID dari URL parameter
	id := c.Param("id")

	// Convert ID ke tipe ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Mengakses database
	db := database.New()
	addressCollection := db.Client.Database("admin").Collection("Addresses")
	deliveryCollection := db.Client.Database("admin").Collection("Deliveries")

	// Hapus delivery terkait yang memiliki address ID tersebut
	// Pastikan penghapusan delivery berdasarkan address._id, bukan hanya _id saja
	deleteResult, err := deliveryCollection.DeleteMany(c, bson.M{"address._id": objectID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete related deliveries"})
		return
	}

	// Jika tidak ada data yang terhapus pada collection deliveries
	if deleteResult.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No related deliveries found to delete"})
		return
	}

	// Hapus alamat dari koleksi Addresses
	_, err = addressCollection.DeleteOne(c, bson.M{"_id": objectID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete address"})
		return
	}

	// Jika berhasil dihapus, kembalikan respons sukses
	c.JSON(http.StatusOK, gin.H{"message": "Address and related delivery deleted successfully", "id": id})
}

// GetDeliveryAddresses handles retrieving all delivery addresses
func GetDeliveryAddresses(c *gin.Context) {
	var deliveries []models.Delivery

	db := database.New()
	collection := db.Client.Database("admin").Collection("Deliveries")

	cursor, err := collection.Find(c, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve deliveries"})
		return
	}
	defer cursor.Close(c)

	// Iterasi melalui cursor untuk mendekodekan setiap delivery
	for cursor.Next(c) {
		var delivery models.Delivery

		// Cek apakah ada error saat mendekode
		if err := cursor.Decode(&delivery); err != nil {
			// Debugging log untuk mencetak kesalahan yang terjadi
			log.Printf("Error decoding delivery: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode delivery", "details": err.Error()})
			return
		}

		// Periksa apakah ObjectID valid atau kosong
		if delivery.ID.IsZero() {
			// Jika ObjectID tidak valid, beri log dan lanjutkan ke delivery berikutnya
			log.Printf("Invalid ObjectID found for delivery: %+v\n", delivery)
			continue // Skip delivery dengan ID kosong
		}

		// Simpan delivery asli ke slice deliveries
		deliveries = append(deliveries, delivery)
	}

	// Cek error saat iterasi cursor
	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to iterate over deliveries"})
		return
	}

	// Membentuk response dari data deliveries
	var response []map[string]interface{}
	for _, delivery := range deliveries {
		// Pastikan ID dikonversi ke string
		deliveryResponse := map[string]interface{}{
			"id":              delivery.ID.Hex(),
			"receiver":        delivery.Receiver,
			"address":         delivery.Address,
			"status":          delivery.Status,
			"proofOfDelivery": delivery.ProofOfDelivery,
			"courierId":       delivery.CourierId,
		}
		response = append(response, deliveryResponse)
	}

	// Kirim response dengan data deliveries
	c.JSON(http.StatusOK, gin.H{"deliveries": response})
}

// UpdateDeliveryStatus handles updating delivery status
func UpdateDeliveryStatus(c *gin.Context) {
	id := c.Param("id")
	var statusUpdate struct {
		Status string `json:"status"`
	}

	// Bind JSON input ke struct statusUpdate
	if err := c.ShouldBindJSON(&statusUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Konversi ID dari string ke ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Mengakses database
	db := database.New()
	deliveryCollection := db.Client.Database("admin").Collection("Deliveries")

	// Update status di MongoDB
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"status": statusUpdate.Status}}

	result, err := deliveryCollection.UpdateOne(c, filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update delivery status"})
		return
	}

	// Periksa apakah ada dokumen yang diperbarui
	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Delivery not found"})
		return
	}

	// Berikan respons sukses
	c.JSON(http.StatusOK, gin.H{
		"message": "Delivery status updated successfully",
		"id":      id,
		"status":  statusUpdate.Status,
	})
}
