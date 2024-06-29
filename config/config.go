package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/cast"
	"log"
)

type Config struct {
	HTTP_PORT   string
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
}

func Load() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := &Config{}

	config.HTTP_PORT = cast.ToString("HTTP_PORT")
	config.DB_HOST = cast.ToString("DB_HOST")
	config.DB_PORT = cast.ToString("DB_PORT")
	config.DB_USER = cast.ToString("DB_USER")
	config.DB_PASSWORD = cast.ToString("DB_PASSWORD")
	config.DB_NAME = cast.ToString("DB_NAME")

	return config
}
