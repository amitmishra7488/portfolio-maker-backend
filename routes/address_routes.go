package routes

import (
	"portfolio-user-service/middleware"
	"portfolio-user-service/controller"

	"github.com/gin-gonic/gin"
)

var (
	addressController= new(controller.AddressController)
)

func AddressRoutes(router *gin.Engine) {
	addressGroup := router.Group("/address")
	addressGroup.Use(middleware.JWTAuthMiddleware())

	addressGroup.POST("/", addressController.CreateAddress)
	addressGroup.PATCH("/:addressID", addressController.UpdateExistingAddress)
	addressGroup.GET("/", addressController.GetAllAddresses)
}