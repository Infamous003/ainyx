package config

import (
	"errors"
	"os"
)

type Config struct {
	Port string
	DSN  string
}

func Load() (*Config, error) {
	cfg := &Config{
		Port: getEnv("PORT", "8080"),
		DSN:  getEnv("DB_DSN", ""),
	}
	if cfg.DSN == "" {
		return nil, errors.New("DB_DSN environment variable is required")
	}
	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}
	return val
}
