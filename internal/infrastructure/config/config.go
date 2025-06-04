package config

import (
	"os"
)

type Config struct {
	RedisHost  string
	RedisPort  string
	ServerPort string
}

func Load() *Config {
	return &Config{
		RedisHost:  getEnv("REDIS_HOST", "localhost"),
		RedisPort:  getEnv("REDIS_PORT", "6379"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
