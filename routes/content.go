package routes

import (
	"portfolio-user-service/controller"
	"portfolio-user-service/middleware"
	"portfolio-user-service/repository/content"
	contentService "portfolio-user-service/services/content"
	 

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// routes/content_type_routes.go



func ContentRoutes(router *gin.Engine, db *gorm.DB, log *zap.Logger) {
	ctRepo := content.NewContentTypeRepository(db)
	ctService := contentService.NewContentTypeService(ctRepo, db, log)
	ctController := controller.NewContentTypeController(ctService, log)

	ctGroup := router.Group("/content")
	ctGroup.Use(middleware.JWTAuthMiddleware())

	// content types
	ctGroup.POST("/", ctController.CreateContentType)
	ctGroup.GET("/", ctController.GetAllContentTypes)
	// ctGroup.GET("/:id", ctController.GetContentTypeByID)
	// ctGroup.PUT("/:id", contentTypeController.UpdateContentType)
	// ctGroup.DELETE("/:id", contentTypeController.DeleteContentType)



	ciRepo := content.NewContentItemRepository(db)
	ciService := contentService.NewContentItemService(ciRepo, db, log)
	ciController := controller.NewContentItemController(ciService, log)
	// content item
	ctGroup.POST("/item", ciController.CreateContentItem)

}
