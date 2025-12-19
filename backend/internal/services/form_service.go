package services

import (
	"encoding/json"
	"errors"

	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/models"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/repositories"
)

type FormService struct {
	formRepo *repositories.FormRepository
}

func NewFormService(formRepo *repositories.FormRepository) *FormService {
	return &FormService{formRepo: formRepo}
}

// CreateForm 创建表单
func (s *FormService) CreateForm(name, description, formConfig string) (*models.Form, error) {
	// 验证formConfig是否为有效的JSON
	var config interface{}
	if err := json.Unmarshal([]byte(formConfig), &config); err != nil {
		return nil, errors.New("表单配置必须是有效的JSON格式")
	}

	form := &models.Form{
		Name:        name,
		Description: description,
		FormConfig:  formConfig,
	}

	if err := s.formRepo.Create(form); err != nil {
		return nil, err
	}

	return form, nil
}

// SubmitForm 提交表单
func (s *FormService) SubmitForm(formID uint, data string, submittedBy uint) (*models.FormSubmission, error) {
	// 验证表单是否存在
	_, err := s.formRepo.GetByID(formID)
	if err != nil {
		return nil, errors.New("表单不存在")
	}

	// 验证data是否为有效的JSON
	var formData interface{}
	if err := json.Unmarshal([]byte(data), &formData); err != nil {
		return nil, errors.New("提交数据必须是有效的JSON格式")
	}

	// 这里可以添加更多的验证逻辑，比如根据formConfig验证字段

	submission := &models.FormSubmission{
		FormID:      formID,
		Data:         data,
		SubmittedBy: submittedBy,
	}

	if err := s.formRepo.CreateSubmission(submission); err != nil {
		return nil, err
	}

	return submission, nil
}

