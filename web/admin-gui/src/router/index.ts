import { createRouter, createWebHistory, type NavigationGuard } from 'vue-router'
import Home from '@/views/Home.vue'
import Debug from '@/views/Debug.vue'
import Login from '@/views/Login.vue'
import { useAuthStore } from '@/stores/AuthStore'
import { nextTick } from 'vue'

const router = createRouter({
	history: createWebHistory(import.meta.env.BASE_URL),
	routes: [
		{
			path: '/',
			name: 'login',
			component: Login,
			meta: {
				title: 'GohCMS - Connexion'
			}
		},
		{
			path: '/debug',
			name: 'debug',
			component: Debug,
			meta: {
				title: 'GohCMS - Debug'
			}
		},
		{
			path: '/home',
			name: 'home',
			component: Home,
			meta: {
				title: 'GohCMS - Accueil'
			}
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

router.afterEach((to, from) => {
	nextTick(() => {
		document.title = to.meta.title as string ?? 'GohCMS'
	})
})

export default router
