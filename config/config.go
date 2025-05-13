package config

import (
	"os"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	DBHOST     string

	MONGOUser string
	MONGOPass string
	MONGODB   string
	MONGOURI  string

	Port string
}

func LoadConfig() (*Config, error) {
	return &Config{
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     os.Getenv("DB_PORT"),
		DBHOST:     os.Getenv("DB_HOST"),

		MONGOUser: os.Getenv("MONGO_USER"),
		MONGOPass: os.Getenv("MONGO_PASSWORD"),
		MONGODB:   os.Getenv("MONGO_DB"),
		MONGOURI:  os.Getenv("MONGOURI"),

		Port: os.Getenv("PORT"),
	}, nil
}
