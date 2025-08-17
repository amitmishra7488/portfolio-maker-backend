package controller

import (
	"fmt"
	"net/http"
	"portfolio-user-service/repository/auth/models"
	"portfolio-user-service/services/auth"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// AuthController struct to hold service instance
type AuthController struct {
	AuthService auth.Service  
	Log         *zap.Logger
}

// Constructor
func NewAuthController(authService auth.Service, log *zap.Logger) *AuthController {
	return &AuthController{
		AuthService: authService,
		Log:         log,
	}
}


func (ac *AuthController) RegisterUser(c *gin.Context) {
	var registerRequest models.RegisterRequest

	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		ac.Log.Warn("Invalid registration input", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	err := ac.AuthService.Register(registerRequest)
	if err != nil {
		ac.Log.Warn("User registration failed", zap.Error(err))
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ac.Log.Info("User registered successfully", zap.String("email", registerRequest.Email))
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (ac *AuthController) LoginUser(c *gin.Context) {
	var loginRequest models.LoginRequest
	// Bind JSON request
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	fmt.Println("loginRequest: ", loginRequest)

	// Call service method
	token, err := ac.AuthService.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Return response
	c.JSON(http.StatusOK, gin.H{"token": token, "message": "Login successful"})
}

func (ac *AuthController) UpdateUserDetails(c *gin.Context) {
	var input models.UpdateUserDetailInput // uses *string, *map[...] for optional fields

	// 1. Bind incoming JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. Get userID from context
	userIDRaw, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userID, ok := userIDRaw.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID in context"})
		return
	}

	// 3. Call service layer
	updatedDetail, err := ac.AuthService.UpdateUserDetails(userID, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user details"})
		return
	}

	// 4. Return updated data
	c.JSON(http.StatusOK, gin.H{
		"message": "User details updated successfully",
		"data":    updatedDetail,
	})
}

func (ac *AuthController) VerifyEmail(c *gin.Context) {
	email := c.Query("email")
	name := c.Query("name")
	if email == "" || name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email parameter is required"})
		return
	}

	// Pass request to AuthService for email verification
	err := ac.AuthService.VerifyEmail(email, name)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Otp Sent To Your Email Successfully"})
}

func (ac *AuthController) VerifyRegistrationOTP(c *gin.Context) {
	email := c.Query("email")
	otp := c.Query("otp")
	name := c.Query("name")
	if email == "" || otp == "" || name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and OTP parameters are required"})
		return
	}

	// Pass request to AuthService for OTP verification
	err := ac.AuthService.VerifyRegistrationOTP(name, email, otp)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Otp Verified Successfully"})
}

func (ac *AuthController) GetAllUser(c *gin.Context) {
	userDetails, err := ac.AuthService.GetAllUser()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"data": userDetails})
}
