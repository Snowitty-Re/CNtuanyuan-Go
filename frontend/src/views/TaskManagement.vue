<template>
  <Layout>
    <div class="flex justify-between items-center mb-6">
      <h1 class="text-3xl font-bold text-gray-900">任务管理</h1>
      <button
        @click="showCreateModal = true"
        class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition"
      >
        创建任务
      </button>
    </div>

    <!-- 筛选 -->
    <div class="bg-white rounded-lg shadow p-4 mb-6">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-4">
        <select
          v-model="filters.status"
          class="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
          @change="loadTasks"
        >
          <option value="">全部状态</option>
          <option value="todo">待办</option>
          <option value="in_progress">进行中</option>
          <option value="done">已完成</option>
          <option value="cancelled">已取消</option>
        </select>
        <select
          v-model="filters.priority"
          class="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
          @change="loadTasks"
        >
          <option value="">全部优先级</option>
          <option value="low">低</option>
          <option value="normal">普通</option>
          <option value="high">高</option>
          <option value="urgent">紧急</option>
        </select>
        <button
          @click="loadTasks"
          class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition"
        >
          筛选
        </button>
      </div>
    </div>

    <!-- 任务列表 -->
    <div class="space-y-4">
      <div
        v-for="task in tasks"
        :key="task.id"
        class="bg-white rounded-lg shadow p-6 hover:shadow-lg transition"
      >
        <div class="flex justify-between items-start">
          <div class="flex-1">
            <div class="flex items-center gap-3 mb-2">
              <h3 class="text-xl font-semibold text-gray-900">{{ task.title }}</h3>
              <span
                :class="{
                  'bg-red-100 text-red-800': task.priority === 'urgent',
                  'bg-orange-100 text-orange-800': task.priority === 'high',
                  'bg-blue-100 text-blue-800': task.priority === 'normal',
                  'bg-gray-100 text-gray-800': task.priority === 'low',
                }"
                class="px-2 py-1 text-xs font-semibold rounded-full"
              >
                {{ task.priority === 'urgent' ? '紧急' : task.priority === 'high' ? '高' : task.priority === 'normal' ? '普通' : '低' }}
              </span>
              <span
                :class="{
                  'bg-yellow-100 text-yellow-800': task.status === 'todo',
                  'bg-blue-100 text-blue-800': task.status === 'in_progress',
                  'bg-green-100 text-green-800': task.status === 'done',
                  'bg-gray-100 text-gray-800': task.status === 'cancelled',
                }"
                class="px-2 py-1 text-xs font-semibold rounded-full"
              >
                {{ task.status === 'todo' ? '待办' : task.status === 'in_progress' ? '进行中' : task.status === 'done' ? '已完成' : '已取消' }}
              </span>
            </div>
            <p class="text-gray-600 mb-3">{{ task.description }}</p>
            <div class="flex items-center gap-4 text-sm text-gray-500 mb-2">
              <span>创建人: {{ task.creator?.username }}</span>
              <span v-if="task.assignee">负责人: {{ task.assignee.username }}</span>
              <span v-if="task.due_date">截止: {{ formatDate(task.due_date) }}</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div
                class="bg-blue-600 h-2 rounded-full transition-all"
                :style="{ width: task.progress + '%' }"
              ></div>
            </div>
            <p class="text-xs text-gray-500 mt-1">进度: {{ task.progress }}%</p>
          </div>
          <div class="ml-4 flex gap-2">
            <button
              @click="editTask(task)"
              class="text-blue-600 hover:text-blue-800 text-sm"
            >
              编辑
            </button>
            <button
              @click="handleDelete(task.id)"
              class="text-red-600 hover:text-red-800 text-sm"
            >
              删除
            </button>
          </div>
        </div>
      </div>
    </div>

    <div v-if="tasks.length === 0" class="text-center py-12 text-gray-500">
      暂无任务
    </div>

    <!-- 创建/编辑任务模态框 -->
    <div
      v-if="showCreateModal || editingTask"
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 px-4"
      @click.self="closeModal"
    >
      <div class="bg-white rounded-xl shadow-xl p-8 max-w-2xl w-full max-h-[90vh] overflow-y-auto">
        <h2 class="text-2xl font-bold text-gray-900 mb-6">
          {{ editingTask ? '编辑任务' : '创建任务' }}
        </h2>
        <form @submit.prevent="handleSubmit" class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">标题</label>
            <input
              v-model="taskForm.title"
              type="text"
              required
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">描述</label>
            <textarea
              v-model="taskForm.description"
              rows="4"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
            ></textarea>
          </div>
          <div class="grid grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">优先级</label>
              <select
                v-model="taskForm.priority"
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
              >
                <option value="low">低</option>
                <option value="normal">普通</option>
                <option value="high">高</option>
                <option value="urgent">紧急</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">状态</label>
              <select
                v-model="taskForm.status"
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
              >
                <option value="todo">待办</option>
                <option value="in_progress">进行中</option>
                <option value="done">已完成</option>
                <option value="cancelled">已取消</option>
              </select>
            </div>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">进度 (%)</label>
            <input
              v-model.number="taskForm.progress"
              type="number"
              min="0"
              max="100"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">截止日期</label>
            <input
              v-model="taskForm.due_date"
              type="date"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
            />
          </div>
          <div class="flex gap-3 pt-4">
            <button
              type="button"
              @click="closeModal"
              class="flex-1 bg-gray-200 text-gray-700 py-2 px-4 rounded-lg font-medium hover:bg-gray-300 transition"
            >
              取消
            </button>
            <button
              type="submit"
              :disabled="loading"
              class="flex-1 bg-blue-600 text-white py-2 px-4 rounded-lg font-medium hover:bg-blue-700 disabled:opacity-50 disabled:cursor-not-allowed transition"
            >
              <span v-if="loading">保存中...</span>
              <span v-else>保存</span>
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
import { getTasks, createTask, updateTask, deleteTask } from '../api/task'

