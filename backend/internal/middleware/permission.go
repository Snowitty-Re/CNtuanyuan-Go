package middleware

import (
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/repositories"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RequirePermission 权限验证中间件工厂函数
// 需要在handler中传入db，这里提供一个辅助函数
func RequirePermissionWithDB(db *gorm.DB, resource, action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			utils.Unauthorized(c, "未授权")
			c.Abort()
			return
		}

		// 获取用户
		userRepo := repositories.NewUserRepository(db)
		user, err := userRepo.GetByID(userID.(uint))
		if err != nil {
			utils.Unauthorized(c, "用户不存在")
			c.Abort()
			return
		}

		// 检查用户是否有权限
		hasPermission := false
		for _, role := range user.Roles {
			for _, permission := range role.Permissions {
				if permission.Resource == resource && permission.Action == action {
					hasPermission = true
					break
				}
			}
			if hasPermission {
				break
			}
		}

		if !hasPermission {
			utils.Forbidden(c, "权限不足")
			c.Abort()
			return
		}

		c.Next()
	}
}
