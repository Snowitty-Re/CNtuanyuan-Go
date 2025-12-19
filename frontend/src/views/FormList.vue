<template>
  <Layout>
      <div class="flex justify-between items-center mb-6">
        <h1 class="text-3xl font-bold text-gray-900">表单管理</h1>
        <router-link
          to="/forms/builder"
          class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition"
        >
          创建表单
        </router-link>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div
          v-for="form in forms"
          :key="form.id"
          class="bg-white rounded-lg shadow p-6 hover:shadow-lg transition"
        >
          <h3 class="text-xl font-semibold text-gray-900 mb-2">{{ form.name }}</h3>
          <p class="text-gray-600 text-sm mb-4">{{ form.description || '无描述' }}</p>
          <div class="flex gap-2">
            <button
              @click="viewForm(form)"
              class="flex-1 bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition text-sm"
            >
              查看
            </button>
            <button
              @click="fillForm(form)"
              class="flex-1 bg-green-600 text-white px-4 py-2 rounded-lg hover:bg-green-700 transition text-sm"
            >
              填写
            </button>
          </div>
        </div>
      </div>

      <div v-if="forms.length === 0" class="text-center py-12 text-gray-500">
        暂无表单，点击上方按钮创建一个
      </div>
  </Layout>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import Layout from '../components/Layout.vue'
import { getForms } from '../api/form'

const router = useRouter()
const forms = ref([])

async function loadForms() {
  try {
    const response = await getForms()
    forms.value = response.data.forms || []
  } catch (error) {
    console.error('加载表单失败:', error)
  }
}

function viewForm(form) {
  router.push(`/forms/${form.id}`)
}

function fillForm(form) {
  router.push(`/forms/${form.id}/fill`)
}

onMounted(() => {
  loadForms()
})
</script>

