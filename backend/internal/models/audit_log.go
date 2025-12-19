package models

import (
	"time"

	"gorm.io/gorm"
)

// AuditLog 审计日志模型
type AuditLog struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	UserID    uint   `gorm:"not null" json:"user_id"`
	Username  string `gorm:"not null" json:"username"`
	Action    string `gorm:"not null" json:"action"` // create, update, delete, view, login, logout
	Resource  string `gorm:"not null" json:"resource"` // user, notice, file, task, etc.
	ResourceID *uint `json:"resource_id,omitempty"`
	IP        string `json:"ip"`
	UserAgent string `gorm:"type:text" json:"user_agent"`
	Details   string `gorm:"type:text" json:"details"` // JSON格式的详细信息

	// 关联关系
	User User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

// TableName 指定表名
func (AuditLog) TableName() string {
	return "audit_logs"
}

