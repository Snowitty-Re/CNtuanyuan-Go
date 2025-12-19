package repositories

import (
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/models"
	"gorm.io/gorm"
)

type NoticeRepository struct {
	db *gorm.DB
}

func NewNoticeRepository(db *gorm.DB) *NoticeRepository {
	return &NoticeRepository{db: db}
}

// Create 创建通知
func (r *NoticeRepository) Create(notice *models.Notice) error {
	return r.db.Create(notice).Error
}

// GetByID 根据ID获取通知
func (r *NoticeRepository) GetByID(id uint) (*models.Notice, error) {
	var notice models.Notice
	err := r.db.Preload("Creator").First(&notice, id).Error
	return &notice, err
}

// Update 更新通知
func (r *NoticeRepository) Update(notice *models.Notice) error {
	return r.db.Save(notice).Error
}

// Delete 删除通知
func (r *NoticeRepository) Delete(id uint) error {
	return r.db.Delete(&models.Notice{}, id).Error
}

// List 获取通知列表
func (r *NoticeRepository) List(offset, limit int, filters map[string]interface{}) ([]models.Notice, int64, error) {
	var notices []models.Notice
	var total int64

	query := r.db.Model(&models.Notice{})

	// 应用筛选条件
	if status, ok := filters["status"].(string); ok && status != "" {
		query = query.Where("status = ?", status)
	}
	if noticeType, ok := filters["type"].(string); ok && noticeType != "" {
		query = query.Where("type = ?", noticeType)
	}
	if priority, ok := filters["priority"].(string); ok && priority != "" {
		query = query.Where("priority = ?", priority)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Offset(offset).Limit(limit).Preload("Creator").Order("created_at DESC").Find(&notices).Error
	return notices, total, err
}

// MarkAsRead 标记为已读
func (r *NoticeRepository) MarkAsRead(noticeID, userID uint) error {
	var noticeRead models.NoticeRead
	err := r.db.Where("notice_id = ? AND user_id = ?", noticeID, userID).First(&noticeRead).Error
	if err == nil {
		// 已存在，更新
		noticeRead.ReadAt = r.db.NowFunc()
		return r.db.Save(&noticeRead).Error
	}
	// 不存在，创建
	noticeRead = models.NoticeRead{
		NoticeID: noticeID,
		UserID:   userID,
		ReadAt:   r.db.NowFunc(),
	}
	return r.db.Create(&noticeRead).Error
}

// GetUnreadCount 获取未读通知数量
func (r *NoticeRepository) GetUnreadCount(userID uint) (int64, error) {
	var count int64
	err := r.db.Model(&models.Notice{}).
		Where("status = ?", "published").
		Where("id NOT IN (SELECT notice_id FROM notice_reads WHERE user_id = ?)", userID).
		Count(&count).Error
	return count, err
}

