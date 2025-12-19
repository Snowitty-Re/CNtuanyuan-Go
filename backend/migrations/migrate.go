package migrations

import (
	"gorm.io/gorm"
)

// RunMigrations 运行所有数据库迁移
func RunMigrations(db *gorm.DB) error {
	// 迁移将在后续阶段添加具体的模型
	// 这里先创建一个占位函数
	return nil
}

