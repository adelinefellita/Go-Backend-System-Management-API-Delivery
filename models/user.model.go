package models

import (
	"basic-server/database"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// Fungsi untuk memverifikasi kredensial pengguna (username dan password)
func VerifyUserCredentials(username, password string) (User, error) {
	var user User
	db := database.New()
	collection := db.Client.Database("admin").Collection("Users")

	// Cari user berdasarkan username
	err := collection.FindOne(nil, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return user, errors.New("user not found")
	}

	// Verifikasi password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, errors.New("invalid credentials")
	}

	return user, nil
}

type Credentials struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"` // '_id' akan di-generate otomatis jika kosong
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"`
	Role     string             `json:"role" bson:"role"` // manager atau kurir
}
