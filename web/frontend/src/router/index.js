import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      redirect: '/dashboard',
    },
    {
      path: '/dashboard',
      component: () => import('../layouts/AppLayout.vue'),
      children: [
        {
          path: '',
          name: 'dashboard',
          component: () => import('../views/DashboardView.vue'),
          meta: { title: '工作台' },
        },
        {
          path: 'calendar',
          name: 'calendar',
          component: () => import('../views/PlaceholderView.vue'),
          meta: { title: '日程管理', description: '这里将承载日程、待办和提醒功能。' },
        },
        {
          path: 'notes',
          name: 'notes',
          component: () => import('../views/PlaceholderView.vue'),
          meta: { title: '内容收集', description: '这里将承载笔记、链接和灵感收集功能。' },
        },
        {
          path: 'settings',
          name: 'settings',
          component: () => import('../views/PlaceholderView.vue'),
          meta: { title: '设置', description: '这里将承载主题、数据和应用偏好设置。' },
        },
      ],
    },
    {
      path: '/:pathMatch(.*)*',
      redirect: '/dashboard',
    },
  ],
})

export default router
