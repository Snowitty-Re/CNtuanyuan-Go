<template>
  <div class="min-h-screen bg-gray-50">
    <!-- 导航栏 -->
    <nav class="bg-white shadow-sm border-b">
      <div class="container mx-auto px-4">
        <div class="flex items-center justify-between h-16">
          <div class="flex items-center">
            <router-link to="/" class="text-xl font-bold text-gray-900">
              志愿者OA系统
            </router-link>
          </div>
          <div class="hidden md:flex items-center space-x-4">
            <router-link
              to="/dashboard"
              class="px-3 py-2 text-sm font-medium text-gray-700 hover:text-gray-900 hover:bg-gray-100 rounded-lg transition"
              active-class="bg-blue-50 text-blue-700"
            >
              仪表盘
            </router-link>
            <router-link
              to="/missing-persons"
              class="px-3 py-2 text-sm font-medium text-gray-700 hover:text-gray-900 hover:bg-gray-100 rounded-lg transition"
              active-class="bg-blue-50 text-blue-700"
            >
              走失人员
            </router-link>
            <router-link
              to="/workflows"
              class="px-3 py-2 text-sm font-medium text-gray-700 hover:text-gray-900 hover:bg-gray-100 rounded-lg transition"
              active-class="bg-blue-50 text-blue-700"
            >
              工作流
            </router-link>
            <router-link
              to="/forms"
              class="px-3 py-2 text-sm font-medium text-gray-700 hover:text-gray-900 hover:bg-gray-100 rounded-lg transition"
              active-class="bg-blue-50 text-blue-700"
            >
              表单
            </router-link>
            <router-link
              to="/roles-permissions"
              class="px-3 py-2 text-sm font-medium text-gray-700 hover:text-gray-900 hover:bg-gray-100 rounded-lg transition"
              active-class="bg-blue-50 text-blue-700"
            >
              权限管理
            </router-link>
            <router-link
              to="/notices"
              class="px-3 py-2 text-sm font-medium text-gray-700 hover:text-gray-900 hover:bg-gray-100 rounded-lg transition"
              active-class="bg-blue-50 text-blue-700"
            >
              通知公告
            </router-link>
            <router-link
              to="/messages"
              class="px-3 py-2 text-sm font-medium text-gray-700 hover:text-gray-900 hover:bg-gray-100 rounded-lg transition"
              active-class="bg-blue-50 text-blue-700"
            >
              消息中心
            </router-link>
            <router-link
              to="/files"
              class="px-3 py-2 text-sm font-medium text-gray-700 hover:text-gray-900 hover:bg-gray-100 rounded-lg transition"
              active-class="bg-blue-50 text-blue-700"
            >
              文件管理
            </router-link>
            <router-link
              to="/tasks"
              class="px-3 py-2 text-sm font-medium text-gray-700 hover:text-gray-900 hover:bg-gray-100 rounded-lg transition"
              active-class="bg-blue-50 text-blue-700"
            >
              任务管理
            </router-link>
          </div>
          <div class="flex items-center space-x-4">
            <div v-if="authStore.user" class="hidden md:block text-sm text-gray-700">
              {{ authStore.user.username }}
            </div>
            <button
              @click="handleLogout"
              class="px-4 py-2 text-sm text-gray-700 hover:text-gray-900 hover:bg-gray-100 rounded-lg transition"
            >
              退出
            </button>
            <!-- 移动端菜单按钮 -->
            <button
              @click="mobileMenuOpen = !mobileMenuOpen"
              class="md:hidden p-2 text-gray-700 hover:text-gray-900"
            >
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
              </svg>
            </button>
          </div>
        </div>
      </div>
      <!-- 移动端菜单 -->
      <div v-if="mobileMenuOpen" class="md:hidden border-t">
        <div class="px-4 py-2 space-y-1">
          <router-link
            to="/dashboard"
            @click="mobileMenuOpen = false"
            class="block px-3 py-2 text-sm font-medium text-gray-700 hover:text-gray-900 hover:bg-gray-100 rounded-lg transition"
          >
            仪表盘
          </router-link>
          <router-link
            to="/missing-persons"
            @click="mobileMenuOpen = false"
            class="block px-3 py-2 text-sm font-medium text-gray-700 hover:text-gray-900 hover:bg-gray-100 rounded-lg transition"
          >
            走失人员
          </router-link>
          <router-link
            to="/workflows"
            @click="mobileMenuOpen = false"
            class="block px-3 py-2 text-sm font-medium text-gray-700 hover:text-gray-900 hover:bg-gray-100 rounded-lg transition"
          >
            工作流
          </router-link>
          <router-link
            to="/forms"
            @click="mobileMenuOpen = false"
            class="block px-3 py-2 text-sm font-medium text-gray-700 hover:text-gray-900 hover:bg-gray-100 rounded-lg transition"
          >
            表单
          </router-link>
          <router-link
            to="/roles-permissions"
            @click="mobileMenuOpen = false"
            class="block px-3 py-2 text-sm font-medium text-gray-700 hover:text-gray-900 hover:bg-gray-100 rounded-lg transition"
          >
            权限管理
          </router-link>
          <router-link
            to="/notices"
            @click="mobileMenuOpen = false"
            class="block px-3 py-2 text-sm font-medium text-gray-700 hover:text-gray-900 hover:bg-gray-100 rounded-lg transition"
          >
            通知公告
          </router-link>
          <router-link
            to="/messages"
            @click="mobileMenuOpen = false"
            class="block px-3 py-2 text-sm font-medium text-gray-700 hover:text-gray-900 hover:bg-gray-100 rounded-lg transition"
          >
            消息中心
          </router-link>
          <router-link
            to="/files"
            @click="mobileMenuOpen = false"
            class="block px-3 py-2 text-sm font-medium text-gray-700 hover:text-gray-900 hover:bg-gray-100 rounded-lg transition"
          >
            文件管理
          </router-link>
          <router-link
            to="/tasks"
            @click="mobileMenuOpen = false"
            class="block px-3 py-2 text-sm font-medium text-gray-700 hover:text-gray-900 hover:bg-gray-100 rounded-lg transition"
          >
            任务管理
          </router-link>
        </div>
      </div>
    </nav>

    <!-- 主内容区 -->
    <main class="container mx-auto px-4 py-8">
      <slot />
    </main>

    <!-- 页脚 -->
    <footer class="bg-white border-t mt-12">
      <div class="container mx-auto px-4 py-6">
        <p class="text-center text-sm text-gray-600">
          © 2024 志愿者OA系统 - 走失人员寻亲管理系统
        </p>
      </div>
    </footer>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'

const router = useRouter()
const authStore = useAuthStore()
const mobileMenuOpen = ref(false)

function handleLogout() {
  authStore.logout()
}
</script>

