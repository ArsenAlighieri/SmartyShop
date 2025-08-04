package config

import "os"

// GetGeminiAPIKey returns the Gemini API key from the environment variables.
func GetGeminiAPIKey() string {
	return os.Getenv("GEMINI_API_KEY")
}
