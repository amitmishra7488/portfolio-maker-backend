package routes

import (
	"portfolio-user-service/controller"
	"portfolio-user-service/middleware"

	"github.com/gin-gonic/gin"
)

var (
	authController = new(controller.AuthController)
)

// AuthRoutes defines authentication-related routes
func AuthRoutes(router *gin.Engine) {
	authGroup := router.Group("/auth")

	// 🔓 Public routes
	authGroup.POST("/register", authController.RegisterUser)
	authGroup.GET("/verify-email/", authController.VerifyEmail)
	authGroup.GET("/verify-otp/", authController.VerifyRegistrationOTP)
	authGroup.POST("/login", authController.LoginUser)

	// 🔐 Protected routes
	protected := authGroup.Group("/")
	protected.Use(middleware.JWTAuthMiddleware()) // apply JWT middleware here
	protected.PATCH("/user-details", authController.UpdateUserDetails)
}
