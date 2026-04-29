package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName    string
	AppEnv     string
	AppPort    string
	AppHost    string
	AppBaseURL string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	JWTSecret string
	UploadDir string
}

func Load() Config {
	_ = godotenv.Load()
	cfg := Config{
		AppName:    getEnv("APP_NAME", "cms-be"),
		AppEnv:     getEnv("APP_ENV", "development"),
		AppPort:    getEnv("APP_PORT", "8080"),
		AppHost:    getEnv("APP_HOST", "0.0.0.0"),
		AppBaseURL: getEnv("APP_BASE_URL", "http://localhost:8080"),
		DBHost:     getEnv("DB_HOST", "127.0.0.1"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "cms_be"),
		DBSSLMode:  getEnv("DB_SSLMODE", "disable"),
		JWTSecret:  getEnv("JWT_SECRET", "change-me"),
		UploadDir:  getEnv("UPLOAD_DIR", "storage/uploads"),
	}
	if cfg.JWTSecret == "change-me" {
		log.Println("warning: JWT_SECRET still uses default value")
	}
	return cfg
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
