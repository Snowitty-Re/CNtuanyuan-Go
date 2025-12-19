<template>
  <div class="min-h-screen bg-gray-50">
    <nav class="bg-white shadow-sm border-b">
      <div class="container mx-auto px-4 py-4">
        <div class="flex items-center justify-between">
          <h1 class="text-xl font-bold text-gray-900">志愿者OA系统</h1>
          <div class="flex items-center gap-4">
            <span v-if="authStore.user" class="text-gray-700">
              {{ authStore.user.username }}
            </span>
            <button
              @click="handleLogout"
              class="px-4 py-2 text-sm text-gray-700 hover:text-gray-900 hover:bg-gray-100 rounded-lg transition"
            >
              退出登录
            </button>
          </div>
        </div>
      </div>
    </nav>

    <div class="container mx-auto px-4 py-8">
      <div class="bg-white rounded-lg shadow p-6">
        <h2 class="text-2xl font-bold text-gray-900 mb-4">欢迎使用志愿者OA系统</h2>
        <p class="text-gray-600 mb-6">
          这是一个用于管理走失人员寻亲的系统，包含完整的权限管理、工作流、自定义表单等功能。
        </p>

        <div v-if="authStore.user" class="mt-6">
          <h3 class="text-lg font-semibold text-gray-900 mb-2">用户信息</h3>
          <div class="bg-gray-50 rounded-lg p-4">
            <p class="text-sm text-gray-600">
              <span class="font-medium">用户名：</span>{{ authStore.user.username }}
            </p>
            <p class="text-sm text-gray-600">
              <span class="font-medium">邮箱：</span>{{ authStore.user.email }}
            </p>
          </div>
        </div>

        <div class="mt-8 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
          <router-link
            to="/dashboard"
            class="bg-blue-50 rounded-lg p-4 border border-blue-200 hover:bg-blue-100 transition"
          >
            <h4 class="font-semibold text-blue-900 mb-2">仪表盘</h4>
            <p class="text-sm text-blue-700">查看系统统计和快速操作</p>
          </router-link>
          <router-link
            to="/missing-persons"
            class="bg-purple-50 rounded-lg p-4 border border-purple-200 hover:bg-purple-100 transition"
          >
            <h4 class="font-semibold text-purple-900 mb-2">走失人员</h4>
            <p class="text-sm text-purple-700">走失人员信息管理数据库</p>
          </router-link>
          <router-link
            to="/workflows"
            class="bg-green-50 rounded-lg p-4 border border-green-200 hover:bg-green-100 transition"
          >
            <h4 class="font-semibold text-green-900 mb-2">工作流</h4>
            <p class="text-sm text-green-700">可自定义字段的工作流系统</p>
          </router-link>
          <router-link
            to="/forms"
            class="bg-yellow-50 rounded-lg p-4 border border-yellow-200 hover:bg-yellow-100 transition"
          >
            <h4 class="font-semibold text-yellow-900 mb-2">表单管理</h4>
            <p class="text-sm text-yellow-700">自定义表单构建和管理</p>
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'

const router = useRouter()
const authStore = useAuthStore()

function handleLogout() {
  authStore.logout()
}
</script>
