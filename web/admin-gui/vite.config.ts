import { fileURLToPath, URL } from 'node:url'
import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig(({command, mode}) => {
	return { 
		plugins: [vue()],
		resolve: {
			alias: {
				'@': fileURLToPath(new URL('./src', import.meta.url))
			}
		},
		define: {
			__APP_ENV__: {
				...loadEnv(mode, "./../../", ""), // when running directly | access .env from ./web/admin-gui/ 
				...loadEnv(mode, process.cwd(), ""), // when running from docker
			}
		}
	}
})
