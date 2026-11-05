package config

import "os"

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func JWTSecret() string {
	return getEnv("JWT_SECRET", "supersecret")
}

func DBPath() string {
	return getEnv("DB_PATH", "api.db")
}

func ServerPort() string {
	return getEnv("SERVER_PORT", ":8080")
}
