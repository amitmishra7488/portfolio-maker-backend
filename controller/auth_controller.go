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
// @Summary Register a new user
// @Description Register a new user with email, password, and other details
// @Tags Auth
// @Accept json
// @Produce json
// @Param data body models.RegisterRequest true "User registration data"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Router /auth/register [post]
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
// @Summary User login
// @Description Login with email and password to get a JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param data body models.LoginRequest true "Login credentials"
// @Router /auth/login [post]
func (ac AuthController) LoginUser(c *gin.Context) {
	var loginRequest models.LoginRequest
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


// UpdateUserDetails updates user details
// @Summary Update user details
// @Description Update user details such as name, email, etc.
// @Tags User
// @Accept json
// @Produce json
// @Param Authorization header string true "Token"
// @Param data body models.UpdateUserDetailInput true "User details to update"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/user-details [patch]
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


// VerifyEmail handles the email verification process by extracting the email 
// @Summary Verify email
// @Description Verify email by sending an OTP
// @Tags Auth
// @Accept json
// @Produce json
// @Param email query string true "Email address"
// @Param name query string true "User's name"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Router /auth/verify-email [get]
func (ac AuthController) VerifyEmail(c *gin.Context) {
	email := c.Query("email")
	name := c.Query("name")
	if email == "" || name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email parameter is required"})
		return
	}

	// Pass request to AuthService for email verification
	err := authService.VerifyEmail(email, name)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Otp Sent To Your Email Successfully"})
}

// VerifyRegistrationOTP handles OTP verification
// @Summary Verify OTP
// @Description Verify OTP sent to the user's email
// @Tags Auth
// @Accept json
// @Produce json
// @Param email query string true "Email address"
// @Param otp query string true "One-Time Password"
// @Param name query string true "User's name"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Router /auth/verify-otp [get]
func (ac AuthController) VerifyRegistrationOTP(c *gin.Context) {
	email := c.Query("email")
	otp := c.Query("otp")
	name := c.Query("name")
	if email == "" || otp == "" || name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and OTP parameters are required"})
		return
	}

	// Pass request to AuthService for OTP verification
	err := authService.VerifyRegistrationOTP(name, email, otp)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Otp Verified Successfully"})
}