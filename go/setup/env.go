package setup

import (
	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from a .env file
func LoadEnv() error {
	// if this fails try "../../env/.env"
	return godotenv.Load("../env/.env")
}
