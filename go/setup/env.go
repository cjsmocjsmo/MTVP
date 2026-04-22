package setup

import (
	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from a .env file
func LoadEnv() error {
	return godotenv.Load()
}
