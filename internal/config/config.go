package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName     string
	AppEnv      string
	AppPort     string
	AppHost     string
	AppBaseURL  string
	APIBasePath string
	APIVersion  string
	APIPrefix   string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	DBTimeZone string

	JWTSecret         string
	JWTExpiresInHours string
	UploadDir         string
	UploadDriver      string
	AssetBaseURL      string
	AdminSeedName     string
	AdminSeedEmail    string
	AdminSeedPassword string
}

func Load() Config {
	_ = godotenv.Load()
	apiBasePath := getEnv("API_BASE_PATH", "/api")
	apiVersion := getEnv("API_VERSION", "v1")
	cfg := Config{
		AppName:           getEnv("APP_NAME", "cms-be"),
		AppEnv:            getEnv("APP_ENV", "development"),
		AppPort:           getEnv("APP_PORT", "8080"),
		AppHost:           getEnv("APP_HOST", "0.0.0.0"),
		AppBaseURL:        getEnv("APP_BASE_URL", "http://localhost:8080"),
		APIBasePath:       apiBasePath,
		APIVersion:        apiVersion,
		APIPrefix:         getEnv("API_PREFIX", apiBasePath+"/"+apiVersion),
		DBHost:            getEnv("DB_HOST", "127.0.0.1"),
		DBPort:            getEnv("DB_PORT", "5432"),
		DBUser:            getEnv("DB_USER", "postgres"),
		DBPassword:        getEnv("DB_PASSWORD", ""),
		DBName:            getEnv("DB_NAME", "cms_be"),
		DBSSLMode:         getEnv("DB_SSLMODE", "disable"),
		DBTimeZone:        getEnv("DB_TIMEZONE", "Asia/Jakarta"),
		JWTSecret:         getEnv("JWT_SECRET", "change-me"),
		JWTExpiresInHours: getEnv("JWT_EXPIRES_IN_HOURS", "24"),
		UploadDir:         getEnv("UPLOAD_DIR", "storage/uploads"),
		UploadDriver:      getEnv("UPLOAD_DRIVER", "local"),
		AssetBaseURL:      getEnv("ASSET_BASE_URL", ""),
		AdminSeedName:     getEnv("ADMIN_SEED_NAME", "Admin"),
		AdminSeedEmail:    getEnv("ADMIN_SEED_EMAIL", ""),
		AdminSeedPassword: getEnv("ADMIN_SEED_PASSWORD", ""),
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
