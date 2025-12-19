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
  ],
})

// 路由守卫
router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()

  if (to.meta.requiresAuth && !authStore.isAuthenticated) {
    next({ name: 'login', query: { redirect: to.fullPath } })
  } else if (to.name === 'login' && authStore.isAuthenticated) {
    next({ name: 'home' })
  } else {
    next()
  }
})

export default router
