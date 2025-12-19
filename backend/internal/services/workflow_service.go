package services

import (
	"errors"

	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/models"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/repositories"
)

type WorkflowService struct {
	workflowRepo *repositories.WorkflowRepository
}

func NewWorkflowService(workflowRepo *repositories.WorkflowRepository) *WorkflowService {
	return &WorkflowService{workflowRepo: workflowRepo}
}

// CreateWorkflow 创建工作流
func (s *WorkflowService) CreateWorkflow(name, description string) (*models.Workflow, error) {
	workflow := &models.Workflow{
		Name:        name,
		Description: description,
		Status:      "active",
	}

	if err := s.workflowRepo.Create(workflow); err != nil {
		return nil, err
	}

	return workflow, nil
}

// AddField 添加字段
func (s *WorkflowService) AddField(workflowID uint, fieldName, fieldType, fieldConfig string, order int, required bool) (*models.WorkflowField, error) {
	// 验证工作流是否存在
	_, err := s.workflowRepo.GetByID(workflowID)
	if err != nil {
		return nil, errors.New("工作流不存在")
	}

	field := &models.WorkflowField{
		WorkflowID:  workflowID,
		FieldName:   fieldName,
		FieldType:   fieldType,
		FieldConfig: fieldConfig,
		Order:       order,
		Required:    required,
	}

	if err := s.workflowRepo.CreateField(field); err != nil {
		return nil, err
	}

	return field, nil
}

// AddState 添加状态
func (s *WorkflowService) AddState(workflowID uint, stateName, stateType string, order int) (*models.WorkflowState, error) {
	// 验证工作流是否存在
	_, err := s.workflowRepo.GetByID(workflowID)
	if err != nil {
		return nil, errors.New("工作流不存在")
	}

	state := &models.WorkflowState{
		WorkflowID: workflowID,
		StateName:  stateName,
		StateType:  stateType,
		Order:      order,
	}

	if err := s.workflowRepo.CreateState(state); err != nil {
		return nil, err
	}

	return state, nil
}

// CreateInstance 创建实例
func (s *WorkflowService) CreateInstance(workflowID uint, data string, createdBy uint) (*models.WorkflowInstance, error) {
	// 获取工作流
	workflow, err := s.workflowRepo.GetByID(workflowID)
	if err != nil {
		return nil, errors.New("工作流不存在")
	}

	// 获取初始状态
	var initialState string
	for _, state := range workflow.States {
		if state.StateType == "initial" {
			initialState = state.StateName
			break
		}
	}

	if initialState == "" {
		return nil, errors.New("工作流没有初始状态")
	}

	instance := &models.WorkflowInstance{
		WorkflowID:   workflowID,
		CurrentState: initialState,
		Data:         data,
		CreatedBy:    createdBy,
	}

	if err := s.workflowRepo.CreateInstance(instance); err != nil {
		return nil, err
	}

	return instance, nil
}

// UpdateInstanceState 更新实例状态
func (s *WorkflowService) UpdateInstanceState(instanceID uint, newState string) error {
	instance, err := s.workflowRepo.GetInstanceByID(instanceID)
	if err != nil {
		return errors.New("实例不存在")
	}

	// 验证新状态是否在工作流中
	workflow, err := s.workflowRepo.GetByID(instance.WorkflowID)
	if err != nil {
		return errors.New("工作流不存在")
	}

	validState := false
	for _, state := range workflow.States {
		if state.StateName == newState {
			validState = true
			break
		}
	}

	if !validState {
		return errors.New("无效的状态")
	}

	instance.CurrentState = newState
	return s.workflowRepo.UpdateInstance(instance)
}

