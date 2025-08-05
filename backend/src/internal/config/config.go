package config

import (
	"fmt"
	"os"
)

type Config struct {
	Database DatabaseConfig
	Logger   LoggerConfig
	HTTP     HTTPConfig
}

func New(isProd bool) (*Config, error) {
	getEnv := func(key, defaultValue string) (string, error) {
		if value, exists := os.LookupEnv(key); exists {
			return value, nil
		}

		if isProd {
			return "", fmt.Errorf("не задано значение для переменной: %s", key)
		}
		return defaultValue, nil
	}

	database := DatabaseConfig{}
	if err := database.fill(getEnv); err != nil {
		return nil, err
	}

	logger := LoggerConfig{}
	if err := logger.fill(getEnv); err != nil {
		return nil, err
	}

	http := HTTPConfig{}
	if err := http.fill(getEnv); err != nil {
		return nil, err
	}

	return &Config{
		Logger:   logger,
		HTTP:     http,
		Database: database,
	}, nil
}
