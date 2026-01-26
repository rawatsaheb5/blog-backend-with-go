package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/rawatsaheb5/blog-backend-with-go/internal/config"
	"github.com/rawatsaheb5/blog-backend-with-go/internal/database"
	"github.com/rawatsaheb5/blog-backend-with-go/internal/modules/user"
	"github.com/rawatsaheb5/blog-backend-with-go/internal/server"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	cfg := config.LoadConfig()
	db := database.Connect(cfg)

	// Run database migrations
	if err := db.AutoMigrate(&user.User{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	log.Println("Database migration completed successfully")

	server.Start(cfg, db)
}