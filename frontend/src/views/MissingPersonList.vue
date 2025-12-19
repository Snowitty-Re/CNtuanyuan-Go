<template>
  <div class="min-h-screen bg-gray-50">
    <div class="container mx-auto px-4 py-8">
      <div class="flex justify-between items-center mb-6">
        <h1 class="text-3xl font-bold text-gray-900">走失人员管理</h1>
        <router-link
          to="/missing-persons/new"
          class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition"
        >
          新增走失人员
        </router-link>
      </div>

      <!-- 搜索和筛选 -->
      <div class="bg-white rounded-lg shadow p-6 mb-6">
        <div class="grid grid-cols-1 md:grid-cols-4 gap-4">
          <input
            v-model="filters.name"
            type="text"
            placeholder="搜索姓名"
            class="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
            @input="handleSearch"
          />
          <select
            v-model="filters.gender"
            class="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
            @change="handleSearch"
          >
            <option value="">全部性别</option>
            <option value="male">男</option>
            <option value="female">女</option>
            <option value="unknown">未知</option>
          </select>
          <select
            v-model="filters.status"
            class="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
            @change="handleSearch"
          >
            <option value="">全部状态</option>
            <option value="missing">走失中</option>
            <option value="found">已找到</option>
            <option value="confirmed">已确认</option>
          </select>
          <button
            @click="handleSearch"
            class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition"
          >
            搜索
          </button>
        </div>
      </div>

      <!-- 列表 -->
      <div class="bg-white rounded-lg shadow overflow-hidden">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                姓名
              </th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                性别
              </th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                年龄
              </th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                走失地点
              </th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                状态
              </th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                操作
              </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="person in persons" :key="person.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                {{ person.name }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ person.gender === 'male' ? '男' : person.gender === 'female' ? '女' : '未知' }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ person.age || '-' }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ person.missing_location || '-' }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span
                  :class="{
                    'bg-yellow-100 text-yellow-800': person.status === 'missing',
                    'bg-green-100 text-green-800': person.status === 'found',
                    'bg-blue-100 text-blue-800': person.status === 'confirmed',
                  }"
                  class="px-2 py-1 text-xs font-semibold rounded-full"
                >
                  {{ person.status === 'missing' ? '走失中' : person.status === 'found' ? '已找到' : '已确认' }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <router-link
                  :to="`/missing-persons/${person.id}`"
                  class="text-blue-600 hover:text-blue-900 mr-4"
                >
                  查看
                </router-link>
                <button
                  @click="handleDelete(person.id)"
                  class="text-red-600 hover:text-red-900"
                >
                  删除
                </button>
              </td>
            </tr>
          </tbody>
        </table>

        <div v-if="persons.length === 0" class="text-center py-12 text-gray-500">
          暂无数据
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getMissingPersons, deleteMissingPerson } from '../api/missingPerson'
import { useRouter } from 'vue-router'

const router = useRouter()
const persons = ref([])
const filters = ref({
  name: '',
  gender: '',
  status: '',
})

async function loadData() {
  try {
    const response = await getMissingPersons(filters.value)
    persons.value = response.data.persons || []
  } catch (error) {
    console.error('加载数据失败:', error)
  }
}

function handleSearch() {
  loadData()
}

async function handleDelete(id) {
  if (confirm('确定要删除这条记录吗？')) {
    try {
      await deleteMissingPerson(id)
      loadData()
    } catch (error) {
      console.error('删除失败:', error)
    }
  }
}

onMounted(() => {
  loadData()
})
</script>

