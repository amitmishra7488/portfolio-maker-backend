package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"portfolio-user-service/config"
	"portfolio-user-service/routes"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to database
	config.ConnectDatabase()

	// Initialize Gin router
	router := gin.Default()

	// Register all routes from routes.go
	routes.InitializeRoutes(router)

	
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "server is running"})
	})
	// Start the server
	port := os.Getenv("PORT")
	fmt.Println("✅ Server is running on port " + port)
	router.Run(":" + port)
}
