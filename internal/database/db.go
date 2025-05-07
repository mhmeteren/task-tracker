package database

import (
	"log"
	"task-tracker/config"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var db *gorm.DB
	var err error

	maxAttempts := 10
	for attempts := 1; attempts <= maxAttempts; attempts++ {
		db, err = gorm.Open(postgres.Open(config.Cfg.Database.SQLDBUrl), &gorm.Config{})
		if err == nil {
			sqlDB, sqlErr := db.DB()
			if sqlErr == nil {
				if pingErr := sqlDB.Ping(); pingErr == nil {
					log.Printf("[DB] Connection successful (attempt %d).", attempts)
					break
				} else {
					err = pingErr
				}
			} else {
				err = sqlErr
			}
		}

		log.Printf("[DB] Connection attempt %d failed: %v", attempts, err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Panicf("[DB] PostgreSQL connection failed: %v", err)
	}

	DB = db
	log.Println("[DB] PostgreSQL connection established successfully.")
}
