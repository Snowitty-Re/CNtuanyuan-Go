<template>
  <div class="min-h-screen bg-gray-50">
    <div class="container mx-auto px-4 py-8">
      <div class="mb-6">
        <button
          @click="$router.back()"
          class="text-blue-600 hover:text-blue-800 mb-4"
        >
          ← 返回
        </button>
        <h1 class="text-3xl font-bold text-gray-900">{{ workflow?.name || '工作流详情' }}</h1>
      </div>

      <div v-if="loading" class="text-center py-12">加载中...</div>

      <div v-else-if="workflow" class="space-y-6">
        <!-- 基本信息 -->
        <div class="bg-white rounded-lg shadow p-6">
          <h2 class="text-xl font-semibold text-gray-900 mb-4">基本信息</h2>
          <p class="text-gray-600">{{ workflow.description || '无描述' }}</p>
        </div>

        <!-- 字段管理 -->
        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex justify-between items-center mb-4">
            <h2 class="text-xl font-semibold text-gray-900">自定义字段</h2>
            <button
              @click="showAddFieldModal = true"
              class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition text-sm"
            >
              添加字段
            </button>
          </div>

          <div v-if="workflow.fields && workflow.fields.length > 0" class="space-y-2">
            <div
              v-for="field in workflow.fields"
              :key="field.id"
              class="flex items-center justify-between p-3 bg-gray-50 rounded-lg"
            >
              <div>
                <span class="font-medium text-gray-900">{{ field.field_name }}</span>
                <span class="ml-2 text-sm text-gray-500">({{ field.field_type }})</span>
                <span v-if="field.required" class="ml-2 text-xs text-red-600">必填</span>
              </div>
            </div>
          </div>
          <div v-else class="text-gray-500 text-sm">暂无字段</div>
        </div>

        <!-- 状态管理 -->
        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex justify-between items-center mb-4">
            <h2 class="text-xl font-semibold text-gray-900">工作流状态</h2>
            <button
              @click="showAddStateModal = true"
              class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition text-sm"
            >
              添加状态
            </button>
          </div>

          <div v-if="workflow.states && workflow.states.length > 0" class="space-y-2">
            <div
              v-for="state in workflow.states"
              :key="state.id"
              class="flex items-center justify-between p-3 bg-gray-50 rounded-lg"
            >
              <div>
                <span class="font-medium text-gray-900">{{ state.state_name }}</span>
                <span class="ml-2 text-sm text-gray-500">({{ state.state_type }})</span>
              </div>
            </div>
          </div>
          <div v-else class="text-gray-500 text-sm">暂无状态</div>
        </div>
      </div>

      <!-- 添加字段模态框 -->
      <div
        v-if="showAddFieldModal"
        class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 px-4"
        @click.self="showAddFieldModal = false"
      >
        <div class="bg-white rounded-xl shadow-xl p-8 max-w-md w-full">
          <h3 class="text-xl font-bold text-gray-900 mb-6">添加字段</h3>
          <form @submit.prevent="handleAddField" class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">字段名称</label>
              <input
                v-model="fieldForm.field_name"
                type="text"
                required
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">字段类型</label>
              <select
                v-model="fieldForm.field_type"
                required
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
              >
                <option value="text">文本</option>
                <option value="number">数字</option>
                <option value="date">日期</option>
                <option value="select">选择</option>
                <option value="textarea">多行文本</option>
              </select>
            </div>
            <div class="flex items-center">
              <input
                v-model="fieldForm.required"
                type="checkbox"
                class="mr-2"
              />
              <label class="text-sm text-gray-700">必填</label>
            </div>
            <div class="flex gap-3 pt-4">
              <button
                type="button"
                @click="showAddFieldModal = false"
                class="flex-1 bg-gray-200 text-gray-700 py-2 px-4 rounded-lg font-medium hover:bg-gray-300 transition"
              >
                取消
              </button>
              <button
                type="submit"
                class="flex-1 bg-blue-600 text-white py-2 px-4 rounded-lg font-medium hover:bg-blue-700 transition"
              >
                添加
              </button>
            </div>
          </form>
        </div>
      </div>

      <!-- 添加状态模态框 -->
      <div
        v-if="showAddStateModal"
        class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 px-4"
        @click.self="showAddStateModal = false"
      >
        <div class="bg-white rounded-xl shadow-xl p-8 max-w-md w-full">
          <h3 class="text-xl font-bold text-gray-900 mb-6">添加状态</h3>
          <form @submit.prevent="handleAddState" class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">状态名称</label>
              <input
                v-model="stateForm.state_name"
                type="text"
                required
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">状态类型</label>
              <select
                v-model="stateForm.state_type"
                required
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
              >
                <option value="initial">初始状态</option>
                <option value="intermediate">中间状态</option>
                <option value="final">最终状态</option>
              </select>
            </div>
            <div class="flex gap-3 pt-4">
              <button
                type="button"
                @click="showAddStateModal = false"
                class="flex-1 bg-gray-200 text-gray-700 py-2 px-4 rounded-lg font-medium hover:bg-gray-300 transition"
              >
                取消
              </button>
              <button
                type="submit"
                class="flex-1 bg-blue-600 text-white py-2 px-4 rounded-lg font-medium hover:bg-blue-700 transition"
              >
                添加
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { getWorkflow, addField, addState } from '../api/workflow'

const route = useRoute()
const workflow = ref(null)
const loading = ref(true)
const showAddFieldModal = ref(false)
const showAddStateModal = ref(false)

const fieldForm = ref({
  workflow_id: 0,
  field_name: '',
  field_type: 'text',
  required: false,
})

const stateForm = ref({
  workflow_id: 0,
  state_name: '',
  state_type: 'initial',
})

async function loadWorkflow() {
  try {
    const id = route.params.id
    const response = await getWorkflow(id)
    workflow.value = response.data
    fieldForm.value.workflow_id = workflow.value.id
    stateForm.value.workflow_id = workflow.value.id
  } catch (error) {
    console.error('加载工作流失败:', error)
  } finally {
    loading.value = false
  }
}

async function handleAddField() {
  try {
    await addField(fieldForm.value)
    showAddFieldModal.value = false
    fieldForm.value = {
      workflow_id: workflow.value.id,
      field_name: '',
      field_type: 'text',
      required: false,
    }
    loadWorkflow()
  } catch (error) {
    console.error('添加字段失败:', error)
  }
}

async function handleAddState() {
  try {
    await addState(stateForm.value)
    showAddStateModal.value = false
    stateForm.value = {
      workflow_id: workflow.value.id,
      state_name: '',
      state_type: 'initial',
    }
    loadWorkflow()
  } catch (error) {
    console.error('添加状态失败:', error)
  }
}

onMounted(() => {
  loadWorkflow()
})
</script>

