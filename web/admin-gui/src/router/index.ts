import { createRouter, createWebHistory } from 'vue-router'
import Home from '@/views/Home.vue'
import Debug from '@/views/Debug.vue'
import Login from '@/views/Login.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'login',
      component: Login
    },
    {
      path: '/debug',
      name: 'debug',
      component: Debug
    },
    {
      path: '/home',
      name: 'home',
      component: Home
    },
  ]
})

export default router
