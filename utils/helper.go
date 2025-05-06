package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// LoadSecret loads the JWT secret key from the environment
func LoadEnvVar(key string) string {
	_ = godotenv.Load()
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("Environment variable '%s' not set in .env file", key))
	}
	return value
}
