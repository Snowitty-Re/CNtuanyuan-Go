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
      token.value = response.data.token
      user.value = response.data.user
      localStorage.setItem('token', token.value)
      return response
    } catch (error) {
      throw error
    }
  }

  // 注册
  async function registerUser(userData) {
    try {
      const response = await register(userData)
      token.value = response.data.token
      user.value = response.data.user
      localStorage.setItem('token', token.value)
      return response
    } catch (error) {
      throw error
    }
  }

  // 获取用户信息
  async function fetchUser() {
    try {
      const response = await getProfile()
      user.value = response.data
      return response
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

  // 初始化时如果有token，获取用户信息
  if (token.value) {
    fetchUser().catch(() => {
      logout()
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

