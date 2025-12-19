package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/models"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/repositories"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/services"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/utils"
	"gorm.io/gorm"
)

type MessageHandler struct {
	messageService *services.MessageService
	db              *gorm.DB
}

func NewMessageHandler(db *gorm.DB) *MessageHandler {
	messageRepo := repositories.NewMessageRepository(db)
	messageService := services.NewMessageService(messageRepo)
	return &MessageHandler{
		messageService: messageService,
		db:             db,
	}
}

// SendMessage 发送消息
func (h *MessageHandler) SendMessage(c *gin.Context) {
	var req models.Message
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	fromUserID, _ := c.Get("user_id")
	fromUID := fromUserID.(uint)
	req.FromUserID = &fromUID

	if err := h.messageService.SendMessage(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	utils.Success(c, req, "消息发送成功")
}

// GetMessage 获取消息详情
func (h *MessageHandler) GetMessage(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的ID")
		return
	}

	messageRepo := repositories.NewMessageRepository(h.db)
	message, err := messageRepo.GetByID(uint(id))
	if err != nil {
		utils.NotFound(c, "消息不存在")
		return
	}

	utils.Success(c, message, "获取成功")
}

// ListMessages 获取消息列表
func (h *MessageHandler) ListMessages(c *gin.Context) {
	messageRepo := repositories.NewMessageRepository(h.db)
	userID, _ := c.Get("user_id")

	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	filters := make(map[string]interface{})
	if status := c.Query("status"); status != "" {
		filters["status"] = status
	}
	if messageType := c.Query("type"); messageType != "" {
		filters["type"] = messageType
	}
	if isRead := c.Query("is_read"); isRead != "" {
		filters["is_read"] = isRead == "true"
	}

	messages, total, err := messageRepo.ListByUser(userID.(uint), offset, limit, filters)
	if err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	utils.Success(c, gin.H{
		"messages": messages,
		"total":   total,
	}, "获取成功")
}

// MarkAsRead 标记为已读
func (h *MessageHandler) MarkAsRead(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的ID")
		return
	}

	if err := h.messageService.MarkAsRead(uint(id)); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	utils.Success(c, nil, "标记成功")
}

// GetUnreadCount 获取未读消息数量
func (h *MessageHandler) GetUnreadCount(c *gin.Context) {
	messageRepo := repositories.NewMessageRepository(h.db)
	userID, _ := c.Get("user_id")

	count, err := messageRepo.GetUnreadCount(userID.(uint))
	if err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	utils.Success(c, gin.H{"count": count}, "获取成功")
}

