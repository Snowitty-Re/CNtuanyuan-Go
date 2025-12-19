package services

import (
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/models"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/repositories"
)

type MessageService struct {
	messageRepo *repositories.MessageRepository
}

func NewMessageService(messageRepo *repositories.MessageRepository) *MessageService {
	return &MessageService{messageRepo: messageRepo}
}

// SendMessage 发送消息
func (s *MessageService) SendMessage(message *models.Message) error {
	return s.messageRepo.Create(message)
}

// MarkAsRead 标记消息为已读
func (s *MessageService) MarkAsRead(messageID uint) error {
	return s.messageRepo.MarkAsRead(messageID)
}

