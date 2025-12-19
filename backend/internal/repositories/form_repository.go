package repositories

import (
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/models"
	"gorm.io/gorm"
)

type FormRepository struct {
	db *gorm.DB
}

func NewFormRepository(db *gorm.DB) *FormRepository {
	return &FormRepository{db: db}
}

// Create 创建表单
func (r *FormRepository) Create(form *models.Form) error {
	return r.db.Create(form).Error
}

// GetByID 根据ID获取表单
func (r *FormRepository) GetByID(id uint) (*models.Form, error) {
	var form models.Form
	err := r.db.First(&form, id).Error
	return &form, err
}

// Update 更新表单
func (r *FormRepository) Update(form *models.Form) error {
	return r.db.Save(form).Error
}

// Delete 删除表单
func (r *FormRepository) Delete(id uint) error {
	return r.db.Delete(&models.Form{}, id).Error
}

// List 获取表单列表
func (r *FormRepository) List(offset, limit int) ([]models.Form, int64, error) {
	var forms []models.Form
	var total int64

	if err := r.db.Model(&models.Form{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.Offset(offset).Limit(limit).Find(&forms).Error
	return forms, total, err
}

// CreateSubmission 创建提交
func (r *FormRepository) CreateSubmission(submission *models.FormSubmission) error {
	return r.db.Create(submission).Error
}

// GetSubmissionByID 根据ID获取提交
func (r *FormRepository) GetSubmissionByID(id uint) (*models.FormSubmission, error) {
	var submission models.FormSubmission
	err := r.db.Preload("Form").Preload("Submitter").First(&submission, id).Error
	return &submission, err
}

// ListSubmissions 获取提交列表
func (r *FormRepository) ListSubmissions(formID uint, offset, limit int) ([]models.FormSubmission, int64, error) {
	var submissions []models.FormSubmission
	var total int64

	query := r.db.Model(&models.FormSubmission{})
	if formID > 0 {
		query = query.Where("form_id = ?", formID)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Offset(offset).Limit(limit).Preload("Form").Preload("Submitter").Order("created_at DESC").Find(&submissions).Error
	return submissions, total, err
}

