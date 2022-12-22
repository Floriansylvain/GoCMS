<script setup lang="ts">
// import EditorJS from '@editorjs/editorjs';
import { useAuthStore, type jwtFormat } from '@/stores/AuthStore';
import { ref } from 'vue';

const email = ref('')
const password = ref('')

const baseURL = `http://${__APP_ENV__.APP_HOST_ADDRESS}:${__APP_ENV__.APP_API_PORT}`

function getArticles() {
	fetch(`${baseURL}/articles/`, {
		headers: { "Authorization": `Bearer ${useAuthStore().token}` }
	})
	.then(response => response.json())
	.then(result => console.log(result))
}

function ping() {
	fetch(`${baseURL}/ping/`, {
		headers: { "Authorization": `Bearer ${useAuthStore().token}` }
	})
	.then(response => response.json())
	.then(result => console.log(result))
}

function jwtHandler(apiResponse: jwtFormat): void {
	if (apiResponse.code !== 200) {
		console.log(apiResponse)
		return
	}
	useAuthStore().token = apiResponse.token
	useAuthStore().expire = apiResponse.expire
	console.log('success! logged in.')
}

function login(email: string, password: string): void {
	fetch(`${baseURL}/login/`, {
		method: "POST",
		body: JSON.stringify({
			email: email,
			password: password
		})
	})
	.then(response => response.json())
	.then(result => jwtHandler(result))
}
</script>

<template>
	<h2>Connexion</h2>
	
	<form @submit.prevent="login(email, password)">
		<div>
			<label for="email">Email address</label>
			<input id="email" name="email" placeholder="email" type="text" v-model="email">
		</div>
		<div>
			<label for="password">Password</label>
			<input id="password" name="password" placeholder="password" type="password" v-model="password">
		</div>
		<button type="submit">Se connecter</button>
	</form>
	
	<div>
		<div>
			<button class="button-primary" @click="getArticles()">get all articles</button>
			<button class="button-primary" @click="ping()">ping</button>
		</div>
	</div>
</template>
