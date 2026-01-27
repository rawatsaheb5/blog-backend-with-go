package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rawatsaheb5/blog-backend-with-go/internal/config"
	"github.com/rawatsaheb5/blog-backend-with-go/internal/middleware"
	"github.com/rawatsaheb5/blog-backend-with-go/internal/modules/post"
	"github.com/rawatsaheb5/blog-backend-with-go/internal/modules/user"
	"gorm.io/gorm"
)

func Start(cfg config.Config, db *gorm.DB) {
	r := gin.New()

	// Optional: Add logging middleware for all routes
	// r.Use(middleware.Logging())

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
		protected.Group("/posts").Use(middleware.AuthMiddleware(cfg.JWTKey))
		post.CreatePostRoutes(protected.Group("/posts"), db)

		// User profile routes (protected)
		// protected.GET("/users/me", GetCurrentUser)
		// protected.PUT("/users/me", UpdateCurrentUser)
		
		// Add other protected routes here
	}

	r.Run(":" + cfg.Port)
}