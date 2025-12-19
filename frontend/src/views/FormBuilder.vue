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
        <h1 class="text-3xl font-bold text-gray-900">表单构建器</h1>
      </div>

      <div class="bg-white rounded-lg shadow p-6 mb-6">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">表单名称</label>
            <input
              v-model="formName"
              type="text"
              required
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
              placeholder="请输入表单名称"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">表单描述</label>
            <input
              v-model="formDescription"
              type="text"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
              placeholder="请输入表单描述"
            />
          </div>
        </div>
      </div>

      <FormBuilder v-model="formFields" @preview="handlePreview" />

      <div class="mt-6 flex justify-end gap-4">
        <button
          @click="$router.back()"
          class="px-6 py-2 border border-gray-300 rounded-lg text-gray-700 hover:bg-gray-50 transition"
        >
          取消
        </button>
        <button
          @click="saveForm"
          :disabled="!formName || formFields.length === 0"
          class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 disabled:opacity-50 disabled:cursor-not-allowed transition"
        >
          保存表单
        </button>
      </div>

      <!-- 预览模态框 -->
      <div
        v-if="showPreview"
        class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 px-4"
        @click.self="showPreview = false"
      >
        <div class="bg-white rounded-xl shadow-xl p-8 max-w-2xl w-full max-h-[90vh] overflow-y-auto">
          <div class="flex justify-between items-center mb-6">
            <h3 class="text-xl font-bold text-gray-900">表单预览</h3>
            <button
              @click="showPreview = false"
              class="text-gray-500 hover:text-gray-700"
            >
              ✕
            </button>
          </div>
          <div class="space-y-4">
            <div
              v-for="(field, index) in previewFields"
              :key="index"
              class="space-y-2"
            >
              <label class="block text-sm font-medium text-gray-700">
                {{ field.label || `字段 ${index + 1}` }}
                <span v-if="field.required" class="text-red-500">*</span>
              </label>
              <input
                v-if="field.type === 'text'"
                type="text"
                :placeholder="field.placeholder"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg"
              />
              <input
                v-else-if="field.type === 'number'"
                type="number"
                :placeholder="field.placeholder"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg"
              />
              <input
                v-else-if="field.type === 'date'"
                type="date"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg"
              />
              <select
                v-else-if="field.type === 'select'"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg"
              >
                <option v-for="option in getOptions(field.options)" :key="option">
                  {{ option }}
                </option>
              </select>
              <textarea
                v-else-if="field.type === 'textarea'"
                :placeholder="field.placeholder"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg"
                rows="3"
              ></textarea>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import FormBuilder from '../components/FormBuilder.vue'
import { createForm } from '../api/form'

const router = useRouter()
const formName = ref('')
const formDescription = ref('')
const formFields = ref([])
const showPreview = ref(false)
const previewFields = ref([])

function handlePreview(fields) {
  previewFields.value = fields
  showPreview.value = true
}

function getOptions(optionsString) {
  if (!optionsString) return []
  return optionsString.split('\n').filter(opt => opt.trim())
}

async function saveForm() {
  try {
    const formConfig = JSON.stringify(formFields.value)
    await createForm({
      name: formName.value,
      description: formDescription.value,
      form_config: formConfig,
    })
    router.push('/forms')
  } catch (error) {
    console.error('保存表单失败:', error)
    alert('保存失败，请重试')
  }
}
</script>

