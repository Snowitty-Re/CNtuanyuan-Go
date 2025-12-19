package services

import (
	"errors"
	"time"

	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/models"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/repositories"
)

type NoticeService struct {
	noticeRepo *repositories.NoticeRepository
}

func NewNoticeService(noticeRepo *repositories.NoticeRepository) *NoticeService {
	return &NoticeService{noticeRepo: noticeRepo}
}

// CreateNotice 创建通知
func (s *NoticeService) CreateNotice(notice *models.Notice) error {
	if notice.PublishTime.IsZero() {
		notice.PublishTime = time.Now()
	}
	return s.noticeRepo.Create(notice)
}

// MarkAsRead 标记通知为已读
func (s *NoticeService) MarkAsRead(noticeID, userID uint) error {
	return s.noticeRepo.MarkAsRead(noticeID, userID)
}

