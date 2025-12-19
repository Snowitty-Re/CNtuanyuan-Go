import request from './request'

// 获取走失人员列表
export const getMissingPersons = (params) => request.get('/missing-persons', { params })

// 搜索走失人员
export const searchMissingPersons = (params) => request.get('/missing-persons/search', { params })

// 获取走失人员详情
export const getMissingPerson = (id) => request.get(`/missing-persons/${id}`)

// 创建走失人员
export const createMissingPerson = (data) => request.post('/missing-persons', data)

// 更新走失人员
export const updateMissingPerson = (id, data) => request.put(`/missing-persons/${id}`, data)

// 删除走失人员
export const deleteMissingPerson = (id) => request.delete(`/missing-persons/${id}`)

// 上传照片
export const uploadPhoto = (file) => {
  const formData = new FormData()
  formData.append('photo', file)
  return request.post('/missing-persons/upload-photo', formData, {
    headers: {
      'Content-Type': 'multipart/form-data',
    },
  })
}

