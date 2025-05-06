package controller

import (
	"fmt"
	"net/http"
	"portfolio-user-service/repository/content/models"
	"portfolio-user-service/services/content"

	"github.com/gin-gonic/gin"
)

type ContentItemController struct{}

var (
	contentItemService = new(content.ContentItemService)
)

func (ci ContentItemController) CreateContentItem(c *gin.Context) {
	fmt.Println("coming here")
	var req models.ContentItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item := models.ContentItem{
		ContentTypeID: req.ContentTypeID,
		Title:         req.Title,
		Body:          req.Body,
	}

	if err := contentItemService.CreateContentItem(&item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, item)
}
