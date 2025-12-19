<template>
  <Layout>
    <div class="flex justify-between items-center mb-6">
      <h1 class="text-3xl font-bold text-gray-900">消息中心</h1>
      <div class="flex items-center gap-4">
        <span v-if="unreadCount > 0" class="bg-red-500 text-white px-3 py-1 rounded-full text-sm">
          未读: {{ unreadCount }}
        </span>
      </div>
    </div>

    <!-- 筛选 -->
    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
        <select
          v-model="filters.status"
          class="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
          @change="loadMessages"
        >
          <option value="">全部状态</option>
          <option value="unread">未读</option>
          <option value="read">已读</option>
          <option value="archived">已归档</option>
        </select>
        <select
          v-model="filters.type"
          class="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
          @change="loadMessages"
        >
          <option value="">全部类型</option>
          <option value="system">系统消息</option>
          <option value="workflow">工作流</option>
          <option value="task">任务</option>
          <option value="custom">自定义</option>
        </select>
        <button
          @click="loadMessages"
          class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition"
        >
          筛选
        </button>
      </div>
    </div>

    <!-- 消息列表 -->
    <div class="space-y-4">
      <div
        v-for="message in messages"
        :key="message.id"
        :class="{
          'bg-blue-50 border-blue-200': !message.is_read,
          'bg-white': message.is_read,
        }"
        class="border rounded-lg shadow p-6 hover:shadow-lg transition cursor-pointer"
        @click="viewMessage(message)"
      >
        <div class="flex justify-between items-start">
          <div class="flex-1">
            <div class="flex items-center gap-3 mb-2">
              <h3 class="text-lg font-semibold text-gray-900">{{ message.title }}</h3>
              <span
                v-if="!message.is_read"
                class="bg-red-500 text-white text-xs px-2 py-1 rounded-full"
              >
                未读
              </span>
            </div>
            <p class="text-gray-600 mb-3">{{ message.content }}</p>
            <div class="flex items-center gap-4 text-sm text-gray-500">
              <span v-if="message.from_user">来自: {{ message.from_user.username }}</span>
              <span>时间: {{ formatDate(message.created_at) }}</span>
            </div>
          </div>
          <button
            v-if="!message.is_read"
            @click.stop="markAsRead(message.id)"
            class="ml-4 text-blue-600 hover:text-blue-800 text-sm"
          >
            标记已读
          </button>
        </div>
      </div>
    </div>

    <div v-if="messages.length === 0" class="text-center py-12 text-gray-500">
      暂无消息
    </div>
  </Layout>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import Layout from '../components/Layout.vue'
import { getMessages, markMessageAsRead, getUnreadMessageCount } from '../api/message'

const messages = ref([])
const unreadCount = ref(0)

const filters = ref({
  status: '',
  type: '',
})

function formatDate(dateString) {
  if (!dateString) return '-'
  return new Date(dateString).toLocaleString('zh-CN')
}

async function loadMessages() {
  try {
    const response = await getMessages(filters.value)
    messages.value = response.data.messages || []
    loadUnreadCount()
  } catch (error) {
    console.error('加载消息失败:', error)
  }
}

async function loadUnreadCount() {
  try {
    const response = await getUnreadMessageCount()
    unreadCount.value = response.data.count || 0
  } catch (error) {
    console.error('加载未读数量失败:', error)
  }
}

async function markAsRead(id) {
  try {
    await markMessageAsRead(id)
    loadMessages()
  } catch (error) {
    console.error('标记已读失败:', error)
  }
}

function viewMessage(message) {
  if (!message.is_read) {
    markAsRead(message.id)
  }
  // 可以打开消息详情
}

onMounted(() => {
  loadMessages()
})
</script>

