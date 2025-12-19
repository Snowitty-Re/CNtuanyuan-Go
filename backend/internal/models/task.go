package models

import (
	"time"

	"gorm.io/gorm"
)

// Task 任务模型
type Task struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Title       string     `gorm:"not null" json:"title"`
	Description string     `gorm:"type:text" json:"description"`
	Priority    string     `gorm:"default:normal" json:"priority"` // low, normal, high, urgent
	Status      string     `gorm:"default:todo" json:"status"` // todo, in_progress, done, cancelled
	Progress    int        `gorm:"default:0" json:"progress"` // 0-100
	DueDate     *time.Time `json:"due_date,omitempty"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`

	// 关联关系
	CreatedBy uint `gorm:"not null" json:"created_by"`
	Creator   User `gorm:"foreignKey:CreatedBy" json:"creator,omitempty"`
	AssignedTo *uint `json:"assigned_to,omitempty"`
	Assignee  *User `gorm:"foreignKey:AssignedTo" json:"assignee,omitempty"`
}

// TableName 指定表名
func (Task) TableName() string {
	return "tasks"
}

