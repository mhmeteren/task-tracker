package config

import (
	"log"
	"os"

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

	Cfg = Config{
		AppEnv:     env,
		ServerPort: os.Getenv("SERVER_PORT"),
	}
}
