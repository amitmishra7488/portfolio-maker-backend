package routes

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// InitializeRoutes registers all routes in the router
func InitializeRoutes(router *gin.Engine, db *gorm.DB, logger *zap.Logger) {
	// health + root
	hostname, _ := os.Hostname()
	// health + root
	router.GET("/", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "server is running","server":  hostname,}) })
	router.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "ok"}) })

	// Register all routes here
	AuthRoutes(router, db, logger)
	AddressRoutes(router, db, logger)
	ContentRoutes(router, db, logger)
}
