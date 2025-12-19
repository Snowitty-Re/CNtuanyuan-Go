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

type MissingPersonHandler struct {
	personService *services.MissingPersonService
	db            *gorm.DB
}

func NewMissingPersonHandler(db *gorm.DB, uploadDir string) *MissingPersonHandler {
	personRepo := repositories.NewMissingPersonRepository(db)
	personService := services.NewMissingPersonService(personRepo, uploadDir)
	return &MissingPersonHandler{
		personService: personService,
		db:            db,
	}
}

// CreateMissingPerson 创建走失人员记录
func (h *MissingPersonHandler) CreateMissingPerson(c *gin.Context) {
	var req models.MissingPerson
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	userID, _ := c.Get("user_id")
	req.CreatedBy = userID.(uint)

	if err := h.personService.CreateMissingPerson(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	utils.Success(c, req, "创建成功")
}

// GetMissingPerson 获取走失人员详情
func (h *MissingPersonHandler) GetMissingPerson(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的ID")
		return
	}

	personRepo := repositories.NewMissingPersonRepository(h.db)
	person, err := personRepo.GetByID(uint(id))
	if err != nil {
		utils.NotFound(c, "记录不存在")
		return
	}

	utils.Success(c, person, "获取成功")
}

// UpdateMissingPerson 更新走失人员记录
func (h *MissingPersonHandler) UpdateMissingPerson(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的ID")
		return
	}

	var req models.MissingPerson
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	if err := h.personService.UpdateMissingPerson(uint(id), &req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	utils.Success(c, nil, "更新成功")
}

// DeleteMissingPerson 删除走失人员记录
func (h *MissingPersonHandler) DeleteMissingPerson(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的ID")
		return
	}

	personRepo := repositories.NewMissingPersonRepository(h.db)
	if err := personRepo.Delete(uint(id)); err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	utils.Success(c, nil, "删除成功")
}

// ListMissingPersons 获取走失人员列表
func (h *MissingPersonHandler) ListMissingPersons(c *gin.Context) {
	personRepo := repositories.NewMissingPersonRepository(h.db)

	// 解析分页参数
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	// 解析筛选条件
	filters := make(map[string]interface{})
	if name := c.Query("name"); name != "" {
		filters["name"] = name
	}
	if gender := c.Query("gender"); gender != "" {
		filters["gender"] = gender
	}
	if status := c.Query("status"); status != "" {
		filters["status"] = status
	}
	if location := c.Query("location"); location != "" {
		filters["location"] = location
	}

	persons, total, err := personRepo.List(offset, limit, filters)
	if err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	utils.Success(c, gin.H{
		"persons": persons,
		"total":   total,
		"offset":  offset,
		"limit":   limit,
	}, "获取成功")
}

// SearchMissingPersons 搜索走失人员
func (h *MissingPersonHandler) SearchMissingPersons(c *gin.Context) {
	personRepo := repositories.NewMissingPersonRepository(h.db)

	keyword := c.Query("keyword")
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	persons, total, err := personRepo.Search(keyword, offset, limit)
	if err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	utils.Success(c, gin.H{
		"persons": persons,
		"total":   total,
		"offset":  offset,
		"limit":   limit,
	}, "搜索成功")
}

// UploadPhoto 上传照片
func (h *MissingPersonHandler) UploadPhoto(c *gin.Context) {
	file, err := c.FormFile("photo")
	if err != nil {
		utils.BadRequest(c, "请选择文件")
		return
	}

	photoURL, err := h.personService.UploadPhoto(file)
	if err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	utils.Success(c, gin.H{"photo_url": photoURL}, "上传成功")
}

