package models

import (
	"time"

	"gorm.io/gorm"
)

// Form 表单模型
type Form struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Name        string `gorm:"not null" json:"name"`
	Description string `json:"description"`
	FormConfig  string `gorm:"type:text" json:"form_config"` // JSON配置：字段定义、验证规则等

	// 关联关系
	Submissions []FormSubmission `gorm:"foreignKey:FormID" json:"submissions,omitempty"`
}

// TableName 指定表名
func (Form) TableName() string {
	return "forms"
}

// FormSubmission 表单提交模型
type FormSubmission struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	FormID     uint   `gorm:"not null" json:"form_id"`
	Data       string `gorm:"type:text" json:"data"` // JSON数据：存储提交的表单数据
	SubmittedBy uint `gorm:"not null" json:"submitted_by"`

	// 关联关系
	Form      Form `gorm:"foreignKey:FormID" json:"form,omitempty"`
	Submitter User `gorm:"foreignKey:SubmittedBy" json:"submitter,omitempty"`
}

// TableName 指定表名
func (FormSubmission) TableName() string {
	return "form_submissions"
}

