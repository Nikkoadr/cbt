package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"cbt/config"
	"cbt/models"
)

// GenerateJWT generates a new JWT token
func GenerateJWT(user models.User, cfg config.Config) (string, error) {
	claims := jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * time.Duration(cfg.JWTExpire)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(cfg.JWTSecret))
	if err != nil {
		return "", err
	}

	return t, nil
}
