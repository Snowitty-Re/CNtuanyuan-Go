import request from './request'

// 创建任务
export const createTask = (data) => request.post('/tasks', data)

// 获取任务列表
export const getTasks = (params) => request.get('/tasks', { params })

// 获取任务详情
export const getTask = (id) => request.get(`/tasks/${id}`)

// 更新任务
export const updateTask = (id, data) => request.put(`/tasks/${id}`, data)

// 删除任务
export const deleteTask = (id) => request.delete(`/tasks/${id}`)

