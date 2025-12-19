<template>
  <Layout>
    <div class="flex justify-between items-center mb-6">
      <h1 class="text-3xl font-bold text-gray-900">通知公告</h1>
      <button
        @click="showCreateModal = true"
        class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition"
      >
        发布通知
      </button>
    </div>

    <!-- 筛选 -->
    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-4">
        <select
          v-model="filters.status"
          class="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
          @change="loadNotices"
        >
          <option value="">全部状态</option>
          <option value="draft">草稿</option>
          <option value="published">已发布</option>
          <option value="archived">已归档</option>
        </select>
        <select
          v-model="filters.type"
          class="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
          @change="loadNotices"
        >
          <option value="">全部类型</option>
          <option value="notice">通知</option>
          <option value="announcement">公告</option>
          <option value="system">系统</option>
        </select>
        <select
          v-model="filters.priority"
          class="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
          @change="loadNotices"
        >
          <option value="">全部优先级</option>
          <option value="low">低</option>
          <option value="normal">普通</option>
          <option value="high">高</option>
          <option value="urgent">紧急</option>
        </select>
        <button
          @click="loadNotices"
          class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition"
        >
          筛选
        </button>
      </div>
    </div>

    <!-- 通知列表 -->
    <div class="space-y-4">
      <div
        v-for="notice in notices"
        :key="notice.id"
        class="bg-white rounded-lg shadow p-6 hover:shadow-lg transition cursor-pointer"
        @click="viewNotice(notice)"
      >
        <div class="flex justify-between items-start">
          <div class="flex-1">
            <div class="flex items-center gap-3 mb-2">
              <h3 class="text-xl font-semibold text-gray-900">{{ notice.title }}</h3>
              <span
                :class="{
                  'bg-red-100 text-red-800': notice.priority === 'urgent',
                  'bg-orange-100 text-orange-800': notice.priority === 'high',
                  'bg-blue-100 text-blue-800': notice.priority === 'normal',
                  'bg-gray-100 text-gray-800': notice.priority === 'low',
                }"
                class="px-2 py-1 text-xs font-semibold rounded-full"
              >
                {{ notice.priority === 'urgent' ? '紧急' : notice.priority === 'high' ? '高' : notice.priority === 'normal' ? '普通' : '低' }}
              </span>
            </div>
            <p class="text-gray-600 mb-3 line-clamp-2">{{ notice.content }}</p>
            <div class="flex items-center gap-4 text-sm text-gray-500">
              <span>发布人: {{ notice.creator?.username || '系统' }}</span>
              <span>发布时间: {{ formatDate(notice.publish_time) }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div v-if="notices.length === 0" class="text-center py-12 text-gray-500">
      暂无通知
    </div>

    <!-- 创建通知模态框 -->
    <div
      v-if="showCreateModal"
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 px-4"
      @click.self="showCreateModal = false"
    >
      <div class="bg-white rounded-xl shadow-xl p-8 max-w-2xl w-full max-h-[90vh] overflow-y-auto">
        <h2 class="text-2xl font-bold text-gray-900 mb-6">发布通知</h2>
        <form @submit.prevent="handleCreate" class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">标题</label>
            <input
              v-model="noticeForm.title"
              type="text"
              required
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">内容</label>
            <textarea
              v-model="noticeForm.content"
              rows="6"
              required
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
            ></textarea>
          </div>
          <div class="grid grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">类型</label>
              <select
                v-model="noticeForm.type"
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
              >
                <option value="notice">通知</option>
                <option value="announcement">公告</option>
                <option value="system">系统</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">优先级</label>
              <select
                v-model="noticeForm.priority"
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
              >
                <option value="low">低</option>
                <option value="normal">普通</option>
                <option value="high">高</option>
                <option value="urgent">紧急</option>
              </select>
            </div>
          </div>
          <div class="flex gap-3 pt-4">
            <button
              type="button"
              @click="showCreateModal = false"
              class="flex-1 bg-gray-200 text-gray-700 py-2 px-4 rounded-lg font-medium hover:bg-gray-300 transition"
            >
              取消
            </button>
            <button
              type="submit"
              :disabled="loading"
              class="flex-1 bg-blue-600 text-white py-2 px-4 rounded-lg font-medium hover:bg-blue-700 disabled:opacity-50 disabled:cursor-not-allowed transition"
            >
              <span v-if="loading">发布中...</span>
              <span v-else>发布</span>
            </button>
          </div>
        </form>
      </div>
    </div>
  </Layout>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import Layout from '../components/Layout.vue'
import { getNotices, createNotice } from '../api/notice'

const notices = ref([])
const showCreateModal = ref(false)
const loading = ref(false)

const filters = ref({
  status: '',
  type: '',
  priority: '',
})

const noticeForm = ref({
  title: '',
  content: '',
  type: 'notice',
  priority: 'normal',
  status: 'published',
})

function formatDate(dateString) {
  if (!dateString) return '-'
  return new Date(dateString).toLocaleString('zh-CN')
}

function viewNotice(notice) {
  // 查看通知详情
  console.log('查看通知:', notice)
}

async function loadNotices() {
  try {
    const response = await getNotices(filters.value)
    notices.value = response.data.notices || []
  } catch (error) {
    console.error('加载通知失败:', error)
  }
}

async function handleCreate() {
  loading.value = true
  try {
    await createNotice(noticeForm.value)
    showCreateModal.value = false
    noticeForm.value = {
      title: '',
      content: '',
      type: 'notice',
      priority: 'normal',
      status: 'published',
    }
    loadNotices()
  } catch (error) {
    console.error('创建通知失败:', error)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  loadNotices()
})
</script>

