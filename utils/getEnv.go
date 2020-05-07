package utils

import (
	"os"
)

// get variables from env
var (
	Host     = getEnv("REDIS_HOST", "localhost")
	Port     = string(getEnv("REDIS_PORT", "6379"))
	Password = getEnv("REDIS_PASSWORD", "")
)

// getEnv - takes key & default value and return value from env or, if key is empty - default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
