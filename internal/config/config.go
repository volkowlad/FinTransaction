package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	HTTPServer
	DB
}

type HTTPServer struct {
	Address string
}

type DB struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
	SSLMode  string
}

func InitConfig() *Config {
	err := godotenv.Load("config.env")
	if err != nil {
		fmt.Errorf("Error loading config.env file", err.Error())
	}

	return &Config{
		HTTPServer: HTTPServer{
			Address: os.Getenv("HTTP_PORT"),
		},

		DB: DB{
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			DBName:   os.Getenv("DB_NAME"),
			SSLMode:  os.Getenv("DB_SSLMODE"),
		},
	}
}
