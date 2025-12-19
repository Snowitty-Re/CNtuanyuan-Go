package handlers

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/models"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/repositories"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/services"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/utils"
	"gorm.io/gorm"
)

type NoticeHandler struct {
	noticeService *services.NoticeService
	db            *gorm.DB
}

func NewNoticeHandler(db *gorm.DB) *NoticeHandler {
	noticeRepo := repositories.NewNoticeRepository(db)
	noticeService := services.NewNoticeService(noticeRepo)
	return &NoticeHandler{
		noticeService: noticeService,
		db:            db,
	}
}

// CreateNotice 创建通知
func (h *NoticeHandler) CreateNotice(c *gin.Context) {
	var req models.Notice
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	userID, _ := c.Get("user_id")
	req.CreatedBy = userID.(uint)

	if err := h.noticeService.CreateNotice(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	utils.Success(c, req, "通知创建成功")
}

// GetNotice 获取通知详情
func (h *NoticeHandler) GetNotice(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的ID")
		return
	}

	noticeRepo := repositories.NewNoticeRepository(h.db)
	notice, err := noticeRepo.GetByID(uint(id))
	if err != nil {
		utils.NotFound(c, "通知不存在")
		return
	}

	utils.Success(c, notice, "获取成功")
}

// ListNotices 获取通知列表
func (h *NoticeHandler) ListNotices(c *gin.Context) {
	noticeRepo := repositories.NewNoticeRepository(h.db)

	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	filters := make(map[string]interface{})
	if status := c.Query("status"); status != "" {
		filters["status"] = status
	}
	if noticeType := c.Query("type"); noticeType != "" {
		filters["type"] = noticeType
	}
	if priority := c.Query("priority"); priority != "" {
		filters["priority"] = priority
	}

	notices, total, err := noticeRepo.List(offset, limit, filters)
	if err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	utils.Success(c, gin.H{
		"notices": notices,
		"total":   total,
	}, "获取成功")
}

// MarkAsRead 标记为已读
func (h *NoticeHandler) MarkAsRead(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的ID")
		return
	}

	userID, _ := c.Get("user_id")
	if err := h.noticeService.MarkAsRead(uint(id), userID.(uint)); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	utils.Success(c, nil, "标记成功")
}

// GetUnreadCount 获取未读通知数量
func (h *NoticeHandler) GetUnreadCount(c *gin.Context) {
	noticeRepo := repositories.NewNoticeRepository(h.db)
	userID, _ := c.Get("user_id")

	count, err := noticeRepo.GetUnreadCount(userID.(uint))
	if err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	utils.Success(c, gin.H{"count": count}, "获取成功")
}

