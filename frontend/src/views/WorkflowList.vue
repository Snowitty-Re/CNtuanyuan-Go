<template>
  <Layout>
      <div class="flex justify-between items-center mb-6">
        <h1 class="text-3xl font-bold text-gray-900">工作流管理</h1>
        <button
          @click="showCreateModal = true"
          class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition"
        >
          创建工作流
        </button>
      </div>

      <!-- 工作流列表 -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div
          v-for="workflow in workflows"
          :key="workflow.id"
          class="bg-white rounded-lg shadow p-6 hover:shadow-lg transition"
        >
          <h3 class="text-xl font-semibold text-gray-900 mb-2">{{ workflow.name }}</h3>
          <p class="text-gray-600 text-sm mb-4">{{ workflow.description || '无描述' }}</p>
          <div class="flex items-center justify-between">
            <span
              :class="{
                'bg-green-100 text-green-800': workflow.status === 'active',
                'bg-gray-100 text-gray-800': workflow.status === 'inactive',
              }"
              class="px-2 py-1 text-xs font-semibold rounded-full"
            >
              {{ workflow.status === 'active' ? '启用' : '禁用' }}
            </span>
            <div class="flex gap-2">
              <button
                @click="editWorkflow(workflow)"
                class="text-blue-600 hover:text-blue-800 text-sm"
              >
                编辑
              </button>
              <button
                @click="viewWorkflow(workflow)"
                class="text-green-600 hover:text-green-800 text-sm"
              >
                查看
              </button>
            </div>
          </div>
        </div>
      </div>

      <div v-if="workflows.length === 0" class="text-center py-12 text-gray-500">
        暂无工作流，点击上方按钮创建一个
      </div>

      <!-- 创建/编辑工作流模态框 -->
      <div
        v-if="showCreateModal || editingWorkflow"
        class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 px-4"
        @click.self="closeModal"
      >
        <div class="bg-white rounded-xl shadow-xl p-8 max-w-2xl w-full max-h-[90vh] overflow-y-auto">
          <h2 class="text-2xl font-bold text-gray-900 mb-6">
            {{ editingWorkflow ? '编辑工作流' : '创建工作流' }}
          </h2>

          <form @submit.prevent="handleSubmit" class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">名称</label>
              <input
                v-model="workflowForm.name"
                type="text"
                required
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
                placeholder="工作流名称"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">描述</label>
              <textarea
                v-model="workflowForm.description"
                rows="3"
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
                placeholder="工作流描述"
              ></textarea>
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
import { useRouter } from 'vue-router'
import Layout from '../components/Layout.vue'
import { getWorkflows, createWorkflow } from '../api/workflow'

const router = useRouter()
const workflows = ref([])
const showCreateModal = ref(false)
const editingWorkflow = ref(null)
const loading = ref(false)

const workflowForm = ref({
  name: '',
  description: '',
})

async function loadWorkflows() {
  try {
    const response = await getWorkflows()
    workflows.value = response.data.workflows || []
  } catch (error) {
    console.error('加载工作流失败:', error)
  }
}

function editWorkflow(workflow) {
  editingWorkflow.value = workflow
  workflowForm.value = {
    name: workflow.name,
    description: workflow.description,
  }
}

function viewWorkflow(workflow) {
  router.push(`/workflows/${workflow.id}`)
}

function closeModal() {
  showCreateModal.value = false
  editingWorkflow.value = null
  workflowForm.value = {
    name: '',
    description: '',
  }
}

async function handleSubmit() {
  loading.value = true
  try {
    await createWorkflow(workflowForm.value)
    closeModal()
    loadWorkflows()
  } catch (error) {
    console.error('保存失败:', error)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  loadWorkflows()
})
</script>

