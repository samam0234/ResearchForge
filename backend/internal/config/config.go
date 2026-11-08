package config

import (
	"os"
)

type Config struct {
	Port         string
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	RedisHost    string
	RedisPort    string
	RedisPass    string
	AIServerURL  string
	JWTSecret    string
}

func LoadConfig() *Config {
	return &Config{
		Port:        getEnv("PORT", "8080"),
		DBHost:      getEnv("DB_HOST", "postgres"),
		DBPort:      getEnv("DB_PORT", "5432"),
		DBUser:      getEnv("DB_USER", "postgres"),
		DBPassword:  getEnv("DB_PASSWORD", "postgres_secret_password"),
		DBName:      getEnv("DB_NAME", "researchforge"),
		RedisHost:   getEnv("REDIS_HOST", "redis"),
		RedisPort:   getEnv("REDIS_PORT", "6379"),
		RedisPass:   getEnv("REDIS_PASSWORD", ""),
		AIServerURL: getEnv("AI_SERVER_URL", "http://ai-server:8000"),
		JWTSecret:   getEnv("JWT_SECRET", "super-secret-jwt-token-change-in-production"),
	}
}

func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}
