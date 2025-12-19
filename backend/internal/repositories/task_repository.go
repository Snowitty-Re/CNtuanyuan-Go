package repositories

import (
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/models"
	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

// Create 创建任务
func (r *TaskRepository) Create(task *models.Task) error {
	return r.db.Create(task).Error
}

// GetByID 根据ID获取任务
func (r *TaskRepository) GetByID(id uint) (*models.Task, error) {
	var task models.Task
	err := r.db.Preload("Creator").Preload("Assignee").First(&task, id).Error
	return &task, err
}

// Update 更新任务
func (r *TaskRepository) Update(task *models.Task) error {
	return r.db.Save(task).Error
}

// Delete 删除任务
func (r *TaskRepository) Delete(id uint) error {
	return r.db.Delete(&models.Task{}, id).Error
}

// List 获取任务列表
func (r *TaskRepository) List(offset, limit int, filters map[string]interface{}) ([]models.Task, int64, error) {
	var tasks []models.Task
	var total int64

	query := r.db.Model(&models.Task{})

	// 应用筛选条件
	if status, ok := filters["status"].(string); ok && status != "" {
		query = query.Where("status = ?", status)
	}
	if priority, ok := filters["priority"].(string); ok && priority != "" {
		query = query.Where("priority = ?", priority)
	}
	if createdBy, ok := filters["created_by"].(uint); ok && createdBy > 0 {
		query = query.Where("created_by = ?", createdBy)
	}
	if assignedTo, ok := filters["assigned_to"].(uint); ok && assignedTo > 0 {
		query = query.Where("assigned_to = ?", assignedTo)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Offset(offset).Limit(limit).Preload("Creator").Preload("Assignee").Order("created_at DESC").Find(&tasks).Error
	return tasks, total, err
}

