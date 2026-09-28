// backend/internal/config/config.go

// config.go contains application configuration loading for the backend service.
// It reads required environment variables from the local .env file and validates
// that the values needed to start the application are available.

package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

// GetSecrets loads and validates the backend configuration.
//
// The function reads the local .env file, retrieves the OpenAI API key and
// server port from environment variables, and returns an error if either
// required value is missing.
//
// Returns:
//   - openApiKey: API key used to authenticate with the OpenAI API.
//   - portNumber: port on which the backend HTTP server should listen.
//   - error: non-nil when the .env file cannot be loaded or required values are missing.
func GetSecrets() (string, string, error) {
	// Load .env from the current working directory
	if err := godotenv.Load(".env"); err != nil {
		return "", "", fmt.Errorf("error loading .env file: %w", err)
	}

	openApiKey := os.Getenv("OPENAI_API_KEY")
	portNumber := os.Getenv("PORT")

	if openApiKey == "" {
		return "", "", fmt.Errorf("OPENAI_API_KEY is missing")
	}

	if portNumber == "" {
		return "", "", fmt.Errorf("PORT is missing")
	}

	return openApiKey, portNumber, nil
}
