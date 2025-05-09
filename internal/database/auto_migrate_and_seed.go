package database

import (
	"log"
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
		log.Panicf("[MIGRATION] Migration failed: %v", err)
	} else {
		log.Println("[MIGRATION] Migration completed successfully.")

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
			log.Println("[SEED] Error occurred during role seeding:", err)
		} else {
			log.Println("[SEED] Role seeding completed successfully")
		}

	}
}

func userSeed() {
	var count int64
	DB.Model(&model.User{}).Count(&count)

	if count == 0 {

		hashedPassword, err := util.HashPassword("tasktracker")
		if err != nil {
			log.Println("[SEED] User seed sırasında hata:", err)
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
			log.Println("[SEED] Error occurred during user seeding:", err)
		} else {
			log.Println("[SEED] Admin user seeding completed successfully")
		}
	}
}
