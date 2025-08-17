package routes

import (
	"portfolio-user-service/controller"
	"portfolio-user-service/middleware"
	"portfolio-user-service/repository/auth"
	authService "portfolio-user-service/services/auth"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// AuthRoutes defines authentication-related routes
func AuthRoutes(router *gin.Engine, db *gorm.DB, log *zap.Logger) {
	// Create repository
	authRepo := auth.NewAuthRepository(db)

	// Create service
	authService := authService.NewAuthService(authRepo, db, log)

	// Create controller with service
	authController := controller.NewAuthController(authService, log)

	authGroup := router.Group("/auth")

	// 🔓 Public routes
	authGroup.POST("/register", authController.RegisterUser)
	authGroup.GET("/verify-email/", authController.VerifyEmail)
	authGroup.GET("/verify-otp/", authController.VerifyRegistrationOTP)
	authGroup.POST("/login", authController.LoginUser)
	authGroup.GET("/all", authController.GetAllUser)

	// 🔐 Protected routes
	protected := authGroup.Group("/")
	protected.Use(middleware.JWTAuthMiddleware())
	protected.PATCH("/user-details", authController.UpdateUserDetails)
}
