package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rawatsaheb5/blog-backend-with-go/internal/config"
	"github.com/rawatsaheb5/blog-backend-with-go/internal/modules/user"
	"gorm.io/gorm"
)

func Start(cfg config.Config, db *gorm.DB) {
	r := gin.New()

	// TODO: Add middleware when implemented
	// r.Use(middleware.Logging())
	// r.Use(middleware.Auth())

	api := r.Group("/api")
	user.RegisterRoutes(api, db)

	r.Run(":" + cfg.Port)
}