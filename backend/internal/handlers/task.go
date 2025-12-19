package handlers

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/models"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/repositories"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/utils"
	"gorm.io/gorm"
)

type TaskHandler struct {
	taskRepo *repositories.TaskRepository
	db       *gorm.DB
}

func NewTaskHandler(db *gorm.DB) *TaskHandler {
	return &TaskHandler{
		taskRepo: repositories.NewTaskRepository(db),
		db:       db,
	}
}

// CreateTask 创建任务
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var req models.Task
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	userID, _ := c.Get("user_id")
	req.CreatedBy = userID.(uint)

	if err := h.taskRepo.Create(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	utils.Success(c, req, "任务创建成功")
}

// GetTask 获取任务详情
func (h *TaskHandler) GetTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的ID")
		return
	}

	task, err := h.taskRepo.GetByID(uint(id))
	if err != nil {
		utils.NotFound(c, "任务不存在")
		return
	}

	utils.Success(c, task, "获取成功")
}

// UpdateTask 更新任务
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的ID")
		return
	}

	var req models.Task
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	req.ID = uint(id)
	
	// 如果状态变为done，设置完成时间
	if req.Status == "done" && req.CompletedAt == nil {
		now := time.Now()
		req.CompletedAt = &now
		req.Progress = 100
	}

	if err := h.taskRepo.Update(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	utils.Success(c, req, "更新成功")
}

// ListTasks 获取任务列表
func (h *TaskHandler) ListTasks(c *gin.Context) {
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	filters := make(map[string]interface{})
	if status := c.Query("status"); status != "" {
		filters["status"] = status
	}
	if priority := c.Query("priority"); priority != "" {
		filters["priority"] = priority
	}
	if createdBy := c.Query("created_by"); createdBy != "" {
		if uid, err := strconv.ParseUint(createdBy, 10, 32); err == nil {
			filters["created_by"] = uint(uid)
		}
	}
	if assignedTo := c.Query("assigned_to"); assignedTo != "" {
		if uid, err := strconv.ParseUint(assignedTo, 10, 32); err == nil {
			filters["assigned_to"] = uint(uid)
		}
	}

	tasks, total, err := h.taskRepo.List(offset, limit, filters)
	if err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	utils.Success(c, gin.H{
		"tasks": tasks,
		"total": total,
	}, "获取成功")
}

// DeleteTask 删除任务
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的ID")
		return
	}

	if err := h.taskRepo.Delete(uint(id)); err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	utils.Success(c, nil, "删除成功")
}

