package configs

import (
	"log"
	"os"
)

func GetEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Printf("Environment variable %s not found, using default value: %s", key, defaultValue)
		return defaultValue
	}
	return value
}
