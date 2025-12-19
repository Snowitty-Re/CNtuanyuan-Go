import request from './request'

// 健康检查
export const healthCheck = () => request.get('/health')

export default {
  healthCheck,
}

