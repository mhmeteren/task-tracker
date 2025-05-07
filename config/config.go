package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Cfg Config

func LoadConfig() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	envFile := ".env." + env
	_ = godotenv.Load(".env")         // load common first
	err := godotenv.Overload(envFile) // override with environment-specific
	if err != nil {
		log.Fatalf("Error loading %s file: %v", envFile, err)
	}

	ExpiryMinute, err := strconv.Atoi(os.Getenv("JWT_EXPIRY_MINUTE"))
	if err != nil {
		log.Fatalf("Error Expiry Minute %s file: %v", envFile, err)
	}

	Cfg = Config{
		AppEnv:     env,
		ServerPort: os.Getenv("SERVER_PORT"),
		Database: Database{
			SQLDBUrl: os.Getenv("SQL_DATABASE_URL"),
		},
		JWT: JWTConfig{
			Secret:       os.Getenv("JWT_SECRET"),
			ExpiryMinute: ExpiryMinute,
			RefreshTTL:   os.Getenv("REFRESH_TTL"),
		},
	}
}
