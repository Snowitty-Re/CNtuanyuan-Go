package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/config"
)

// GenerateToken 生成JWT Token
func GenerateToken(userID uint, username string) (string, error) {
	cfg := config.Get()
	expirationTime := time.Now().Add(time.Duration(cfg.JWT.Expiration) * time.Hour)

	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"exp":      expirationTime.Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWT.Secret))
}

// ParseToken 解析JWT Token
func ParseToken(tokenString string) (jwt.MapClaims, error) {
	cfg := config.Get()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWT.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}

