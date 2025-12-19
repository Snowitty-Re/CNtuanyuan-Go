import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { login, register, getProfile } from '../api/auth'
import router from '../router'

export const useAuthStore = defineStore('auth', () => {
  const user = ref(null)
  const token = ref(localStorage.getItem('token') || null)

  const isAuthenticated = computed(() => !!token.value)

  // 登录
  async function loginUser(credentials) {
    try {
      const response = await login(credentials)
      // 后端返回格式: {success: true, data: {user: {...}, token: "..."}, message: "..."}
      if (response.success && response.data) {
        token.value = response.data.token
        user.value = response.data.user
        localStorage.setItem('token', token.value)
        return response
      } else {
        throw new Error(response.error || '登录失败')
      }
    } catch (error) {
      throw error
    }
  }

  // 注册
  async function registerUser(userData) {
    try {
      const response = await register(userData)
      // 后端返回格式: {success: true, data: {user: {...}, token: "..."}, message: "..."}
      if (response.success && response.data) {
        token.value = response.data.token
        user.value = response.data.user
        localStorage.setItem('token', token.value)
        return response
      } else {
        throw new Error(response.error || '注册失败')
      }
    } catch (error) {
      throw error
    }
  }

  // 获取用户信息
  async function fetchUser() {
    try {
      const response = await getProfile()
      // 后端返回格式: {success: true, data: {...}, message: "..."}
      if (response.success && response.data) {
        user.value = response.data
        return response
      } else {
        throw new Error(response.error || '获取用户信息失败')
      }
    } catch (error) {
      logout()
      throw error
    }
  }

  // 登出
  function logout() {
    user.value = null
    token.value = null
    localStorage.removeItem('token')
    router.push('/login')
  }

  // 初始化时如果有token，获取用户信息（静默失败，不自动登出）
  if (token.value) {
    fetchUser().catch((err) => {
      // 只在明确是401错误时才清除token
      console.warn('获取用户信息失败，但保留token:', err)
      // 不清除token，让用户手动重试
    })
  }

  return {
    user,
    token,
    isAuthenticated,
    loginUser,
    registerUser,
    fetchUser,
    logout,
  }
})

