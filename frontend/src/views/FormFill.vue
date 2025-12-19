<template>
  <div class="min-h-screen bg-gray-50">
    <div class="container mx-auto px-4 py-8 max-w-2xl">
      <div v-if="loading" class="text-center py-12">加载中...</div>

      <div v-else-if="form" class="bg-white rounded-lg shadow p-8">
        <h1 class="text-2xl font-bold text-gray-900 mb-2">{{ form.name }}</h1>
        <p v-if="form.description" class="text-gray-600 mb-6">{{ form.description }}</p>

        <form @submit.prevent="handleSubmit" class="space-y-6">
          <div
            v-for="(field, index) in formFields"
            :key="index"
            class="space-y-2"
          >
            <label class="block text-sm font-medium text-gray-700">
              {{ field.label || `字段 ${index + 1}` }}
              <span v-if="field.required" class="text-red-500">*</span>
            </label>
            <input
              v-if="field.type === 'text'"
              v-model="formData[field.id]"
              type="text"
              :required="field.required"
              :placeholder="field.placeholder"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
            />
            <input
              v-else-if="field.type === 'number'"
              v-model.number="formData[field.id]"
              type="number"
              :required="field.required"
              :placeholder="field.placeholder"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
            />
            <input
              v-else-if="field.type === 'date'"
              v-model="formData[field.id]"
              type="date"
              :required="field.required"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
            />
            <select
              v-else-if="field.type === 'select'"
              v-model="formData[field.id]"
              :required="field.required"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
            >
              <option value="">请选择</option>
              <option v-for="option in getOptions(field.options)" :key="option" :value="option">
                {{ option }}
              </option>
            </select>
            <textarea
              v-else-if="field.type === 'textarea'"
              v-model="formData[field.id]"
              :required="field.required"
              :placeholder="field.placeholder"
              rows="4"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
            ></textarea>
          </div>

          <div class="flex gap-4 pt-4">
            <button
              type="button"
              @click="$router.back()"
              class="flex-1 bg-gray-200 text-gray-700 py-2 px-4 rounded-lg font-medium hover:bg-gray-300 transition"
            >
              取消
            </button>
            <button
              type="submit"
              :disabled="submitting"
              class="flex-1 bg-blue-600 text-white py-2 px-4 rounded-lg font-medium hover:bg-blue-700 disabled:opacity-50 disabled:cursor-not-allowed transition"
            >
              <span v-if="submitting">提交中...</span>
              <span v-else>提交</span>
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { getForm, submitForm } from '../api/form'

const route = useRoute()
const router = useRouter()
const form = ref(null)
const loading = ref(true)
const submitting = ref(false)
const formData = ref({})

const formFields = computed(() => {
  if (!form.value || !form.value.form_config) return []
  try {
    return JSON.parse(form.value.form_config)
  } catch {
    return []
  }
})

function getOptions(optionsString) {
  if (!optionsString) return []
  return optionsString.split('\n').filter(opt => opt.trim())
}

async function loadForm() {
  try {
    const id = route.params.id
    const response = await getForm(id)
    form.value = response.data
    // 初始化表单数据
    formFields.value.forEach(field => {
      formData.value[field.id] = ''
    })
  } catch (error) {
    console.error('加载表单失败:', error)
  } finally {
    loading.value = false
  }
}

async function handleSubmit() {
  submitting.value = true
  try {
    await submitForm({
      form_id: form.value.id,
      data: JSON.stringify(formData.value),
    })
    alert('提交成功！')
    router.push('/forms')
  } catch (error) {
    console.error('提交失败:', error)
    alert('提交失败，请重试')
  } finally {
    submitting.value = false
  }
}

onMounted(() => {
  loadForm()
})
</script>

