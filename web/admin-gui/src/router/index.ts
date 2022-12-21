import { createRouter, createWebHistory } from 'vue-router'
import Home from '@/views/Home.vue'
import Debug from '@/views/Debug.vue'
import Login from '@/views/Login.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/debug',
      name: 'debug',
      component: Debug
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    },
  ]
})

export default router
