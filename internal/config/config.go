package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	HTTPport    string
}

func New() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to load .env file, system variables are used")
	}

	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		log.Fatal("The DATABASE_URL environment variable is not set")
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		log.Println("The HTTP_PORT environment variable is not set, the default port 3000 is used")
		httpPort = ":3000"
	}

	return &Config{
		DatabaseURL: databaseUrl,
		HTTPport:    httpPort,
	}
}
