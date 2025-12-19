package repositories

import (
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/models"
	"gorm.io/gorm"
)

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{db: db}
}

// Create 创建消息
func (r *MessageRepository) Create(message *models.Message) error {
	return r.db.Create(message).Error
}

// GetByID 根据ID获取消息
func (r *MessageRepository) GetByID(id uint) (*models.Message, error) {
	var message models.Message
	err := r.db.Preload("FromUser").Preload("ToUser").First(&message, id).Error
	return &message, err
}

// Update 更新消息
func (r *MessageRepository) Update(message *models.Message) error {
	return r.db.Save(message).Error
}

// Delete 删除消息
func (r *MessageRepository) Delete(id uint) error {
	return r.db.Delete(&models.Message{}, id).Error
}

// ListByUser 获取用户的消息列表
func (r *MessageRepository) ListByUser(userID uint, offset, limit int, filters map[string]interface{}) ([]models.Message, int64, error) {
	var messages []models.Message
	var total int64

	query := r.db.Model(&models.Message{}).Where("to_user_id = ?", userID)

	// 应用筛选条件
	if status, ok := filters["status"].(string); ok && status != "" {
		query = query.Where("status = ?", status)
	}
	if messageType, ok := filters["type"].(string); ok && messageType != "" {
		query = query.Where("type = ?", messageType)
	}
	if isRead, ok := filters["is_read"].(bool); ok {
		query = query.Where("is_read = ?", isRead)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Offset(offset).Limit(limit).Preload("FromUser").Order("created_at DESC").Find(&messages).Error
	return messages, total, err
}

// MarkAsRead 标记为已读
func (r *MessageRepository) MarkAsRead(id uint) error {
	now := r.db.NowFunc()
	return r.db.Model(&models.Message{}).Where("id = ?", id).Updates(map[string]interface{}{
		"is_read": true,
		"read_at": now,
		"status":  "read",
	}).Error
}

// GetUnreadCount 获取未读消息数量
func (r *MessageRepository) GetUnreadCount(userID uint) (int64, error) {
	var count int64
	err := r.db.Model(&models.Message{}).
		Where("to_user_id = ? AND is_read = ?", userID, false).
		Count(&count).Error
	return count, err
}

