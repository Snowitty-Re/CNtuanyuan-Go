import request from './request'

// 发送消息
export const sendMessage = (data) => request.post('/messages', data)

// 获取消息列表
export const getMessages = (params) => request.get('/messages', { params })

// 获取消息详情
export const getMessage = (id) => request.get(`/messages/${id}`)

// 标记为已读
export const markMessageAsRead = (id) => request.post(`/messages/${id}/read`)

// 获取未读数量
export const getUnreadMessageCount = () => request.get('/messages/unread/count')

