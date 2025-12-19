package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/repositories"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/services"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/utils"
	"gorm.io/gorm"
)

type AuthHandler struct {
	authService *services.AuthService
	db          *gorm.DB
}

func NewAuthHandler(db *gorm.DB) *AuthHandler {
	userRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)
	return &AuthHandler{
		authService: authService,
		db:          db,
	}
}

// Register 用户注册
func (h *AuthHandler) Register(c *gin.Context) {
	var req services.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	user, token, err := h.authService.Register(&req)
	if err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	utils.Success(c, gin.H{
		"user":  user,
		"token": token,
	}, "注册成功")
}

// Login 用户登录
func (h *AuthHandler) Login(c *gin.Context) {
	var req services.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	user, token, err := h.authService.Login(&req)
	if err != nil {
		utils.Unauthorized(c, err.Error())
		return
	}

	utils.Success(c, gin.H{
		"user":  user,
		"token": token,
	}, "登录成功")
}

// GetProfile 获取当前用户信息
func (h *AuthHandler) GetProfile(c *gin.Context) {
	userIDInterface, exists := c.Get("user_id")
	if !exists {
		utils.Unauthorized(c, "未授权")
		return
	}

	// 安全地转换user_id
	var userID uint
	switch v := userIDInterface.(type) {
	case uint:
		userID = v
	case float64:
		userID = uint(v)
	case int:
		userID = uint(v)
	default:
		utils.Unauthorized(c, "无效的用户ID")
		return
	}

	userRepo := repositories.NewUserRepository(h.db)
	user, err := userRepo.GetByID(userID)
	if err != nil {
		utils.NotFound(c, "用户不存在")
		return
	}

	utils.Success(c, user, "获取成功")
}

