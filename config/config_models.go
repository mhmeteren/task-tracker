package config

type Config struct {
	AppEnv     string
	ServerPort string
	Database   Database
	JWT        JWTConfig
}

type Database struct {
	SQLDBUrl string
}

type JWTConfig struct {
	Secret       string
	ExpiryMinute int
	RefreshTTL   string
}
