package controller

import (
	"fmt"
	"net/http"
	"portfolio-user-service/repository/auth/models"
	"portfolio-user-service/services/auth"

	"github.com/gin-gonic/gin"
)

// AuthController struct to hold service instance
type AuthController struct{}

var (
	authService = new(auth.AuthService)
)

// RegisterUser handles user registration
func (ac AuthController) RegisterUser(c *gin.Context) {
	var registerRequest models.RegisterRequest

	// Bind JSON request to struct
	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	// Pass request to AuthService for user creation
	err := authService.Register(registerRequest)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// LoginUser handles user login request
func (ac AuthController) LoginUser(c *gin.Context) {
	var loginRequest struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// Bind JSON request
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	fmt.Println("loginRequest: ", loginRequest)

	// Call service method
	token, err := authService.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Return response
	c.JSON(http.StatusOK, gin.H{"token": token, "message": "Login successful"})
}

func (ac AuthController) UpdateUserDetails(c *gin.Context) {
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
	updatedDetail, err := authService.UpdateUserDetails(userID, input)
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


