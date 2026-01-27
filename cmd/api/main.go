package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/rawatsaheb5/blog-backend-with-go/internal/config"
	"github.com/rawatsaheb5/blog-backend-with-go/internal/database"
	"github.com/rawatsaheb5/blog-backend-with-go/internal/modules/user"
	"github.com/rawatsaheb5/blog-backend-with-go/internal/server"
	"github.com/rawatsaheb5/blog-backend-with-go/pkg/logger"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	cfg := config.LoadConfig()

	// Initialize logger based on environment
	if err := logger.Init(cfg.Environment); err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	logger.Sugar.Infof("Starting application in %s environment", cfg.Environment)

	db := database.Connect(cfg)

	// Run database migrations
	if err := db.AutoMigrate(&user.User{}); err != nil {
		logger.Sugar.Fatalf("Failed to migrate database: %v", err)
	}
	logger.Sugar.Info("Database migration completed successfully")

	server.Start(cfg, db)
}