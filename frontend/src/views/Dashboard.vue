<template>
  <Layout>
      <h1 class="text-3xl font-bold text-gray-900 mb-8">仪表盘</h1>

      <!-- 统计卡片 -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm text-gray-600 mb-1">走失人员总数</p>
              <p class="text-2xl font-bold text-gray-900">{{ stats.missingPersons || 0 }}</p>
            </div>
            <div class="bg-blue-100 rounded-full p-3">
              <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
              </svg>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm text-gray-600 mb-1">工作流总数</p>
              <p class="text-2xl font-bold text-gray-900">{{ stats.workflows || 0 }}</p>
            </div>
            <div class="bg-green-100 rounded-full p-3">
              <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
              </svg>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm text-gray-600 mb-1">表单总数</p>
              <p class="text-2xl font-bold text-gray-900">{{ stats.forms || 0 }}</p>
            </div>
            <div class="bg-purple-100 rounded-full p-3">
              <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
              </svg>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm text-gray-600 mb-1">用户总数</p>
              <p class="text-2xl font-bold text-gray-900">{{ stats.users || 0 }}</p>
            </div>
            <div class="bg-yellow-100 rounded-full p-3">
              <svg class="w-6 h-6 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z" />
              </svg>
            </div>
          </div>
        </div>
      </div>

      <!-- 快速操作 -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <div class="bg-white rounded-lg shadow p-6">
          <h2 class="text-xl font-semibold text-gray-900 mb-4">快速操作</h2>
          <div class="space-y-3">
            <router-link
              to="/missing-persons"
              class="block p-3 bg-gray-50 hover:bg-gray-100 rounded-lg transition"
            >
              <span class="font-medium text-gray-900">走失人员管理</span>
            </router-link>
            <router-link
              to="/workflows"
              class="block p-3 bg-gray-50 hover:bg-gray-100 rounded-lg transition"
            >
              <span class="font-medium text-gray-900">工作流管理</span>
            </router-link>
            <router-link
              to="/forms"
              class="block p-3 bg-gray-50 hover:bg-gray-100 rounded-lg transition"
            >
              <span class="font-medium text-gray-900">表单管理</span>
            </router-link>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <h2 class="text-xl font-semibold text-gray-900 mb-4">最近活动</h2>
          <div class="text-gray-500 text-sm">
            暂无最近活动
          </div>
        </div>
      </div>
  </Layout>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import Layout from '../components/Layout.vue'
import { getMissingPersons } from '../api/missingPerson'
import { getWorkflows } from '../api/workflow'
import { getForms } from '../api/form'

const stats = ref({
  missingPersons: 0,
  workflows: 0,
  forms: 0,
  users: 0,
})

async function loadStats() {
  try {
    const [missingPersonsRes, workflowsRes, formsRes] = await Promise.all([
      getMissingPersons({ limit: 1 }),
      getWorkflows(),
      getForms(),
    ])
    stats.value.missingPersons = missingPersonsRes.data?.total || 0
    stats.value.workflows = workflowsRes.data?.total || 0
    stats.value.forms = formsRes.data?.total || 0
  } catch (error) {
    console.error('加载统计失败:', error)
  }
}

onMounted(() => {
  loadStats()
})
</script>

