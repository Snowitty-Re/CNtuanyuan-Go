import request from './request'

// 创建角色
export const createRole = (data) => request.post('/permissions/roles', data)

// 获取角色列表
export const getRoles = () => request.get('/permissions/roles')

// 创建权限
export const createPermission = (data) => request.post('/permissions/permissions', data)

// 获取权限列表
export const getPermissions = () => request.get('/permissions/permissions')

// 为用户分配角色
export const assignRoleToUser = (data) => request.post('/permissions/assign-role', data)

// 为角色分配权限
export const assignPermissionsToRole = (data) => request.post('/permissions/assign-permissions', data)

