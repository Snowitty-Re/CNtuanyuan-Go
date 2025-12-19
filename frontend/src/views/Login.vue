<template>
  <div class="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 flex items-center justify-center px-4">
    <div class="max-w-md w-full bg-white rounded-xl shadow-lg p-8">
      <div class="text-center mb-8">
        <h1 class="text-3xl font-bold text-gray-900 mb-2">志愿者OA系统</h1>
        <p class="text-gray-600">走失人员寻亲管理系统</p>
      </div>

      <div v-if="error" class="mb-4 p-3 bg-red-50 border border-red-200 rounded-lg text-red-700 text-sm">
        {{ error }}
      </div>

      <form @submit.prevent="handleLogin" class="space-y-6">
        <div>
          <label for="username" class="block text-sm font-medium text-gray-700 mb-2">
            用户名
          </label>
          <input
            id="username"
            v-model="form.username"
            type="text"
            required
            class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none transition"
            placeholder="请输入用户名"
          />
        </div>

        <div>
          <label for="password" class="block text-sm font-medium text-gray-700 mb-2">
            密码
          </label>
          <input
            id="password"
            v-model="form.password"
            type="password"
            required
            class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none transition"
            placeholder="请输入密码"
          />
        </div>

        <button
          type="submit"
          :disabled="loading"
          class="w-full bg-blue-600 text-white py-2 px-4 rounded-lg font-medium hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 disabled:opacity-50 disabled:cursor-not-allowed transition"
        >
          <span v-if="loading">登录中...</span>
          <span v-else>登录</span>
        </button>
      </form>

      <div class="mt-6 text-center">
        <p class="text-sm text-gray-600">
          还没有账号？
          <a href="#" @click.prevent="showRegister = true" class="text-blue-600 hover:text-blue-700 font-medium">
            立即注册
          </a>
        </p>
      </div>
    </div>

    <!-- 注册对话框 -->
    <div
      v-if="showRegister"
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 px-4"
      @click.self="showRegister = false"
    >
      <div class="bg-white rounded-xl shadow-xl p-8 max-w-md w-full max-h-[90vh] overflow-y-auto">
        <h2 class="text-2xl font-bold text-gray-900 mb-6">注册账号</h2>

        <div v-if="registerError" class="mb-4 p-3 bg-red-50 border border-red-200 rounded-lg text-red-700 text-sm">
          {{ registerError }}
        </div>

        <form @submit.prevent="handleRegister" class="space-y-4">
          <div>
            <label for="reg-username" class="block text-sm font-medium text-gray-700 mb-2">
              用户名
            </label>
            <input
              id="reg-username"
              v-model="registerForm.username"
              type="text"
              required
              minlength="3"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none"
              placeholder="至少3个字符"
            />
          </div>

          <div>
            <label for="reg-email" class="block text-sm font-medium text-gray-700 mb-2">
              邮箱
            </label>
            <input
              id="reg-email"
              v-model="registerForm.email"
              type="email"
              required
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none"
              placeholder="example@email.com"
            />
          </div>

          <div>
            <label for="reg-password" class="block text-sm font-medium text-gray-700 mb-2">
              密码
            </label>
            <input
              id="reg-password"
              v-model="registerForm.password"
              type="password"
              required
              minlength="6"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none"
              placeholder="至少6个字符"
            />
          </div>

          <div class="flex gap-3">
            <button
              type="button"
              @click="showRegister = false"
              class="flex-1 bg-gray-200 text-gray-700 py-2 px-4 rounded-lg font-medium hover:bg-gray-300 transition"
            >
              取消
            </button>
            <button
              type="submit"
              :disabled="registerLoading"
              class="flex-1 bg-blue-600 text-white py-2 px-4 rounded-lg font-medium hover:bg-blue-700 disabled:opacity-50 disabled:cursor-not-allowed transition"
            >
              <span v-if="registerLoading">注册中...</span>
              <span v-else>注册</span>
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useAuthStore } from '../stores/auth'

const router = useRouter()
const route = useRoute()
const authStore = useAuthStore()

const form = ref({
  username: '',
  password: '',
})

const registerForm = ref({
  username: '',
  email: '',
  password: '',
})

const loading = ref(false)
const registerLoading = ref(false)
const error = ref('')
const registerError = ref('')
const showRegister = ref(false)

async function handleLogin() {
  loading.value = true
  error.value = ''

  try {
    await authStore.loginUser(form.value)
    const redirect = route.query.redirect || '/'
    router.push(redirect)
  } catch (err) {
    // 错误可能来自响应拦截器，message已经设置
    error.value = err.message || err.response?.data?.error || '登录失败，请检查用户名和密码'
  } finally {
    loading.value = false
  }
}

async function handleRegister() {
  registerLoading.value = true
  registerError.value = ''

  try {
    await authStore.registerUser(registerForm.value)
    showRegister.value = false
    const redirect = route.query.redirect || '/'
    router.push(redirect)
  } catch (err) {
    // 错误可能来自响应拦截器，message已经设置
    registerError.value = err.message || err.response?.data?.error || '注册失败，请重试'
  } finally {
    registerLoading.value = false
  }
}
</script>
