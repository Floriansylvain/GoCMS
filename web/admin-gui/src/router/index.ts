import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/stores/AuthStore'
import { nextTick } from 'vue'
import Debug from '@/views/Debug.vue'
import Login from '@/views/Login.vue'
import Home from '@/views/Home.vue'
import Articles from '@/views/Articles.vue'
import ArticlesEdit from '@/views/ArticlesEdit.vue'
import ArticlesNew from '@/views/ArticlesNew.vue'

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
		{
			path: '/articles',
			name: 'articles',
			component: Articles,
			meta: {
				title: 'GohCMS - Articles'
			}
		},
		{
			path: '/articles/edit/:articleID',
			name: 'edition',
			component: ArticlesEdit,
			meta: {
				title: 'GohCMS - Edition'
			}
		},
		{
			path: '/articles/new',
			name: 'edition',
			component: ArticlesNew,
			meta: {
				title: 'GohCMS - Edition'
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
