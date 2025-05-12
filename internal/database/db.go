package database

import (
	"task-tracker/config"
	"task-tracker/internal/logger"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var db *gorm.DB
	var err error

	tags := []string{"db", "sql", "pgsql"}

	maxAttempts := 10
	for attempts := 1; attempts <= maxAttempts; attempts++ {
		db, err = gorm.Open(postgres.Open(config.Cfg.Database.SQLDBUrl), &gorm.Config{})
		if err == nil {
			sqlDB, sqlErr := db.DB()
			if sqlErr == nil {
				if pingErr := sqlDB.Ping(); pingErr == nil {
					logger.GlobalLogger.Info("Connection successful", &logger.LogFields{"tags": tags, "attempts": attempts})
					break
				} else {
					err = pingErr
				}
			} else {
				err = sqlErr
			}
		}
		logger.GlobalLogger.Error("Connection failed", &logger.LogFields{"tags": tags, "attempts": attempts, "error": err})
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		logger.GlobalLogger.Panic("PostgreSQL connection failed", &logger.LogFields{"tags": tags, "error": err})
	}

	DB = db
	logger.GlobalLogger.Info("PostgreSQL connection established successfully.", &logger.LogFields{"tags": tags})
}
