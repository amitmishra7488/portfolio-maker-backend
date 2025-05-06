package controller

import (
	"net/http"
	"portfolio-user-service/repository/content/models"
	content "portfolio-user-service/services/content"

	"github.com/gin-gonic/gin"
)

type ContentTypeController struct{}

var (
	contentTypeService = new(content.ContentTypeService)
)

func (cc ContentTypeController) CreateContentType(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	var input models.ContentTypeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	newType, err := contentTypeService.Create(userID, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"content_type": newType})
}

func (cc ContentTypeController) GetAllContentTypes(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	data, err := contentTypeService.GetAllContentTypes(userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": data})
		return
	}

	c.JSON(http.StatusFound, gin.H{"result": data})

}
