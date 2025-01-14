package config

import (
	"fmt"
	"os"
	"strconv"
	"tech-challenge/src/models"

	"github.com/joho/godotenv"
)

type Config struct {
	Database models.Database
	Server   models.Server
}

func loadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading config: %v", err)
	}

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		dbPort = 5432
	}

	serverPort, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		serverPort = 8080
	}

	config := &Config{
		Database: models.Database{
			Host:     os.Getenv("DB_HOST"),
			Port:     dbPort,
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
			SSLMode:  os.Getenv("DB_SSLMODE"),
		},
		Server: models.Server{
			Port: serverPort,
			Host: os.Getenv("SERVER_HOST"),
		},
	}
	return config, nil
}

func (c *Config) getDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.Database.Host, c.Database.Port, c.Database.User, c.Database.Password, c.Database.DBName)
}
