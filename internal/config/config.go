package config

import (
	"os"
)

type Config struct{
	DBUrl string
	Port string
	JWTKey string
	Environment string
}

func LoadConfig() Config{
	env := os.Getenv("ENV")
	if env == "" {
		env = os.Getenv("ENVIRONMENT")
	}
	if env == "" {
		env = "development" // Default to development
	}

	return Config{
		DBUrl: os.Getenv("DB_URL"),
		Port: os.Getenv("PORT"),
		JWTKey: os.Getenv("JWT_SECRET"),
		Environment: env,
	}
}