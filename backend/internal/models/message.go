package models

import (
	"time"

	"gorm.io/gorm"
)

// Message 消息模型
type Message struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Title       string    `gorm:"not null" json:"title"`
	Content     string    `gorm:"type:text" json:"content"`
	Type        string    `gorm:"default:system" json:"type"` // system, workflow, task, custom
	Status      string    `gorm:"default:unread" json:"status"` // unread, read, archived
	IsRead      bool      `gorm:"default:false" json:"is_read"`
	ReadAt      *time.Time `json:"read_at,omitempty"`

	// 关联关系
	FromUserID *uint `json:"from_user_id,omitempty"`
	FromUser   *User `gorm:"foreignKey:FromUserID" json:"from_user,omitempty"`
	ToUserID   uint  `gorm:"not null" json:"to_user_id"`
	ToUser     User  `gorm:"foreignKey:ToUserID" json:"to_user,omitempty"`
}

// TableName 指定表名
func (Message) TableName() string {
	return "messages"
}

