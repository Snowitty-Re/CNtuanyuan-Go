<template>
  <div class="form-builder">
    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <!-- 字段库 -->
      <div class="bg-white rounded-lg shadow p-4">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">字段库</h3>
        <div class="space-y-2">
          <button
            v-for="fieldType in fieldTypes"
            :key="fieldType.type"
            @click="addField(fieldType)"
            class="w-full text-left px-4 py-2 bg-gray-50 hover:bg-gray-100 rounded-lg transition"
          >
            <span class="font-medium text-gray-900">{{ fieldType.label }}</span>
            <span class="text-sm text-gray-500 ml-2">{{ fieldType.type }}</span>
          </button>
        </div>
      </div>

      <!-- 表单设计区 -->
      <div class="lg:col-span-2 bg-white rounded-lg shadow p-4">
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-lg font-semibold text-gray-900">表单设计</h3>
          <button
            @click="previewForm"
            class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition text-sm"
          >
            预览
          </button>
        </div>

        <div v-if="fields.length === 0" class="text-center py-12 text-gray-500 border-2 border-dashed rounded-lg">
          从左侧拖拽字段到此处开始构建表单
        </div>

        <div v-else class="space-y-4">
          <div
            v-for="(field, index) in fields"
            :key="field.id"
            class="border border-gray-200 rounded-lg p-4 hover:border-blue-500 transition"
          >
            <div class="flex justify-between items-start mb-2">
              <div class="flex-1">
                <label class="block text-sm font-medium text-gray-700 mb-1">
                  {{ field.label || `字段 ${index + 1}` }}
                  <span v-if="field.required" class="text-red-500">*</span>
                </label>
                <input
                  v-if="field.type === 'text'"
                  type="text"
                  :placeholder="field.placeholder || '请输入'"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg"
                  disabled
                />
                <input
                  v-else-if="field.type === 'number'"
                  type="number"
                  :placeholder="field.placeholder || '请输入数字'"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg"
                  disabled
                />
                <input
                  v-else-if="field.type === 'date'"
                  type="date"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg"
                  disabled
                />
                <select
                  v-else-if="field.type === 'select'"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg"
                  disabled
                >
                  <option>请选择</option>
                  <option v-for="option in getOptions(field.options)" :key="option">
                    {{ option }}
                  </option>
                </select>
                <textarea
                  v-else-if="field.type === 'textarea'"
                  :placeholder="field.placeholder || '请输入'"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg"
                  rows="3"
                  disabled
                ></textarea>
              </div>
              <div class="ml-4 flex gap-2">
                <button
                  @click="editField(index)"
                  class="text-blue-600 hover:text-blue-800 text-sm"
                >
                  编辑
                </button>
                <button
                  @click="removeField(index)"
                  class="text-red-600 hover:text-red-800 text-sm"
                >
                  删除
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 字段编辑模态框 -->
    <div
      v-if="editingFieldIndex !== null"
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 px-4"
      @click.self="editingFieldIndex = null"
    >
      <div class="bg-white rounded-xl shadow-xl p-8 max-w-md w-full">
        <h3 class="text-xl font-bold text-gray-900 mb-6">编辑字段</h3>
        <form @submit.prevent="saveField" class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">字段标签</label>
            <input
              v-model="editingField.label"
              type="text"
              required
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">占位符</label>
            <input
              v-model="editingField.placeholder"
              type="text"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
            />
          </div>
          <div class="flex items-center">
            <input
              v-model="editingField.required"
              type="checkbox"
              class="mr-2"
            />
            <label class="text-sm text-gray-700">必填</label>
          </div>
          <div v-if="editingField.type === 'select'">
            <label class="block text-sm font-medium text-gray-700 mb-2">选项（每行一个）</label>
            <textarea
              v-model="editingField.options"
              rows="4"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
            ></textarea>
          </div>
          <div class="flex gap-3 pt-4">
            <button
              type="button"
              @click="editingFieldIndex = null"
              class="flex-1 bg-gray-200 text-gray-700 py-2 px-4 rounded-lg font-medium hover:bg-gray-300 transition"
            >
              取消
            </button>
            <button
              type="submit"
              class="flex-1 bg-blue-600 text-white py-2 px-4 rounded-lg font-medium hover:bg-blue-700 transition"
            >
              保存
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'

const props = defineProps({
  modelValue: {
    type: Array,
    default: () => [],
  },
})

const emit = defineEmits(['update:modelValue'])

const fields = computed({
  get: () => props.modelValue,
  set: (value) => emit('update:modelValue', value),
})

const editingFieldIndex = ref(null)
const editingField = ref({})

const fieldTypes = [
  { type: 'text', label: '文本输入' },
  { type: 'number', label: '数字输入' },
  { type: 'date', label: '日期选择' },
  { type: 'select', label: '下拉选择' },
  { type: 'textarea', label: '多行文本' },
]

let fieldIdCounter = 0

function addField(fieldType) {
  const newField = {
    id: ++fieldIdCounter,
    type: fieldType.type,
    label: '',
    placeholder: '',
    required: false,
    options: '',
  }
  fields.value.push(newField)
  editField(fields.value.length - 1)
}

function editField(index) {
  editingFieldIndex.value = index
  editingField.value = { ...fields.value[index] }
}

function saveField() {
  if (editingFieldIndex.value !== null) {
    fields.value[editingFieldIndex.value] = { ...editingField.value }
    editingFieldIndex.value = null
  }
}

function removeField(index) {
  fields.value.splice(index, 1)
}

function previewForm() {
  // 预览功能可以在父组件中实现
  emit('preview', fields.value)
}

function getOptions(optionsString) {
  if (!optionsString) return []
  return optionsString.split('\n').filter(opt => opt.trim())
}
</script>

