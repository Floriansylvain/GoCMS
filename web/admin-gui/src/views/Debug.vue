<script setup lang="ts">
// import EditorJS from '@editorjs/editorjs';
import { useAuthStore, type jwtFormat } from '@/stores/AuthStore';
import { ref } from 'vue';

const email = ref('')
const password = ref('')

function getArticles() {
	fetch(`http://localhost:${__APP_ENV__.API_PORT}/articles/`, {
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
	fetch(`http://localhost:${__APP_ENV__.API_PORT}/login/`, {
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
			<button @click="getArticles()">get all articles</button>
			<button @click="getArticles()">get all articles</button>
			<button @click="getArticles()">get all articles</button>
			<button @click="getArticles()">get all articles</button>
		</div>
	</div>
</template>
