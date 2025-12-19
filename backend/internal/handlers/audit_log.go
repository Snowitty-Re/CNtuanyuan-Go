package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/repositories"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/utils"
	"gorm.io/gorm"
)

type AuditLogHandler struct {
	logRepo *repositories.AuditLogRepository
}

func NewAuditLogHandler(db *gorm.DB) *AuditLogHandler {
	return &AuditLogHandler{
		logRepo: repositories.NewAuditLogRepository(db),
	}
}

// ListLogs 获取审计日志列表
func (h *AuditLogHandler) ListLogs(c *gin.Context) {
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))

	filters := make(map[string]interface{})
	if userID := c.Query("user_id"); userID != "" {
		if uid, err := strconv.ParseUint(userID, 10, 32); err == nil {
			filters["user_id"] = uint(uid)
		}
	}
	if action := c.Query("action"); action != "" {
		filters["action"] = action
	}
	if resource := c.Query("resource"); resource != "" {
		filters["resource"] = resource
	}

	logs, total, err := h.logRepo.List(offset, limit, filters)
	if err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	utils.Success(c, gin.H{
		"logs":  logs,
		"total": total,
	}, "获取成功")
}

