import request from './request'

// 创建工作流
export const createWorkflow = (data) => request.post('/workflows', data)

// 获取工作流列表
export const getWorkflows = () => request.get('/workflows')

// 获取工作流详情
export const getWorkflow = (id) => request.get(`/workflows/${id}`)

// 添加字段
export const addField = (data) => request.post('/workflows/fields', data)

// 添加状态
export const addState = (data) => request.post('/workflows/states', data)

// 创建实例
export const createInstance = (data) => request.post('/workflows/instances', data)

// 获取实例列表
export const getInstances = (params) => request.get('/workflows/instances', { params })

// 更新实例状态
export const updateInstanceState = (data) => request.put('/workflows/instances/state', data)

