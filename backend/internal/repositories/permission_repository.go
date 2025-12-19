package repositories

import (
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/models"
	"gorm.io/gorm"
)

type PermissionRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) *PermissionRepository {
	return &PermissionRepository{db: db}
}

// Create 创建权限
func (r *PermissionRepository) Create(permission *models.Permission) error {
	return r.db.Create(permission).Error
}

// GetByID 根据ID获取权限
func (r *PermissionRepository) GetByID(id uint) (*models.Permission, error) {
	var permission models.Permission
	err := r.db.First(&permission, id).Error
	return &permission, err
}

// GetByResourceAndAction 根据资源和操作获取权限
func (r *PermissionRepository) GetByResourceAndAction(resource, action string) (*models.Permission, error) {
	var permission models.Permission
	err := r.db.Where("resource = ? AND action = ?", resource, action).First(&permission).Error
	return &permission, err
}

// Update 更新权限
func (r *PermissionRepository) Update(permission *models.Permission) error {
	return r.db.Save(permission).Error
}

// Delete 删除权限
func (r *PermissionRepository) Delete(id uint) error {
	return r.db.Delete(&models.Permission{}, id).Error
}

// List 获取权限列表
func (r *PermissionRepository) List(offset, limit int) ([]models.Permission, int64, error) {
	var permissions []models.Permission
	var total int64

	if err := r.db.Model(&models.Permission{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.Offset(offset).Limit(limit).Find(&permissions).Error
	return permissions, total, err
}
