package database

import (
	"log"
	"task-tracker/internal/model"
	"time"

	"gorm.io/gorm"
)

func AutoMigrateAndSeed() {
	err := DB.AutoMigrate(
		&model.Role{},
		&model.User{},
		&model.Task{},
		&model.Log{},
	)
	if err != nil {
		log.Panicf("[MIGRATION] Migration başarısız: %v", err)
	} else {

		log.Println("[MIGRATION] Migration başarıyla tamamlandı.")
	}

	roleSeed()
	userSeed()
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
			log.Println("[SEED] Role seed sırasında hata:", err)
		}

		log.Println("[SEED] Role seed işlemi tamamlandı")

	}
}

func userSeed() {
	var count int64
	DB.Model(&model.User{}).Count(&count)

	if count == 0 {
		user := model.User{
			Model:                 gorm.Model{ID: 1},
			Name:                  "Admin User",
			Email:                 "admin@example.com",
			Password:              "hashed_password", // [NOTE] Burayı user register işleminde düzelt
			UserKey:               "",                // [NOTE] Burayı uuidkey ile user register işleminde düzelt
			RefreshToken:          "dummy_refresh_token",
			RefreshTokenExpiresAt: time.Now().AddDate(-1, 0, 0),
			RoleID:                1,
		}

		if err := DB.Create(&user).Error; err != nil {
			log.Println("[SEED] User seed sırasında hata:", err)
		} else {
			log.Println("[SEED] Admin kullanıcı seed işlemi tamamlandı")
		}
	}
}
