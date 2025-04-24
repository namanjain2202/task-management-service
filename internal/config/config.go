package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds the configuration settings for the application.
type Config struct {
	DatabaseURL string
	ServerPort  string
}

// LoadConfig loads the configuration from environment variables or .env file.
func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using default values")
	}

	return &Config{
		DatabaseURL: getEnv("DATABASE_URL", "default_database_url"),
		ServerPort:  getEnv("SERVER_PORT", "8080"),
	}
}

// getEnv retrieves the value of the environment variable named by the key.
// If the variable is not present, it returns the fallback value.
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}