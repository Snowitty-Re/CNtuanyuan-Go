<template>
  <Layout>
    <div class="flex justify-between items-center mb-6">
      <h1 class="text-3xl font-bold text-gray-900">文件管理</h1>
      <div class="flex gap-4">
        <label class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition cursor-pointer">
          上传文件
          <input
            type="file"
            @change="handleFileSelect"
            class="hidden"
            multiple
          />
        </label>
      </div>
    </div>

    <!-- 筛选 -->
    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
        <select
          v-model="filters.category"
          class="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
          @change="loadFiles"
        >
          <option value="">全部分类</option>
          <option value="document">文档</option>
          <option value="image">图片</option>
          <option value="video">视频</option>
          <option value="other">其他</option>
        </select>
        <button
          @click="loadFiles"
          class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition"
        >
          筛选
        </button>
      </div>
    </div>

    <!-- 文件列表 -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <div
        v-for="file in files"
        :key="file.id"
        class="bg-white rounded-lg shadow p-4 hover:shadow-lg transition"
      >
        <div class="flex items-start justify-between mb-2">
          <div class="flex-1">
            <h3 class="font-semibold text-gray-900 truncate">{{ file.original_name }}</h3>
            <p class="text-sm text-gray-500 mt-1">
              {{ formatFileSize(file.file_size) }} · {{ file.category }}
            </p>
          </div>
        </div>
        <div class="flex items-center gap-2 mt-4">
          <a
            :href="file.file_path"
            target="_blank"
            class="flex-1 bg-blue-600 text-white px-3 py-2 rounded-lg hover:bg-blue-700 transition text-sm text-center"
          >
            下载
          </a>
          <button
            @click="handleDelete(file.id)"
            class="flex-1 bg-red-600 text-white px-3 py-2 rounded-lg hover:bg-red-700 transition text-sm"
          >
            删除
          </button>
        </div>
      </div>
    </div>

    <div v-if="files.length === 0" class="text-center py-12 text-gray-500">
      暂无文件
    </div>
  </Layout>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import Layout from '../components/Layout.vue'
import { getFiles, uploadFile, deleteFile } from '../api/file'

const files = ref([])
const filters = ref({
  category: '',
})

function formatFileSize(bytes) {
  if (bytes === 0) return '0 B'
  const k = 1024
  const sizes = ['B', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return Math.round(bytes / Math.pow(k, i) * 100) / 100 + ' ' + sizes[i]
}

async function loadFiles() {
  try {
    const response = await getFiles(filters.value)
    files.value = response.data.files || []
  } catch (error) {
    console.error('加载文件失败:', error)
  }
}

async function handleFileSelect(event) {
  const selectedFiles = event.target.files
  if (!selectedFiles || selectedFiles.length === 0) return

  for (let i = 0; i < selectedFiles.length; i++) {
    try {
      await uploadFile(selectedFiles[i])
    } catch (error) {
      console.error('上传文件失败:', error)
    }
  }
  loadFiles()
  event.target.value = '' // 清空选择
}

async function handleDelete(id) {
  if (confirm('确定要删除这个文件吗？')) {
    try {
      await deleteFile(id)
      loadFiles()
    } catch (error) {
      console.error('删除文件失败:', error)
    }
  }
}

onMounted(() => {
  loadFiles()
})
</script>

