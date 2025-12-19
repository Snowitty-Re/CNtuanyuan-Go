package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/config"
)

// Auth JWT认证中间件
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   true,
				"message": "Authorization header is required",
			})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   true,
				"message": "Invalid authorization header format",
			})
			c.Abort()
			return
		}

		tokenString := parts[1]
		cfg := config.Get()

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.JWT.Secret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   true,
				"message": "Invalid or expired token",
			})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   true,
				"message": "Invalid token claims",
			})
			c.Abort()
			return
		}

		// 转换user_id为uint（JWT claims中的数字可能是float64）
		var userID uint
		if uid, ok := claims["user_id"].(float64); ok {
			userID = uint(uid)
		} else if uid, ok := claims["user_id"].(uint); ok {
			userID = uid
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   true,
				"message": "Invalid user_id in token",
			})
			c.Abort()
			return
		}

		c.Set("user_id", userID)
		if username, ok := claims["username"].(string); ok {
			c.Set("username", username)
		}
		c.Next()
	}
}

