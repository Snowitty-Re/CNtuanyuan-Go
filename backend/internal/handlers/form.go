package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/repositories"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/services"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/utils"
	"gorm.io/gorm"
)

type FormHandler struct {
	formService *services.FormService
	db          *gorm.DB
}

func NewFormHandler(db *gorm.DB) *FormHandler {
	formRepo := repositories.NewFormRepository(db)
	formService := services.NewFormService(formRepo)
	return &FormHandler{
		formService: formService,
		db:          db,
	}
}

// CreateForm 创建表单
func (h *FormHandler) CreateForm(c *gin.Context) {
	var req struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
		FormConfig  string `json:"form_config" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	form, err := h.formService.CreateForm(req.Name, req.Description, req.FormConfig)
	if err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	utils.Success(c, form, "表单创建成功")
}

// GetForm 获取表单
func (h *FormHandler) GetForm(c *gin.Context) {
	var id struct {
		ID uint `uri:"id" binding:"required"`
	}

	if err := c.ShouldBindUri(&id); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	formRepo := repositories.NewFormRepository(h.db)
	form, err := formRepo.GetByID(id.ID)
	if err != nil {
		utils.NotFound(c, "表单不存在")
		return
	}

	utils.Success(c, form, "获取成功")
}

// ListForms 获取表单列表
func (h *FormHandler) ListForms(c *gin.Context) {
	formRepo := repositories.NewFormRepository(h.db)
	offset := 0
	limit := 100

	forms, total, err := formRepo.List(offset, limit)
	if err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	utils.Success(c, gin.H{
		"forms": forms,
		"total": total,
	}, "获取成功")
}

// SubmitForm 提交表单
func (h *FormHandler) SubmitForm(c *gin.Context) {
	var req struct {
		FormID uint   `json:"form_id" binding:"required"`
		Data   string `json:"data" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	userID, _ := c.Get("user_id")

	submission, err := h.formService.SubmitForm(req.FormID, req.Data, userID.(uint))
	if err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	utils.Success(c, submission, "提交成功")
}

// ListSubmissions 获取提交列表
func (h *FormHandler) ListSubmissions(c *gin.Context) {
	formRepo := repositories.NewFormRepository(h.db)
	formID := uint(0)
	offset := 0
	limit := 100

	if id := c.Query("form_id"); id != "" {
		// 解析form_id
	}

	submissions, total, err := formRepo.ListSubmissions(formID, offset, limit)
	if err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	utils.Success(c, gin.H{
		"submissions": submissions,
		"total":      total,
	}, "获取成功")
}

