package services

import (
	"errors"

	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/models"
	"github.com/Snowitty-Re/CNtuanyuan-Go/internal/repositories"
)

type PermissionService struct {
	roleRepo       *repositories.RoleRepository
	permissionRepo *repositories.PermissionRepository
	userRepo       *repositories.UserRepository
}

func NewPermissionService(
	roleRepo *repositories.RoleRepository,
	permissionRepo *repositories.PermissionRepository,
	userRepo *repositories.UserRepository,
) *PermissionService {
	return &PermissionService{
		roleRepo:       roleRepo,
		permissionRepo: permissionRepo,
		userRepo:       userRepo,
	}
}

// CreateRole 创建角色
func (s *PermissionService) CreateRole(name, description string) (*models.Role, error) {
	// 检查角色是否已存在
	_, err := s.roleRepo.GetByName(name)
	if err == nil {
		return nil, errors.New("角色已存在")
	}

	role := &models.Role{
		Name:        name,
		Description: description,
	}

	if err := s.roleRepo.Create(role); err != nil {
		return nil, err
	}

	return role, nil
}

// CreatePermission 创建权限
func (s *PermissionService) CreatePermission(name, resource, action string) (*models.Permission, error) {
	// 检查权限是否已存在
	_, err := s.permissionRepo.GetByResourceAndAction(resource, action)
	if err == nil {
		return nil, errors.New("权限已存在")
	}

	permission := &models.Permission{
		Name:     name,
		Resource: resource,
		Action:   action,
	}

	if err := s.permissionRepo.Create(permission); err != nil {
		return nil, err
	}

	return permission, nil
}

// AssignRoleToUser 为用户分配角色
func (s *PermissionService) AssignRoleToUser(userID uint, roleID uint) error {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return errors.New("用户不存在")
	}

	role, err := s.roleRepo.GetByID(roleID)
	if err != nil {
		return errors.New("角色不存在")
	}

	// 检查用户是否已有该角色
	for _, r := range user.Roles {
		if r.ID == roleID {
			return errors.New("用户已拥有该角色")
		}
	}

	// 添加角色
	user.Roles = append(user.Roles, *role)
	return s.userRepo.Update(user)
}

// AssignPermissionsToRole 为角色分配权限
func (s *PermissionService) AssignPermissionsToRole(roleID uint, permissionIDs []uint) error {
	return s.roleRepo.AssignPermissions(roleID, permissionIDs)
}

