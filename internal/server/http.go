package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rawatsaheb5/blog-backend-with-go/internal/config"
	"github.com/rawatsaheb5/blog-backend-with-go/internal/middleware"
	"github.com/rawatsaheb5/blog-backend-with-go/internal/modules/user"
	"github.com/rawatsaheb5/blog-backend-with-go/pkg/logger"
	"gorm.io/gorm"
)

func Start(cfg config.Config, db *gorm.DB) {
	// Set Gin mode based on environment
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.New()

	// Add zap logger and recovery middleware
	r.Use(logger.GinLogger())
	r.Use(logger.GinRecovery())

	api := r.Group("/api")

	// ===== PUBLIC ROUTES (No authentication required) =====
	public := api.Group("")
	{
		// User authentication routes (public)
		user.RegisterRoutes(public, db, cfg.JWTKey)
		
		// Add other public routes here
		// public.GET("/posts", GetPublicPosts) // Example: public post listing
	}

	// ===== PROTECTED ROUTES (Authentication required) =====
	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware(cfg.JWTKey)) // Apply auth middleware only to this group
	{
		// Post routes (protected)

		// User profile routes (protected)
		// protected.GET("/users/me", GetCurrentUser)
		// protected.PUT("/users/me", UpdateCurrentUser)
		
		// Add other protected routes here
	}

	r.Run(":" + cfg.Port)
}