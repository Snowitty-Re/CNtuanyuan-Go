package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/repositories"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/services"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/utils"
	"gorm.io/gorm"
)

type WorkflowHandler struct {
	workflowService *services.WorkflowService
	db              *gorm.DB
}

func NewWorkflowHandler(db *gorm.DB) *WorkflowHandler {
	workflowRepo := repositories.NewWorkflowRepository(db)
	workflowService := services.NewWorkflowService(workflowRepo)
	return &WorkflowHandler{
		workflowService: workflowService,
		db:              db,
	}
}

// CreateWorkflow 创建工作流
func (h *WorkflowHandler) CreateWorkflow(c *gin.Context) {
	var req struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	workflow, err := h.workflowService.CreateWorkflow(req.Name, req.Description)
	if err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	utils.Success(c, workflow, "工作流创建成功")
}

// GetWorkflow 获取工作流
func (h *WorkflowHandler) GetWorkflow(c *gin.Context) {
	var id struct {
		ID uint `uri:"id" binding:"required"`
	}

	if err := c.ShouldBindUri(&id); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	workflowRepo := repositories.NewWorkflowRepository(h.db)
	workflow, err := workflowRepo.GetByID(id.ID)
	if err != nil {
		utils.NotFound(c, "工作流不存在")
		return
	}

	utils.Success(c, workflow, "获取成功")
}

// ListWorkflows 获取工作流列表
func (h *WorkflowHandler) ListWorkflows(c *gin.Context) {
	workflowRepo := repositories.NewWorkflowRepository(h.db)
	offset := 0
	limit := 100

	workflows, total, err := workflowRepo.List(offset, limit)
	if err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	utils.Success(c, gin.H{
		"workflows": workflows,
		"total":     total,
	}, "获取成功")
}

// AddField 添加字段
func (h *WorkflowHandler) AddField(c *gin.Context) {
	var req struct {
		WorkflowID  uint   `json:"workflow_id" binding:"required"`
		FieldName   string `json:"field_name" binding:"required"`
		FieldType   string `json:"field_type" binding:"required"`
		FieldConfig string `json:"field_config"`
		Order       int    `json:"order"`
		Required    bool   `json:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	field, err := h.workflowService.AddField(req.WorkflowID, req.FieldName, req.FieldType, req.FieldConfig, req.Order, req.Required)
	if err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	utils.Success(c, field, "字段添加成功")
}

// AddState 添加状态
func (h *WorkflowHandler) AddState(c *gin.Context) {
	var req struct {
		WorkflowID uint   `json:"workflow_id" binding:"required"`
		StateName  string `json:"state_name" binding:"required"`
		StateType  string `json:"state_type" binding:"required"`
		Order      int    `json:"order"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	state, err := h.workflowService.AddState(req.WorkflowID, req.StateName, req.StateType, req.Order)
	if err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	utils.Success(c, state, "状态添加成功")
}

// CreateInstance 创建实例
func (h *WorkflowHandler) CreateInstance(c *gin.Context) {
	var req struct {
		WorkflowID uint   `json:"workflow_id" binding:"required"`
		Data       string `json:"data" binding:"required"`
	}

	userID, _ := c.Get("user_id")

	instance, err := h.workflowService.CreateInstance(req.WorkflowID, req.Data, userID.(uint))
	if err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	utils.Success(c, instance, "实例创建成功")
}

// UpdateInstanceState 更新实例状态
func (h *WorkflowHandler) UpdateInstanceState(c *gin.Context) {
	var req struct {
		InstanceID uint   `json:"instance_id" binding:"required"`
		NewState   string `json:"new_state" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	if err := h.workflowService.UpdateInstanceState(req.InstanceID, req.NewState); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	utils.Success(c, nil, "状态更新成功")
}

// ListInstances 获取实例列表
func (h *WorkflowHandler) ListInstances(c *gin.Context) {
	workflowRepo := repositories.NewWorkflowRepository(h.db)
	workflowID := uint(0)
	offset := 0
	limit := 100

	if id := c.Query("workflow_id"); id != "" {
		// 解析workflow_id
	}

	instances, total, err := workflowRepo.ListInstances(workflowID, offset, limit)
	if err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	utils.Success(c, gin.H{
		"instances": instances,
		"total":     total,
	}, "获取成功")
}

