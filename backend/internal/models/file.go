package models

import (
	"time"

	"gorm.io/gorm"
)

// File 文件模型
type File struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	FileName     string `gorm:"not null" json:"file_name"`
	OriginalName string `gorm:"not null" json:"original_name"`
	FilePath     string `gorm:"not null" json:"file_path"`
	FileSize     int64  `gorm:"not null" json:"file_size"` // 字节
	FileType     string `json:"file_type"` // MIME类型
	FileExt      string `json:"file_ext"` // 文件扩展名
	Category     string `json:"category"` // 分类：document, image, video, other
	Description  string `gorm:"type:text" json:"description"`

	// 关联关系
	UploadedBy uint `gorm:"not null" json:"uploaded_by"`
	Uploader   User `gorm:"foreignKey:UploadedBy" json:"uploader,omitempty"`
}

// TableName 指定表名
func (File) TableName() string {
	return "files"
}

