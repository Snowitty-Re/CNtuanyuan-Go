package repositories

import (
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/models"
	"gorm.io/gorm"
)

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{db: db}
}

// Create 创建角色
func (r *RoleRepository) Create(role *models.Role) error {
	return r.db.Create(role).Error
}

// GetByID 根据ID获取角色
func (r *RoleRepository) GetByID(id uint) (*models.Role, error) {
	var role models.Role
	err := r.db.Preload("Permissions").First(&role, id).Error
	return &role, err
}

// GetByName 根据名称获取角色
func (r *RoleRepository) GetByName(name string) (*models.Role, error) {
	var role models.Role
	err := r.db.Where("name = ?", name).Preload("Permissions").First(&role).Error
	return &role, err
}

// Update 更新角色
func (r *RoleRepository) Update(role *models.Role) error {
	return r.db.Save(role).Error
}

// Delete 删除角色
func (r *RoleRepository) Delete(id uint) error {
	return r.db.Delete(&models.Role{}, id).Error
}

// List 获取角色列表
func (r *RoleRepository) List(offset, limit int) ([]models.Role, int64, error) {
	var roles []models.Role
	var total int64

	if err := r.db.Model(&models.Role{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.Offset(offset).Limit(limit).Preload("Permissions").Find(&roles).Error
	return roles, total, err
}

// AssignPermissions 为角色分配权限
func (r *RoleRepository) AssignPermissions(roleID uint, permissionIDs []uint) error {
	role, err := r.GetByID(roleID)
	if err != nil {
		return err
	}

	var permissions []models.Permission
	if err := r.db.Where("id IN ?", permissionIDs).Find(&permissions).Error; err != nil {
		return err
	}

	return r.db.Model(role).Association("Permissions").Replace(permissions)
}

