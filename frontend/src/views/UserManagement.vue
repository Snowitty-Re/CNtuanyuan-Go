<template>
  <Layout>
      <div class="flex justify-between items-center mb-6">
        <h1 class="text-3xl font-bold text-gray-900">用户管理</h1>
        <button
          @click="showCreateModal = true"
          class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition"
        >
          添加用户
        </button>
      </div>

      <div class="bg-white rounded-lg shadow overflow-hidden">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                ID
              </th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                用户名
              </th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                邮箱
              </th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                角色
              </th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                创建时间
              </th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                操作
              </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="user in users" :key="user.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ user.id }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                {{ user.username }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ user.email }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                <span
                  v-for="role in user.roles"
                  :key="role.id"
                  class="mr-2 px-2 py-1 bg-blue-100 text-blue-800 text-xs rounded-full"
                >
                  {{ role.name }}
                </span>
                <span v-if="!user.roles || user.roles.length === 0" class="text-gray-400">
                  无角色
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ formatDate(user.created_at) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button
                  @click="editUser(user)"
                  class="text-blue-600 hover:text-blue-900 mr-4"
                >
                  编辑
                </button>
                <button
                  @click="deleteUser(user.id)"
                  class="text-red-600 hover:text-red-900"
                >
                  删除
                </button>
              </td>
            </tr>
          </tbody>
        </table>

        <div v-if="users.length === 0" class="text-center py-12 text-gray-500">
          暂无用户
        </div>
      </div>
  </Layout>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import Layout from '../components/Layout.vue'

const users = ref([])
const showCreateModal = ref(false)

function formatDate(dateString) {
  if (!dateString) return '-'
  return new Date(dateString).toLocaleDateString('zh-CN')
}

function editUser(user) {
  // 编辑用户功能
  console.log('编辑用户:', user)
}

async function deleteUser(id) {
  if (confirm('确定要删除这个用户吗？')) {
    // 删除用户功能
    console.log('删除用户:', id)
  }
}

onMounted(() => {
  // 加载用户列表
  // 这里需要实现用户列表API
})
</script>

