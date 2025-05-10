package config

import "time"

type Config struct {
	AppEnv        string
	ServerPort    string
	Database      Database
	JWT           JWTConfig
	RateLimit     RateLimitConfig
	TaskRateLimit RateLimitConfig
}

type Database struct {
	SQLDBUrl string
}

type JWTConfig struct {
	Secret       string
	ExpiryMinute int
	RefreshTTL   string
}

type RateLimitConfig struct {
	Max        int
	Expiration time.Duration
}
