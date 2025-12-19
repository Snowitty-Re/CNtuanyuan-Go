import request from './request'

// 创建通知
export const createNotice = (data) => request.post('/notices', data)

// 获取通知列表
export const getNotices = (params) => request.get('/notices', { params })

// 获取通知详情
export const getNotice = (id) => request.get(`/notices/${id}`)

// 标记为已读
export const markNoticeAsRead = (id) => request.post(`/notices/${id}/read`)

// 获取未读数量
export const getUnreadNoticeCount = () => request.get('/notices/unread/count')

