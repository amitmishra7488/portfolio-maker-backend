package main

import (
	"portfolio-user-service/config"
	"portfolio-user-service/middleware"
	"portfolio-user-service/pkg/logger"
	"portfolio-user-service/routes"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// 1) config + logger + db
	cfg := config.LoadConfig()

	logr, _ := logger.New(cfg.Env)
	defer logr.Sync()

	db, err := config.ConnectDatabase(cfg)
	if err != nil {
		logr.Fatal("db connect failed", zap.Error(err))
	}
	defer config.CloseDatabase(db)

	// 2) gin with production middlewares
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.RequestID())
	r.Use(middleware.Logging(logr))
	r.Use(middleware.Timeout(cfg.RequestTimeout))
	r.Use(middleware.CORS(cfg.AllowedOrigins))

	routes.InitializeRoutes(r, db, logr)
	logr.Info("🚀 Starting server", zap.String("port", cfg.Port))
	if err := r.Run(":" + cfg.Port); err != nil {
		logr.Fatal("❌ Server failed to start", zap.Error(err))
	}

}
