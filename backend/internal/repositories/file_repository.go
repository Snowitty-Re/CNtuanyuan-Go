package repositories

import (
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/models"
	"gorm.io/gorm"
)

type FileRepository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) *FileRepository {
	return &FileRepository{db: db}
}

// Create 创建文件记录
func (r *FileRepository) Create(file *models.File) error {
	return r.db.Create(file).Error
}

// GetByID 根据ID获取文件
func (r *FileRepository) GetByID(id uint) (*models.File, error) {
	var file models.File
	err := r.db.Preload("Uploader").First(&file, id).Error
	return &file, err
}

// Update 更新文件记录
func (r *FileRepository) Update(file *models.File) error {
	return r.db.Save(file).Error
}

// Delete 删除文件记录
func (r *FileRepository) Delete(id uint) error {
	return r.db.Delete(&models.File{}, id).Error
}

// List 获取文件列表
func (r *FileRepository) List(offset, limit int, filters map[string]interface{}) ([]models.File, int64, error) {
	var files []models.File
	var total int64

	query := r.db.Model(&models.File{})

	// 应用筛选条件
	if category, ok := filters["category"].(string); ok && category != "" {
		query = query.Where("category = ?", category)
	}
	if uploadedBy, ok := filters["uploaded_by"].(uint); ok && uploadedBy > 0 {
		query = query.Where("uploaded_by = ?", uploadedBy)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Offset(offset).Limit(limit).Preload("Uploader").Order("created_at DESC").Find(&files).Error
	return files, total, err
}

