package util

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"task-tracker/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte(config.Cfg.JWT.Secret)

func GenerateJWT(userID uint, roleName string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    roleName,
		"exp":     time.Now().Add(time.Duration(config.Cfg.JWT.ExpiryMinute) * time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateRefreshToken() (string, time.Time, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", time.Time{}, err
	}
	token := base64.URLEncoding.EncodeToString(bytes)

	ttl, _ := time.ParseDuration(config.Cfg.JWT.RefreshTTL)
	expiry := time.Now().Add(ttl)
	return token, expiry, nil
}

func HashToken(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}
