package config

import (
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	DB   DBConfig
	HTTP HTTPConfig
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	return &Config{
		DB:   LoadDBConfig(),
		HTTP: LoadHTTPConfig(),
	}
}
