package main

import (
	"fmt"
	"log"
	"os"

	"portfolio-user-service/config"
	"portfolio-user-service/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

	// Configure CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://your-domain.com"}, // Add your frontend URL here
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

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
