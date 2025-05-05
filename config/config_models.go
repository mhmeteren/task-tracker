package config

type Config struct {
	AppEnv     string
	ServerPort string
	Database   Database
}

type Database struct {
	SQLDBUrl string
}
