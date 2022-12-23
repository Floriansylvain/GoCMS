import { createRouter, createWebHistory, type NavigationGuard } from 'vue-router'
import Home from '@/views/Home.vue'
import Debug from '@/views/Debug.vue'
import Login from '@/views/Login.vue'
import { useAuthStore } from '@/stores/AuthStore'

const router = createRouter({
	history: createWebHistory(import.meta.env.BASE_URL),
	routes: [
		{
			path: '/',
			name: 'login',
			component: Login,
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

router.beforeEach(async (to, from) => {
	const isTokenValid = useAuthStore().isValid()
	if (to.name === 'login') {
		if (isTokenValid) return {
			name: 'home'
		}
	} else {
		if (!isTokenValid) { 
			return {
				name: 'login',
			}
		}
	}
})

export default router
