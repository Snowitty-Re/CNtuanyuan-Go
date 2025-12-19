package models

import (
	"time"

	"gorm.io/gorm"
)

// MissingPerson 走失人员模型
type MissingPerson struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Name            string    `gorm:"not null" json:"name"`
	Gender          string    `json:"gender"` // male, female, unknown
	Age             int       `json:"age"`
	PhotoURL        string    `json:"photo_url"`
	MissingTime     time.Time `json:"missing_time"`
	MissingLocation string    `json:"missing_location"`
	Description     string    `gorm:"type:text" json:"description"`
	Height          float64   `json:"height"` // 身高（cm）
	Weight          float64   `json:"weight"` // 体重（kg）
	Contact         string    `json:"contact"` // 联系方式
	Status          string    `gorm:"default:missing" json:"status"` // missing, found, confirmed

	// 关联关系
	WorkflowInstanceID *uint            `json:"workflow_instance_id,omitempty"`
	WorkflowInstance   *WorkflowInstance `gorm:"foreignKey:WorkflowInstanceID" json:"workflow_instance,omitempty"`
	CreatedBy          uint              `gorm:"not null" json:"created_by"`
	Creator            User              `gorm:"foreignKey:CreatedBy" json:"creator,omitempty"`
}

// TableName 指定表名
func (MissingPerson) TableName() string {
	return "missing_persons"
}

