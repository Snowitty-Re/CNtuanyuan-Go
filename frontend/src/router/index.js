import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '../stores/auth'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'home',
      component: () => import('../views/Home.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/Login.vue'),
      meta: { requiresAuth: false },
    },
    {
      path: '/missing-persons',
      name: 'missingPersonList',
      component: () => import('../views/MissingPersonList.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/workflows',
      name: 'workflowList',
      component: () => import('../views/WorkflowList.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/workflows/:id',
      name: 'workflowDetail',
      component: () => import('../views/WorkflowDetail.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/forms',
      name: 'formList',
      component: () => import('../views/FormList.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/forms/builder',
      name: 'formBuilder',
      component: () => import('../views/FormBuilder.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/forms/:id',
      name: 'formDetail',
      component: () => import('../views/FormFill.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/forms/:id/fill',
      name: 'formFill',
      component: () => import('../views/FormFill.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/dashboard',
      name: 'dashboard',
      component: () => import('../views/Dashboard.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/users',
      name: 'userManagement',
      component: () => import('../views/UserManagement.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/roles-permissions',
      name: 'rolePermissionManagement',
      component: () => import('../views/RolePermissionManagement.vue'),
      meta: { requiresAuth: true },
    },
  ],
})

// 路由守卫
router.beforeEach(async (to, from, next) => {
  const authStore = useAuthStore()

  // 如果需要认证但未认证，跳转到登录页
  if (to.meta.requiresAuth && !authStore.isAuthenticated) {
    next({ name: 'login', query: { redirect: to.fullPath } })
    return
  }

  // 如果已登录但访问登录页，跳转到首页
  if (to.name === 'login' && authStore.isAuthenticated) {
    next({ name: 'home' })
    return
  }

  // 如果需要认证且有token但用户信息未加载，尝试加载用户信息
  if (to.meta.requiresAuth && authStore.isAuthenticated && !authStore.user) {
    try {
      await authStore.fetchUser()
    } catch (error) {
      // 如果获取用户信息失败，清除token并跳转到登录页
      console.error('获取用户信息失败:', error)
      authStore.logout()
      next({ name: 'login', query: { redirect: to.fullPath } })
      return
    }
  }

  next()
})

export default router
