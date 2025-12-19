import request from './request'

// 创建表单
export const createForm = (data) => request.post('/forms', data)

// 获取表单列表
export const getForms = () => request.get('/forms')

// 获取表单详情
export const getForm = (id) => request.get(`/forms/${id}`)

// 提交表单
export const submitForm = (data) => request.post('/forms/submit', data)

// 获取提交列表
export const getSubmissions = (params) => request.get('/forms/submissions', { params })

