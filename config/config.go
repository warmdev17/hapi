package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	IsProduction bool

	Database DatabaseConfig
	Redis    RedisConfig
	JWT      JWTConfig
}

type DatabaseConfig struct {
	Name     string
	Port     int
	User     string
	Password string
}

type RedisConfig struct {
	Port     int
	Password string
}

type JWTConfig struct {
	Secret     string
	AccessTTL  time.Duration
	RefreshTTL time.Duration
}

func getEnv(key string) string {
	return os.Getenv(key)
}

func mustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("missing required environment variable: %s", key))
	}

	return value
}

func getEnvInt(key string) int {
	value := mustGetEnv(key)
	n, err := strconv.Atoi(value)
	if err != nil {
		panic(fmt.Sprintf("invalid integer environment variable %s: %v", key, err))
	}
	return n
}

func getEnvDuration(key string) time.Duration {
	value := mustGetEnv(key)

	duration, err := time.ParseDuration(value)
	if err != nil {
		panic(fmt.Sprintf("invalid duration environment variable %s: %v", key, err))
	}
	return duration
}

func Load() *Config {
	_ = godotenv.Load()
	return &Config{
		IsProduction: getEnv("APP_ENV") == "development",
		Database: DatabaseConfig{
			Name:     mustGetEnv("DB_NAME"),
			Port:     getEnvInt("DB_PORT"),
			User:     mustGetEnv("DB_USER"),
			Password: mustGetEnv("DB_PASSWORD"),
		},
		Redis: RedisConfig{
			Port:     getEnvInt("REDIS_PORT"),
			Password: mustGetEnv("REDIS_PASSWORD"),
		},
		JWT: JWTConfig{
			Secret:     mustGetEnv("JWT_SECRET"),
			AccessTTL:  getEnvDuration("ACCESS_TTL"),
			RefreshTTL: getEnvDuration("REFRESH_TTL"),
		},
	}
}
