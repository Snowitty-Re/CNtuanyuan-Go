package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/config"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/middleware"
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
		// 认证相关路由
		authHandler := NewAuthHandler(db)
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.GET("/profile", middleware.Auth(), authHandler.GetProfile)
		}

		// 权限管理路由（需要认证）
		permissionHandler := NewPermissionHandler(db)
		permissions := api.Group("/permissions", middleware.Auth())
		{
			permissions.POST("/roles", permissionHandler.CreateRole)
			permissions.GET("/roles", permissionHandler.ListRoles)
			permissions.POST("/permissions", permissionHandler.CreatePermission)
			permissions.GET("/permissions", permissionHandler.ListPermissions)
			permissions.POST("/assign-role", permissionHandler.AssignRoleToUser)
			permissions.POST("/assign-permissions", permissionHandler.AssignPermissionsToRole)
		}

		// 工作流路由（需要认证）
		workflowHandler := NewWorkflowHandler(db)
		workflows := api.Group("/workflows", middleware.Auth())
		{
			workflows.POST("", workflowHandler.CreateWorkflow)
			workflows.GET("", workflowHandler.ListWorkflows)
			workflows.GET("/:id", workflowHandler.GetWorkflow)
			workflows.POST("/fields", workflowHandler.AddField)
			workflows.POST("/states", workflowHandler.AddState)
			workflows.POST("/instances", workflowHandler.CreateInstance)
			workflows.GET("/instances", workflowHandler.ListInstances)
			workflows.PUT("/instances/state", workflowHandler.UpdateInstanceState)
		}

		// 走失人员路由（需要认证）
		missingPersonHandler := NewMissingPersonHandler(db, "./uploads")
		missingPersons := api.Group("/missing-persons", middleware.Auth())
		{
			missingPersons.POST("", missingPersonHandler.CreateMissingPerson)
			missingPersons.GET("", missingPersonHandler.ListMissingPersons)
			missingPersons.GET("/search", missingPersonHandler.SearchMissingPersons)
			missingPersons.GET("/:id", missingPersonHandler.GetMissingPerson)
			missingPersons.PUT("/:id", missingPersonHandler.UpdateMissingPerson)
			missingPersons.DELETE("/:id", missingPersonHandler.DeleteMissingPerson)
			missingPersons.POST("/upload-photo", missingPersonHandler.UploadPhoto)
		}

		// 表单路由（需要认证）
		formHandler := NewFormHandler(db)
		forms := api.Group("/forms", middleware.Auth())
		{
			forms.POST("", formHandler.CreateForm)
			forms.GET("", formHandler.ListForms)
			forms.GET("/:id", formHandler.GetForm)
			forms.POST("/submit", formHandler.SubmitForm)
			forms.GET("/submissions", formHandler.ListSubmissions)
		}
	}
}

