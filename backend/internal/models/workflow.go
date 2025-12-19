package models

import (
	"time"

	"gorm.io/gorm"
)

// Workflow 工作流模型
type Workflow struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Name        string `gorm:"not null" json:"name"`
	Description string `json:"description"`
	Status      string `gorm:"default:active" json:"status"` // active, inactive

	// 关联关系
	Fields    []WorkflowField    `gorm:"foreignKey:WorkflowID" json:"fields,omitempty"`
	States    []WorkflowState    `gorm:"foreignKey:WorkflowID" json:"states,omitempty"`
	Instances []WorkflowInstance `gorm:"foreignKey:WorkflowID" json:"instances,omitempty"`
}

// TableName 指定表名
func (Workflow) TableName() string {
	return "workflows"
}

// WorkflowField 工作流字段模型
type WorkflowField struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	WorkflowID  uint   `gorm:"not null" json:"workflow_id"`
	FieldName   string `gorm:"not null" json:"field_name"`
	FieldType   string `gorm:"not null" json:"field_type"` // text, number, date, select, textarea
	FieldConfig string `gorm:"type:text" json:"field_config"` // JSON配置：选项、验证规则等
	Order       int    `gorm:"default:0" json:"order"`
	Required    bool   `gorm:"default:false" json:"required"`

	// 关联关系
	Workflow Workflow `gorm:"foreignKey:WorkflowID" json:"-"`
}

// TableName 指定表名
func (WorkflowField) TableName() string {
	return "workflow_fields"
}

// WorkflowState 工作流状态模型
type WorkflowState struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	WorkflowID uint   `gorm:"not null" json:"workflow_id"`
	StateName  string `gorm:"not null" json:"state_name"`
	StateType  string `gorm:"not null" json:"state_type"` // initial, intermediate, final
	Order      int    `gorm:"default:0" json:"order"`

	// 关联关系
	Workflow Workflow `gorm:"foreignKey:WorkflowID" json:"-"`
}

// TableName 指定表名
func (WorkflowState) TableName() string {
	return "workflow_states"
}

// WorkflowInstance 工作流实例模型
type WorkflowInstance struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	WorkflowID  uint   `gorm:"not null" json:"workflow_id"`
	CurrentState string `gorm:"not null" json:"current_state"`
	Data        string `gorm:"type:text" json:"data"` // JSON数据：存储字段值
	CreatedBy   uint   `gorm:"not null" json:"created_by"`

	// 关联关系
	Workflow  Workflow `gorm:"foreignKey:WorkflowID" json:"workflow,omitempty"`
	Creator   User     `gorm:"foreignKey:CreatedBy" json:"creator,omitempty"`
}

// TableName 指定表名
func (WorkflowInstance) TableName() string {
	return "workflow_instances"
}

