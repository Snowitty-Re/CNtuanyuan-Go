import request from './request'

// 上传文件
export const uploadFile = (file) => {
  const formData = new FormData()
  formData.append('file', file)
  return request.post('/files/upload', formData, {
    headers: {
      'Content-Type': 'multipart/form-data',
    },
  })
}

// 获取文件列表
export const getFiles = (params) => request.get('/files', { params })

// 获取文件详情
export const getFile = (id) => request.get(`/files/${id}`)

// 删除文件
export const deleteFile = (id) => request.delete(`/files/${id}`)

