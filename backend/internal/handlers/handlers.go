package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/config"
	"gorm.io/gorm"
)

// InitHandlers 初始化所有处理器
func InitHandlers(r *gin.Engine, db *gorm.DB, cfg *config.Config) {
	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"message": "Server is running",
		})
	})

	// API信息
	r.GET("/api/v1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name":    "志愿者OA系统 API",
			"version": "1.0.0",
		})
	})

	// API路由组
	api := r.Group("/api/v1")
	{
		// 认证相关路由将在后续阶段添加
		// api.POST("/auth/login", Login)
		// api.POST("/auth/register", Register)
	}
}

