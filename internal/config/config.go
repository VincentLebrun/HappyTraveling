package config

import (
	"fmt"
	"os"
)

type Config struct {
	Database DatabaseConfig
}

type DatabaseConfig struct {
	User     string
	Password string
	Host     string
	Port     string // Added Port
	Name     string
}

func (dc *DatabaseConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dc.User, dc.Password, dc.Host, dc.Port, dc.Name)
}

func LoadConfiguration() Config {
	return Config{
		Database: DatabaseConfig{
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Name:     os.Getenv("DB_NAME"),
		},
	}
}
