package routes

import (
	_ "portfolio-user-service/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files" // Import swagger files
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitializeRoutes registers all routes in the router
func InitializeRoutes(router *gin.Engine) {
	// Register all routes here
	AuthRoutes(router)
	AddressRoutes(router)
	ContentRoutes(router)

	// Register Swagger
	setupSwagger(router)
}

// setupSwagger adds Swagger documentation support
func setupSwagger(router *gin.Engine) {
	// Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
