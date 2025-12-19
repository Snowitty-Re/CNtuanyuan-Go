package handlers

import (
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/models"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/repositories"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/utils"
	"gorm.io/gorm"
)

type FileHandler struct {
	fileRepo *repositories.FileRepository
	uploadDir string
}

func NewFileHandler(db *gorm.DB, uploadDir string) *FileHandler {
	if uploadDir == "" {
		uploadDir = "./uploads/files"
	}
	os.MkdirAll(uploadDir, 0755)

	return &FileHandler{
		fileRepo: repositories.NewFileRepository(db),
		uploadDir: uploadDir,
	}
}

// UploadFile 上传文件
func (h *FileHandler) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		utils.BadRequest(c, "请选择文件")
		return
	}

	// 验证文件大小（最大50MB）
	if file.Size > 50*1024*1024 {
		utils.BadRequest(c, "文件大小不能超过50MB")
		return
	}

	// 生成文件名
	filename := time.Now().Format("20060102150405") + "_" + file.Filename
	filePath := filepath.Join(h.uploadDir, filename)

	// 保存文件
	src, err := file.Open()
	if err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}
	defer src.Close()

	dst, err := os.Create(filePath)
	if err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}
	defer dst.Close()

	if _, err := dst.ReadFrom(src); err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	// 获取文件信息
	ext := filepath.Ext(file.Filename)
	category := getFileCategory(ext)

	// 创建文件记录
	fileRecord := &models.File{
		FileName:     filename,
		OriginalName: file.Filename,
		FilePath:     "/uploads/files/" + filename,
		FileSize:     file.Size,
		FileType:     file.Header.Get("Content-Type"),
		FileExt:      ext,
		Category:     category,
	}

	userID, _ := c.Get("user_id")
	fileRecord.UploadedBy = userID.(uint)

	if err := h.fileRepo.Create(fileRecord); err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	utils.Success(c, fileRecord, "上传成功")
}

// GetFile 获取文件信息
func (h *FileHandler) GetFile(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的ID")
		return
	}

	file, err := h.fileRepo.GetByID(uint(id))
	if err != nil {
		utils.NotFound(c, "文件不存在")
		return
	}

	utils.Success(c, file, "获取成功")
}

// ListFiles 获取文件列表
func (h *FileHandler) ListFiles(c *gin.Context) {
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	filters := make(map[string]interface{})
	if category := c.Query("category"); category != "" {
		filters["category"] = category
	}
	if uploadedBy := c.Query("uploaded_by"); uploadedBy != "" {
		if uid, err := strconv.ParseUint(uploadedBy, 10, 32); err == nil {
			filters["uploaded_by"] = uint(uid)
		}
	}

	files, total, err := h.fileRepo.List(offset, limit, filters)
	if err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	utils.Success(c, gin.H{
		"files": files,
		"total": total,
	}, "获取成功")
}

// DeleteFile 删除文件
func (h *FileHandler) DeleteFile(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的ID")
		return
	}

	file, err := h.fileRepo.GetByID(uint(id))
	if err != nil {
		utils.NotFound(c, "文件不存在")
		return
	}

	// 删除物理文件
	if err := os.Remove(filepath.Join(h.uploadDir, file.FileName)); err != nil {
		// 文件不存在也继续删除记录
	}

	// 删除记录
	if err := h.fileRepo.Delete(uint(id)); err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	utils.Success(c, nil, "删除成功")
}

// getFileCategory 根据扩展名判断文件分类
func getFileCategory(ext string) string {
	// ext已经是扩展名（如.jpg），直接比较
	imageExts := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp"}
	videoExts := []string{".mp4", ".avi", ".mov", ".wmv", ".flv"}
	docExts := []string{".doc", ".docx", ".pdf", ".xls", ".xlsx", ".ppt", ".pptx", ".txt"}

	for _, e := range imageExts {
		if ext == e {
			return "image"
		}
	}
	for _, e := range videoExts {
		if ext == e {
			return "video"
		}
	}
	for _, e := range docExts {
		if ext == e {
			return "document"
		}
	}
	return "other"
}

