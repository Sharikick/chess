package config

import (
	"chess/internal/lib/env"
	"fmt"
)

type (
	HTTPConfig struct {
		Host string
		Port string
	}

	DatabaseConfig struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}

	LoggerConfig struct {
		Level string
	}

	Config struct {
		Database DatabaseConfig
		Logger   LoggerConfig
		HTTP     HTTPConfig
	}
)

func (db DatabaseConfig) GetConn() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", db.User, db.Password, db.Host, db.Port, db.Name)
}

func New() (*Config, error) {
	database := DatabaseConfig{
		Host:     env.GetEnvAsString("DB_HOST", "database"),
		Port:     env.GetEnvAsString("DB_PORT", "5432"),
		User:     env.GetEnvAsString("DB_USER", "tsyden"),
		Password: env.GetEnvAsString("DB_PASSWORD", "chess"),
		Name:     env.GetEnvAsString("DB_DATABASe", "chess_db"),
	}

	logger := LoggerConfig{
		Level: env.GetEnvAsString("LOG_LEVEL", "INFO"),
	}

	http := HTTPConfig{
		Host: env.GetEnvAsString("HTTP_HOST", "127.0.0.1"),
		Port: env.GetEnvAsString("HTTP_PORT", "8080"),
	}

	return &Config{
		Logger:   logger,
		HTTP:     http,
		Database: database,
	}, nil
}
