package migrations

import (
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/models"
	"gorm.io/gorm"
)

// RunMigrations 运行所有数据库迁移
func RunMigrations(db *gorm.DB) error {
	// 自动迁移所有模型
	err := db.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Permission{},
		&models.Workflow{},
		&models.WorkflowField{},
		&models.WorkflowState{},
		&models.WorkflowInstance{},
		&models.MissingPerson{},
		&models.Form{},
		&models.FormSubmission{},
		&models.Notice{},
		&models.NoticeRead{},
		&models.Message{},
		&models.File{},
		&models.Task{},
		&models.AuditLog{},
	)
	return err
}

