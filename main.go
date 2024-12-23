package main

import (
	"basic-server/controllers"
	"basic-server/middlewares"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Konfigurasi CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://example.com"}, // Ganti dengan domain front-end Anda
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Endpoint for Authentication
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/login", controllers.Login)
		authRoutes.POST("/register", controllers.Register) // Jika Anda ingin menyediakan endpoint untuk register
	}

	// Endpoint for Manager
	managerRoutes := router.Group("/manager")
	managerRoutes.Use(middlewares.RoleBasedAuth([]string{"manager"}))
	{
		// Rute untuk addressses
		managerRoutes.POST("/addresses", controllers.CreateDeliveryAddress)
		managerRoutes.PUT("/addresses/:id", controllers.UpdateDeliveryAddress)
		managerRoutes.DELETE("/addresses/:id", controllers.DeleteDeliveryAddress)
		managerRoutes.GET("/addresses", controllers.GetDeliveryAddresses)

		// Rute untuk deliveries
		managerRoutes.POST("/deliveries", controllers.CreateDeliveryAddress)
		managerRoutes.PUT("/deliveries/:id", controllers.UpdateDeliveryAddress)
		managerRoutes.DELETE("/deliveries/:id", controllers.DeleteDeliveryAddress)
		managerRoutes.GET("/deliveries", controllers.GetDeliveryAddresses)

		// Rute untuk person
		managerRoutes.GET("/persons", controllers.GetPerson)
	}

	// Endpoint for Courier
	courierRoutes := router.Group("/courier")
	courierRoutes.Use(middlewares.RoleBasedAuth([]string{"courier"}))
	{
		// Rute untuk addressses
		courierRoutes.GET("/addresses", controllers.GetDeliveryAddresses)

		// Rute untuk deliveries
		courierRoutes.GET("/deliveries", controllers.GetDeliveryAddresses)
		courierRoutes.PUT("/deliveries/:id/status", controllers.UpdateDeliveryStatus)
	}

	router.Use(cors.Default())

	router.Run(":8080")
}
