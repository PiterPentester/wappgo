package utils

import (
	"os"
)

var (
	Host     = getEnv("REDIS_HOST", "localhost")
	Port     = string(getEnv("REDIS_PORT", "6379"))
	Password = getEnv("REDIS_PASSWORD", "")
)

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
