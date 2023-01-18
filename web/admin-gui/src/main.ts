import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'

import '@/assets/base.css'
import '@/assets/tabulator_bootstrap5.min.css'

const app = createApp(App)

app.use(createPinia())
app.use(router)

app.mount('#app')
