package routes

import (
	"portfolio-user-service/controller"
	"portfolio-user-service/middleware"
	"portfolio-user-service/repository/address"
	addressService "portfolio-user-service/services/address"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func AddressRoutes(router *gin.Engine, db *gorm.DB, log *zap.Logger) {
	// Create repository
	addressRepo := address.NewAddressRepository(db)

	// Create service
	addressService := addressService.NewAddressService(addressRepo, db, log)

	// Create controller with service
	addressController := controller.NewAddressController(addressService, log)
	addressGroup := router.Group("/address")
	addressGroup.Use(middleware.JWTAuthMiddleware())

	addressGroup.POST("/", addressController.CreateAddress)
	addressGroup.PATCH("/:addressID", addressController.UpdateExistingAddress)
	addressGroup.GET("/", addressController.GetAllAddresses)
}
