package routes

import (
	"github.com/gin-gonic/gin"
)

// InitializeRoutes registers all routes in the router
func InitializeRoutes(router *gin.Engine) {
	AuthRoutes(router)
	AddressRoutes(router)
	ContentRoutes(router)
	
}
