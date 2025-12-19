package repositories

import (
	"strings"

	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/models"
	"gorm.io/gorm"
)

type MissingPersonRepository struct {
	db *gorm.DB
}

func NewMissingPersonRepository(db *gorm.DB) *MissingPersonRepository {
	return &MissingPersonRepository{db: db}
}

// Create 创建走失人员记录
func (r *MissingPersonRepository) Create(person *models.MissingPerson) error {
	return r.db.Create(person).Error
}

// GetByID 根据ID获取走失人员
func (r *MissingPersonRepository) GetByID(id uint) (*models.MissingPerson, error) {
	var person models.MissingPerson
	err := r.db.Preload("Creator").Preload("WorkflowInstance").First(&person, id).Error
	return &person, err
}

// Update 更新走失人员
func (r *MissingPersonRepository) Update(person *models.MissingPerson) error {
	return r.db.Save(person).Error
}

// Delete 删除走失人员
func (r *MissingPersonRepository) Delete(id uint) error {
	return r.db.Delete(&models.MissingPerson{}, id).Error
}

// List 获取走失人员列表
func (r *MissingPersonRepository) List(offset, limit int, filters map[string]interface{}) ([]models.MissingPerson, int64, error) {
	var persons []models.MissingPerson
	var total int64

	query := r.db.Model(&models.MissingPerson{})

	// 应用筛选条件
	if name, ok := filters["name"].(string); ok && name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if gender, ok := filters["gender"].(string); ok && gender != "" {
		query = query.Where("gender = ?", gender)
	}
	if status, ok := filters["status"].(string); ok && status != "" {
		query = query.Where("status = ?", status)
	}
	if location, ok := filters["location"].(string); ok && location != "" {
		query = query.Where("missing_location LIKE ?", "%"+location+"%")
	}

	// 计算总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 获取列表
	err := query.Offset(offset).Limit(limit).Preload("Creator").Order("created_at DESC").Find(&persons).Error
	return persons, total, err
}

// Search 搜索走失人员
func (r *MissingPersonRepository) Search(keyword string, offset, limit int) ([]models.MissingPerson, int64, error) {
	var persons []models.MissingPerson
	var total int64

	keyword = strings.TrimSpace(keyword)
	if keyword == "" {
		return r.List(offset, limit, map[string]interface{}{})
	}

	query := r.db.Model(&models.MissingPerson{}).Where(
		"name LIKE ? OR description LIKE ? OR missing_location LIKE ?",
		"%"+keyword+"%",
		"%"+keyword+"%",
		"%"+keyword+"%",
	)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Offset(offset).Limit(limit).Preload("Creator").Order("created_at DESC").Find(&persons).Error
	return persons, total, err
}

