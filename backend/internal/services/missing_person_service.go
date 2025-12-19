package services

import (
	"errors"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/models"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/repositories"
)

type MissingPersonService struct {
	personRepo *repositories.MissingPersonRepository
	uploadDir  string
}

func NewMissingPersonService(personRepo *repositories.MissingPersonRepository, uploadDir string) *MissingPersonService {
	// 确保上传目录存在
	if uploadDir == "" {
		uploadDir = "./uploads"
	}
	os.MkdirAll(uploadDir, 0755)

	return &MissingPersonService{
		personRepo: personRepo,
		uploadDir:  uploadDir,
	}
}

// CreateMissingPerson 创建走失人员记录
func (s *MissingPersonService) CreateMissingPerson(person *models.MissingPerson) error {
	return s.personRepo.Create(person)
}

// UpdateMissingPerson 更新走失人员记录
func (s *MissingPersonService) UpdateMissingPerson(id uint, person *models.MissingPerson) error {
	existing, err := s.personRepo.GetByID(id)
	if err != nil {
		return errors.New("走失人员记录不存在")
	}

	person.ID = existing.ID
	return s.personRepo.Update(person)
}

// UploadPhoto 上传照片
func (s *MissingPersonService) UploadPhoto(file *multipart.FileHeader) (string, error) {
	// 验证文件类型
	ext := filepath.Ext(file.Filename)
	allowedExts := []string{".jpg", ".jpeg", ".png", ".gif"}
	validExt := false
	for _, allowedExt := range allowedExts {
		if ext == allowedExt {
			validExt = true
			break
		}
	}

	if !validExt {
		return "", errors.New("不支持的文件类型，仅支持 jpg, jpeg, png, gif")
	}

	// 验证文件大小（最大5MB）
	if file.Size > 5*1024*1024 {
		return "", errors.New("文件大小不能超过5MB")
	}

	// 生成文件名
	filename := time.Now().Format("20060102150405") + "_" + file.Filename
	filepath := filepath.Join(s.uploadDir, filename)

	// 保存文件
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	dst, err := os.Create(filepath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := dst.ReadFrom(src); err != nil {
		return "", err
	}

	// 返回相对路径
	return "/uploads/" + filename, nil
}

