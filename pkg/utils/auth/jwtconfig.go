package auth

import (
	"os"
	"time"
)

type JWTConfig struct {
	SecretKey     string
	ExpirationTTL time.Duration
}

func LoadJWTConfig() *JWTConfig {
	return &JWTConfig{
		SecretKey:     getEnv("JWT_SECRET", "default_secret_key"),
		ExpirationTTL: getEnvAsDuration("JWT_TTL", time.Hour*24),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		d, err := time.ParseDuration(value)
		if err == nil {
			return d
		}
	}
	return defaultValue
}
