<template>
  <Layout>
      <h1 class="text-3xl font-bold text-gray-900 mb-8">角色权限管理</h1>

      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- 角色管理 -->
        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex justify-between items-center mb-4">
            <h2 class="text-xl font-semibold text-gray-900">角色管理</h2>
            <button
              @click="showCreateRoleModal = true"
              class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition text-sm"
            >
              创建角色
            </button>
          </div>

          <div class="space-y-3">
            <div
              v-for="role in roles"
              :key="role.id"
              class="p-4 border border-gray-200 rounded-lg hover:border-blue-500 transition"
            >
              <div class="flex justify-between items-start">
                <div>
                  <h3 class="font-medium text-gray-900">{{ role.name }}</h3>
                  <p class="text-sm text-gray-600 mt-1">{{ role.description || '无描述' }}</p>
                  <div class="mt-2">
                    <span
                      v-for="permission in role.permissions"
                      :key="permission.id"
                      class="mr-2 mb-1 inline-block px-2 py-1 bg-gray-100 text-gray-700 text-xs rounded"
                    >
                      {{ permission.name }}
                    </span>
                  </div>
                </div>
                <button
                  @click="editRole(role)"
                  class="text-blue-600 hover:text-blue-800 text-sm"
                >
                  编辑
                </button>
              </div>
            </div>
          </div>

          <div v-if="roles.length === 0" class="text-center py-8 text-gray-500 text-sm">
            暂无角色
          </div>
        </div>

        <!-- 权限管理 -->
        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex justify-between items-center mb-4">
            <h2 class="text-xl font-semibold text-gray-900">权限管理</h2>
            <button
              @click="showCreatePermissionModal = true"
              class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition text-sm"
            >
              创建权限
            </button>
          </div>

          <div class="space-y-3">
            <div
              v-for="permission in permissions"
              :key="permission.id"
              class="p-4 border border-gray-200 rounded-lg"
            >
              <h3 class="font-medium text-gray-900">{{ permission.name }}</h3>
              <p class="text-sm text-gray-600 mt-1">
                {{ permission.resource }} / {{ permission.action }}
              </p>
            </div>
          </div>

          <div v-if="permissions.length === 0" class="text-center py-8 text-gray-500 text-sm">
            暂无权限
          </div>
        </div>
      </div>
  </Layout>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import Layout from '../components/Layout.vue'
import { getRoles, getPermissions } from '../api/permission'

const roles = ref([])
const permissions = ref([])
const showCreateRoleModal = ref(false)
const showCreatePermissionModal = ref(false)

async function loadData() {
  try {
    const [rolesRes, permissionsRes] = await Promise.all([
      getRoles(),
      getPermissions(),
    ])
    roles.value = rolesRes.data?.roles || []
    permissions.value = permissionsRes.data?.permissions || []
  } catch (error) {
    console.error('加载数据失败:', error)
  }
}

function editRole(role) {
  // 编辑角色功能
  console.log('编辑角色:', role)
}

onMounted(() => {
  loadData()
})
</script>

