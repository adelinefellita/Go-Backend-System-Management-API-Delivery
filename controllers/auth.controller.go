package controllers

import (
	"net/http"
	"time"

	"basic-server/database"
	"basic-server/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("your_secret_key")

// Register user
func Register(ct *gin.Context) {
	var user models.User
	if err := ct.ShouldBindJSON(&user); err != nil {
		ct.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	// Hash password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(hashedPassword)

	// Set ID ke nilai kosong agar MongoDB menghasilkan ID secara otomatis
	user.ID = primitive.NilObjectID

	db := database.New()
	collection := db.Client.Database("admin").Collection("Users")

	_, err := collection.InsertOne(ct, user)
	if err != nil {
		ct.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to register user", "error": err.Error()})
		return
	}

	ct.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// Login user
func Login(c *gin.Context) {
	var credentials models.Credentials

	// Bind JSON to credentials struct
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	// Verify user credentials (this is just an example, implement your own logic)
	user, err := models.VerifyUserCredentials(credentials.Username, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	// Generate JWT token
	token := generateToken(user)

	// Return token and user role in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"role":    user.Role,
		"token":   token,
	})
}

// Function to generate JWT token
func generateToken(user models.User) string {
	claims := jwt.MapClaims{
		"sub":  user.ID,                               // User ID
		"role": user.Role,                             // User role (manager/courier)
		"exp":  time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret key
	signedToken, _ := token.SignedString(secretKey)

	return signedToken
}
