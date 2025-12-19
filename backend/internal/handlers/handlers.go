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

		// OA基础功能路由（需要认证）
		// 通知公告
		noticeHandler := NewNoticeHandler(db)
		notices := api.Group("/notices", middleware.Auth(), middleware.Audit(db))
		{
			notices.POST("", noticeHandler.CreateNotice)
			notices.GET("", noticeHandler.ListNotices)
			notices.GET("/:id", noticeHandler.GetNotice)
			notices.POST("/:id/read", noticeHandler.MarkAsRead)
			notices.GET("/unread/count", noticeHandler.GetUnreadCount)
		}

		// 消息中心
		messageHandler := NewMessageHandler(db)
		messages := api.Group("/messages", middleware.Auth(), middleware.Audit(db))
		{
			messages.POST("", messageHandler.SendMessage)
			messages.GET("", messageHandler.ListMessages)
			messages.GET("/:id", messageHandler.GetMessage)
			messages.POST("/:id/read", messageHandler.MarkAsRead)
			messages.GET("/unread/count", messageHandler.GetUnreadCount)
		}

		// 文件管理
		fileHandler := NewFileHandler(db, "./uploads/files")
		files := api.Group("/files", middleware.Auth(), middleware.Audit(db))
		{
			files.POST("/upload", fileHandler.UploadFile)
			files.GET("", fileHandler.ListFiles)
			files.GET("/:id", fileHandler.GetFile)
			files.DELETE("/:id", fileHandler.DeleteFile)
		}

		// 任务管理
		taskHandler := NewTaskHandler(db)
		tasks := api.Group("/tasks", middleware.Auth(), middleware.Audit(db))
		{
			tasks.POST("", taskHandler.CreateTask)
			tasks.GET("", taskHandler.ListTasks)
			tasks.GET("/:id", taskHandler.GetTask)
			tasks.PUT("/:id", taskHandler.UpdateTask)
			tasks.DELETE("/:id", taskHandler.DeleteTask)
		}

		// 审计日志
		auditLogHandler := NewAuditLogHandler(db)
		logs := api.Group("/audit-logs", middleware.Auth())
		{
			logs.GET("", auditLogHandler.ListLogs)
		}
	}
}

