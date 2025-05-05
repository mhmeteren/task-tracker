package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name                  string    `gorm:"not null"`
	Email                 string    `gorm:"uniqueIndex;not null"`
	Password              string    `gorm:"not null"`
	UserKey               string    `gorm:"uniqueIndex;size:32"`
	RefreshToken          string    `gorm:"uniqueIndex;not null"`
	RefreshTokenExpiresAt time.Time `gorm:"not null"`
	RoleID                uint
	Role                  Role   `gorm:"foreignKey:RoleID"`
	Tasks                 []Task `gorm:"foreignKey:UserID"`
}
