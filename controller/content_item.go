package controller

import (
	"fmt"
	"net/http"
	"portfolio-user-service/repository/content/models"
	"portfolio-user-service/services/content"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ContentItemController struct {
	ContentItemService *content.ContentItemService
	Log                *zap.Logger
}

func NewContentItemController(contentSevice *content.ContentItemService, log *zap.Logger) *ContentItemController {
	return &ContentItemController{
		ContentItemService: contentSevice,
		Log:                log,
	}
}

func (ci *ContentItemController) CreateContentItem(c *gin.Context) {
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

	if err := ci.ContentItemService.CreateContentItem(&item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, item)
}
