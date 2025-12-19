import request from './request'

// 健康检查
export const healthCheck = () => request.get('/health')

// 认证相关API将在后续阶段添加
// export const login = (data) => request.post('/auth/login', data)
// export const register = (data) => request.post('/auth/register', data)

export default {
  healthCheck,
}

