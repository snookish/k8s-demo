package config

import (
	"os"
	"strconv"
	"time"
)

func getEnvIntOrFallback(key string, fallback int) int {
	env, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	if parsed, err := strconv.Atoi(env); err == nil {
		return parsed
	}
	return fallback
}

func getEnvOrFallback(key, fallback string) string {
	if value, ok := os.LookupEnv(key); value != "" && ok {
		return value
	}
	return fallback
}

func getDurationOrFallback(key, fallback string) time.Duration {
	if value, ok := os.LookupEnv(key); value != "" && ok {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}

	duration, err := time.ParseDuration(fallback)
	if err != nil {
		return time.Duration(0)
	}
	return duration
}
