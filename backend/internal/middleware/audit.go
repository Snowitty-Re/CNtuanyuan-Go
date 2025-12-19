package middleware

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/models"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/repositories"
	"gorm.io/gorm"
)

// Audit 审计日志中间件
func Audit(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// 只记录需要认证的请求
		userIDInterface, exists := c.Get("user_id")
		if !exists {
			return
		}

		var userID uint
		switch v := userIDInterface.(type) {
		case uint:
			userID = v
		case float64:
			userID = uint(v)
		case int:
			userID = uint(v)
		default:
			return
		}

		usernameInterface, _ := c.Get("username")
		username := ""
		if u, ok := usernameInterface.(string); ok {
			username = u
		}

		// 获取操作信息
		action := getActionFromMethod(c.Request.Method)
		resource := getResourceFromPath(c.Request.URL.Path)

		// 创建审计日志
		details := map[string]interface{}{
			"method": c.Request.Method,
			"path":   c.Request.URL.Path,
		}

		detailsJSON, _ := json.Marshal(details)

		logRepo := repositories.NewAuditLogRepository(db)
		auditLog := &models.AuditLog{
			UserID:    userID,
			Username:  username,
			Action:    action,
			Resource:  resource,
			IP:        c.ClientIP(),
			UserAgent: c.Request.UserAgent(),
			Details:   string(detailsJSON),
		}

		logRepo.Create(auditLog)
	}
}

// getActionFromMethod 根据HTTP方法获取操作类型
func getActionFromMethod(method string) string {
	switch method {
	case "GET":
		return "view"
	case "POST":
		return "create"
	case "PUT", "PATCH":
		return "update"
	case "DELETE":
		return "delete"
	default:
		return "other"
	}
}

// getResourceFromPath 从路径提取资源类型
func getResourceFromPath(path string) string {
	// 简化处理，从路径中提取资源名
	// 例如: /api/v1/notices -> notices
	if len(path) > 9 && path[:9] == "/api/v1/" {
		parts := path[9:]
		for i, char := range parts {
			if char == '/' {
				return parts[:i]
			}
		}
		return parts
	}
	return "unknown"
}

