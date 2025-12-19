package models

import (
	"time"

	"gorm.io/gorm"
)

// Notice 通知公告模型
type Notice struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Title       string    `gorm:"not null" json:"title"`
	Content     string    `gorm:"type:text" json:"content"`
	Type        string    `gorm:"default:notice" json:"type"` // notice, announcement, system
	Priority    string    `gorm:"default:normal" json:"priority"` // low, normal, high, urgent
	Status      string    `gorm:"default:published" json:"status"` // draft, published, archived
	PublishTime time.Time `json:"publish_time"`
	ExpireTime  *time.Time `json:"expire_time,omitempty"`

	// 关联关系
	CreatedBy uint `gorm:"not null" json:"created_by"`
	Creator   User `gorm:"foreignKey:CreatedBy" json:"creator,omitempty"`
}

// TableName 指定表名
func (Notice) TableName() string {
	return "notices"
}

// NoticeRead 通知已读记录
type NoticeRead struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	NoticeID uint `gorm:"not null;uniqueIndex:idx_notice_user" json:"notice_id"`
	UserID   uint `gorm:"not null;uniqueIndex:idx_notice_user" json:"user_id"`
	ReadAt   time.Time `json:"read_at"`

	// 关联关系
	Notice Notice `gorm:"foreignKey:NoticeID" json:"notice,omitempty"`
	User   User   `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

// TableName 指定表名
func (NoticeRead) TableName() string {
	return "notice_reads"
}

