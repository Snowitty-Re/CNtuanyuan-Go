package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/repositories"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/services"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/utils"
	"gorm.io/gorm"
)

type PermissionHandler struct {
	permissionService *services.PermissionService
	db                *gorm.DB
}

func NewPermissionHandler(db *gorm.DB) *PermissionHandler {
	roleRepo := repositories.NewRoleRepository(db)
	permissionRepo := repositories.NewPermissionRepository(db)
	userRepo := repositories.NewUserRepository(db)
	permissionService := services.NewPermissionService(roleRepo, permissionRepo, userRepo)
	return &PermissionHandler{
		permissionService: permissionService,
		db:                db,
	}
}

// CreateRole 创建角色
func (h *PermissionHandler) CreateRole(c *gin.Context) {
	var req struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	role, err := h.permissionService.CreateRole(req.Name, req.Description)
	if err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	utils.Success(c, role, "角色创建成功")
}

// CreatePermission 创建权限
func (h *PermissionHandler) CreatePermission(c *gin.Context) {
	var req struct {
		Name     string `json:"name" binding:"required"`
		Resource string `json:"resource" binding:"required"`
		Action   string `json:"action" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	permission, err := h.permissionService.CreatePermission(req.Name, req.Resource, req.Action)
	if err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	utils.Success(c, permission, "权限创建成功")
}

// AssignRoleToUser 为用户分配角色
func (h *PermissionHandler) AssignRoleToUser(c *gin.Context) {
	var req struct {
		UserID uint `json:"user_id" binding:"required"`
		RoleID uint `json:"role_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	if err := h.permissionService.AssignRoleToUser(req.UserID, req.RoleID); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	utils.Success(c, nil, "角色分配成功")
}

// AssignPermissionsToRole 为角色分配权限
func (h *PermissionHandler) AssignPermissionsToRole(c *gin.Context) {
	var req struct {
		RoleID        uint   `json:"role_id" binding:"required"`
		PermissionIDs []uint `json:"permission_ids" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	if err := h.permissionService.AssignPermissionsToRole(req.RoleID, req.PermissionIDs); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	utils.Success(c, nil, "权限分配成功")
}

// ListRoles 获取角色列表
func (h *PermissionHandler) ListRoles(c *gin.Context) {
	roleRepo := repositories.NewRoleRepository(h.db)
	offset := 0
	limit := 100

	roles, total, err := roleRepo.List(offset, limit)
	if err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	utils.Success(c, gin.H{
		"roles": roles,
		"total": total,
	}, "获取成功")
}

// ListPermissions 获取权限列表
func (h *PermissionHandler) ListPermissions(c *gin.Context) {
	permissionRepo := repositories.NewPermissionRepository(h.db)
	offset := 0
	limit := 100

	permissions, total, err := permissionRepo.List(offset, limit)
	if err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	utils.Success(c, gin.H{
		"permissions": permissions,
		"total":      total,
	}, "获取成功")
}