const tasks = ref([])
const showCreateModal = ref(false)
const editingTask = ref(null)
const loading = ref(false)

const filters = ref({
  status: '',
  priority: '',
})

const taskForm = ref({
  title: '',
  description: '',
  priority: 'normal',
  status: 'todo',
  progress: 0,
  due_date: '',
})

function formatDate(dateString) {
  if (!dateString) return '-'
  return new Date(dateString).toLocaleDateString('zh-CN')
}

async function loadTasks() {
  try {
    const response = await getTasks(filters.value)
    tasks.value = response.data.tasks || []
  } catch (error) {
    console.error('加载任务失败:', error)
  }
}

function editTask(task) {
  editingTask.value = task
  taskForm.value = {
    title: task.title,
    description: task.description,
    priority: task.priority,
    status: task.status,
    progress: task.progress,
    due_date: task.due_date ? task.due_date.split('T')[0] : '',
  }
}

function closeModal() {
  showCreateModal.value = false
  editingTask.value = null
  taskForm.value = {
    title: '',
    description: '',
    priority: 'normal',
    status: 'todo',
    progress: 0,
    due_date: '',
  }
}

async function handleSubmit() {
  loading.value = true
  try {
    if (editingTask.value) {
      await updateTask(editingTask.value.id, taskForm.value)
    } else {
      await createTask(taskForm.value)
    }
    closeModal()
    loadTasks()
  } catch (error) {
    console.error('保存任务失败:', error)
  } finally {
    loading.value = false
  }
}

async function handleDelete(id) {
  if (confirm('确定要删除这个任务吗？')) {
    try {
      await deleteTask(id)
      loadTasks()
    } catch (error) {
      console.error('删除任务失败:', error)
    }
  }
}

onMounted(() => {
  loadTasks()
})
</script>

