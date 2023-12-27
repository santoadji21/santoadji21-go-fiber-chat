package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config stores all configuration of the application.
// The values are read by godotenv from a .env file.
type Config struct {
    DBUsername string
    DBPassword string
    DBHost     string
    DBName     string
    DBPort     string
	RDHost 	   string
    RDPort     string
    RDPassword string
}

type JwtConfig struct {
    SecretKey string
}

// LoadConfig reads configuration from .env file and environment variables.
func DbCfg() Config {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    return Config{
        DBUsername: os.Getenv("DB_USERNAME"),
        DBPassword: os.Getenv("DB_PASSWORD"),
        DBHost:     os.Getenv("DB_HOST"),
        DBName:     os.Getenv("DB_NAME"),
        DBPort:     os.Getenv("DB_PORT"),
		RDHost:     os.Getenv("RD_HOST"),
        RDPort:     os.Getenv("RD_PORT"),
        RDPassword: os.Getenv("RD_PASSWORD"),
    }
}
