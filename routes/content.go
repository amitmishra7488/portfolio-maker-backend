package routes

import (
	"portfolio-user-service/controller"
	"portfolio-user-service/middleware"

	"github.com/gin-gonic/gin"
)

// routes/content_type_routes.go

var (
	contentTypeController = new(controller.ContentTypeController)
	contentItemController = new(controller.ContentItemController)
)

func ContentRoutes(router *gin.Engine) {
	ctGroup := router.Group("/content")
	ctGroup.Use(middleware.JWTAuthMiddleware())

	// content types
	ctGroup.POST("/", contentTypeController.CreateContentType)
	ctGroup.GET("/", contentTypeController.GetAllContentTypes)
	// ctGroup.GET("/:id", contentTypeController.GetContentTypeByID)
	// ctGroup.PUT("/:id", contentTypeController.UpdateContentType)
	// ctGroup.DELETE("/:id", contentTypeController.DeleteContentType)

	// content item
	ctGroup.POST("/item", contentItemController.CreateContentItem)

}
