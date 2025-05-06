package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHOST     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	MONGOURI   string
	MONGODB    string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, fmt.Errorf("failed to load environment variables: %w", err)
	}
	config := &Config{
		DBHOST:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		MONGOURI:   os.Getenv("MONGO_URI"),
		MONGODB:    os.Getenv("MONGO_DB"),
	}
	return config, nil
}
