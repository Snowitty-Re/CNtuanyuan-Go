package main

import (
	"log"

	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/config"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/handlers"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/middleware"
	"github.com/Snowitty-Re/CNtuanyuan-Go/migrations"
	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// 初始化数据库
	db, err := config.InitDB(cfg)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// 运行数据库迁移
	if err := migrations.RunMigrations(db); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	// 设置Gin模式
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 创建Gin路由
	r := gin.Default()

	// 应用中间件
	r.Use(middleware.CORS())
	r.Use(middleware.Logger())
	r.Use(middleware.ErrorHandler())

	// 静态文件服务（用于上传的文件）
	r.Static("/uploads", "./uploads")
	r.Static("/uploads/files", "./uploads/files")

	// 初始化处理器
	handlers.InitHandlers(r, db, cfg)

	// 启动服务器
	addr := ":" + cfg.Server.Port
	log.Printf("Server starting on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
