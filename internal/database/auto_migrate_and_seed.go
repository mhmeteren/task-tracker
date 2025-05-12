package database

import (
	"task-tracker/internal/logger"
	"task-tracker/internal/model"
	"task-tracker/internal/util"
	"time"
)

func AutoMigrateAndSeed() {
	err := DB.AutoMigrate(
		&model.Role{},
		&model.User{},
		&model.Task{},
		&model.Log{},
		&model.TaskNotification{},
	)
	if err != nil {
		logger.GlobalLogger.Panic("Migration failed", &logger.LogFields{"tags": []string{"migration", "sql", "pgsql"}, "error": err})
	} else {
		logger.GlobalLogger.Info("Migration completed successfully.", &logger.LogFields{"tags": []string{"migration", "sql", "pgsql"}})

		roleSeed()
		userSeed()
	}
}

func roleSeed() {
	var count int64
	DB.Model(&model.Role{}).Count(&count)

	if count == 0 {
		roles := []model.Role{
			{ID: 1, Name: "admin"},
			{ID: 2, Name: "user"},
		}
		if err := DB.Create(&roles).Error; err != nil {
			logger.GlobalLogger.Error("Error occurred during role seeding", &logger.LogFields{"tags": []string{"seed", "sql", "pgsql"}, "error": err})
		} else {
			logger.GlobalLogger.Info("Role seeding completed successfully.", &logger.LogFields{"tags": []string{"seed", "sql", "pgsql"}})
		}

	}
}

func userSeed() {
	var count int64
	DB.Model(&model.User{}).Count(&count)

	if count == 0 {

		hashedPassword, err := util.HashPassword("tasktracker")
		if err != nil {
			logger.GlobalLogger.Error("Error occurred during user seeding", &logger.LogFields{"tags": []string{"seed", "sql", "pgsql"}, "error": err})
		}

		user := model.User{
			Name:                  "Admin User",
			Email:                 "admin@mail.com",
			Password:              hashedPassword,
			RefreshToken:          nil,
			RefreshTokenExpiresAt: time.Now().AddDate(-1, 0, 0),
			RoleID:                1,
		}

		if err := DB.Create(&user).Error; err != nil {
			logger.GlobalLogger.Error("Error occurred during user seeding", &logger.LogFields{"tags": []string{"seed", "sql", "pgsql"}, "error": err})
		} else {
			logger.GlobalLogger.Info("Admin user seeding completed successfully", &logger.LogFields{"tags": []string{"seed", "sql", "pgsql"}})
		}
	}
}
