package controller

import (
	"net/http"
	"portfolio-user-service/repository/address/models"
	service "portfolio-user-service/services/address"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AddressController struct {
	AddressService *service.AddressService
	Log            *zap.Logger
}

// Constructor
func NewAddressController(addressService *service.AddressService, log *zap.Logger) *AddressController {
	return &AddressController{
		AddressService: addressService,
		Log:            log,
	}
}

// POST /api/user/address
func (ac *AddressController) CreateAddress(c *gin.Context) {
	userIDRaw, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID := userIDRaw.(uint)

	var input models.AddressInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	address, err := ac.AddressService.CreateAddress(userID, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Address created successfully", "data": address})
}

func (ac AddressController) UpdateExistingAddress(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	// Get :addressID param and convert to uint
	addressIDParam := c.Param("addressID")
	addressID, err := strconv.ParseUint(addressIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid address ID"})
		return
	}

	// Bind input
	var input models.AddressInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Call service layer
	address, err := ac.AddressService.UpdateAddress(userID, uint(addressID), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, address)
}

func (ac AddressController) GetAllAddresses(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	addresses, err := ac.AddressService.GetAllAddresses(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"addresses": addresses})
}
