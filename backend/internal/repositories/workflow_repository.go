package repositories

import (
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/models"
	"gorm.io/gorm"
)

type WorkflowRepository struct {
	db *gorm.DB
}

func NewWorkflowRepository(db *gorm.DB) *WorkflowRepository {
	return &WorkflowRepository{db: db}
}

// Create 创建工作流
func (r *WorkflowRepository) Create(workflow *models.Workflow) error {
	return r.db.Create(workflow).Error
}

// GetByID 根据ID获取工作流
func (r *WorkflowRepository) GetByID(id uint) (*models.Workflow, error) {
	var workflow models.Workflow
	err := r.db.Preload("Fields").Preload("States").First(&workflow, id).Error
	return &workflow, err
}

// Update 更新工作流
func (r *WorkflowRepository) Update(workflow *models.Workflow) error {
	return r.db.Save(workflow).Error
}

// Delete 删除工作流
func (r *WorkflowRepository) Delete(id uint) error {
	return r.db.Delete(&models.Workflow{}, id).Error
}

// List 获取工作流列表
func (r *WorkflowRepository) List(offset, limit int) ([]models.Workflow, int64, error) {
	var workflows []models.Workflow
	var total int64

	if err := r.db.Model(&models.Workflow{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.Offset(offset).Limit(limit).Preload("Fields").Preload("States").Find(&workflows).Error
	return workflows, total, err
}

// CreateField 创建字段
func (r *WorkflowRepository) CreateField(field *models.WorkflowField) error {
	return r.db.Create(field).Error
}

// UpdateField 更新字段
func (r *WorkflowRepository) UpdateField(field *models.WorkflowField) error {
	return r.db.Save(field).Error
}

// DeleteField 删除字段
func (r *WorkflowRepository) DeleteField(id uint) error {
	return r.db.Delete(&models.WorkflowField{}, id).Error
}

// CreateState 创建状态
func (r *WorkflowRepository) CreateState(state *models.WorkflowState) error {
	return r.db.Create(state).Error
}

// UpdateState 更新状态
func (r *WorkflowRepository) UpdateState(state *models.WorkflowState) error {
	return r.db.Save(state).Error
}

// DeleteState 删除状态
func (r *WorkflowRepository) DeleteState(id uint) error {
	return r.db.Delete(&models.WorkflowState{}, id).Error
}

// CreateInstance 创建实例
func (r *WorkflowRepository) CreateInstance(instance *models.WorkflowInstance) error {
	return r.db.Create(instance).Error
}

// GetInstanceByID 根据ID获取实例
func (r *WorkflowRepository) GetInstanceByID(id uint) (*models.WorkflowInstance, error) {
	var instance models.WorkflowInstance
	err := r.db.Preload("Workflow").Preload("Workflow.Fields").Preload("Creator").First(&instance, id).Error
	return &instance, err
}

// UpdateInstance 更新实例
func (r *WorkflowRepository) UpdateInstance(instance *models.WorkflowInstance) error {
	return r.db.Save(instance).Error
}

// ListInstances 获取实例列表
func (r *WorkflowRepository) ListInstances(workflowID uint, offset, limit int) ([]models.WorkflowInstance, int64, error) {
	var instances []models.WorkflowInstance
	var total int64

	query := r.db.Model(&models.WorkflowInstance{})
	if workflowID > 0 {
		query = query.Where("workflow_id = ?", workflowID)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Offset(offset).Limit(limit).Preload("Workflow").Preload("Creator").Find(&instances).Error
	return instances, total, err
}

