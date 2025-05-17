package config

import (
	"os"
	"strconv"
	"task-tracker/internal/logger"
	"time"

	"github.com/joho/godotenv"
)

var Cfg Config

func LoadConfig() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	envFile := ".env." + env
	err := godotenv.Load(".env") // load common first
	if err != nil {
		logger.GlobalLogger.Fatal("Error loading environment", &logger.LogFields{"file": ".env", "error": err})
	}

	err = godotenv.Load(envFile) // override with environment-specific
	if err != nil {
		logger.GlobalLogger.Fatal("Error loading environment", &logger.LogFields{"file": envFile, "error": err})
	}

	Cfg = Config{
		AppEnv:     env,
		ServerPort: os.Getenv("SERVER_PORT"),
		Database: Database{
			SQLDBUrl: os.Getenv("SQL_DATABASE_URL"),
		},
		JWT:           loadJWTConfig(),
		RateLimit:     loadRateLimitConfig("RATE_LIMIT_MAX", "RATE_LIMIT_EXPIRATION"),
		TaskRateLimit: loadRateLimitConfig("TASK_RATE_LIMIT_MAX", "TASK_RATE_LIMIT_EXPIRATION"),
	}
}

func loadJWTConfig() JWTConfig {
	ExpiryMinute, err := strconv.Atoi(os.Getenv("JWT_EXPIRY_MINUTE"))
	if err != nil {
		logger.GlobalLogger.Fatal("Error Expiry Minute", &logger.LogFields{"error": err})
	}

	return JWTConfig{
		Secret:       os.Getenv("JWT_SECRET"),
		ExpiryMinute: ExpiryMinute,
		RefreshTTL:   os.Getenv("REFRESH_TTL"),
	}
}

func loadRateLimitConfig(max_config_key, expiration_config_key string) RateLimitConfig {
	maxStr := os.Getenv(max_config_key)
	max, err := strconv.Atoi(maxStr)
	if err != nil || max <= 0 {
		max = 5
	}

	expStr := os.Getenv(expiration_config_key)
	exp, err := time.ParseDuration(expStr)
	if err != nil {
		exp = time.Second
	}

	return RateLimitConfig{Max: max, Expiration: exp}
}
