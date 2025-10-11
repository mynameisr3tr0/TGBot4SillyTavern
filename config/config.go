package config

import (
	"os"
	"strconv"
)

// Config holds the application configuration
type Config struct {
	TelegramToken  string
	SillyTavernURL string
	ChromiumPath   string
	HeadlessMode   bool
	Debug          bool
}

// Load loads configuration from environment variables
func Load() *Config {
	headless := true
	if os.Getenv("HEADLESS_MODE") == "false" {
		headless = false
	}

	debug := false
	if os.Getenv("DEBUG") == "true" {
		debug = true
	}

	return &Config{
		TelegramToken:  getEnv("TELEGRAM_BOT_TOKEN", ""),
		SillyTavernURL: getEnv("SILLYTAVERN_URL", "http://localhost:8000"),
		ChromiumPath:   getEnv("CHROMIUM_PATH", ""),
		HeadlessMode:   headless,
		Debug:          debug,
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}
