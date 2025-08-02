package env

import (
	"os"
	"strconv"
)

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func GetEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err != nil {
		return value
	}
	return defaultValue
}

func GetEnvAsBool(key string, defaultValue bool) bool {
	valueStr := getEnv(key, "")
	if value, err := strconv.ParseBool(valueStr); err != nil {
		return value
	}
	return defaultValue
}

func GetEnvAsString(key, defaultValue string) string {
	return getEnv(key, defaultValue)
}
