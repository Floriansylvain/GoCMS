<script setup lang="ts">
import Editor from '@/components/Editor.vue';
import { useAuthStore } from '@/stores/AuthStore';
import { ref } from 'vue';

const baseURL = `http://${__APP_ENV__.APP_HOST_ADDRESS}:${__APP_ENV__.APP_API_PORT}`
const articleID = ref('')

function getArticles(id?: string) {
	fetch(`${baseURL}/articles/${id ?? ''}`, {
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
</script>

<template>
	<div class="inputs-group">
		<div>
			<input type="text" placeholder="article ID" v-model="articleID">
			<button class="button-primary" @click="getArticles(articleID)">get article</button>
		</div>
	</div>
	<div class="buttons-group">
		<button class="button-primary" @click="ping()">ping</button>
	</div>
	<div class="editor">
		<Editor :show-default-data="true"></Editor>
	</div>
</template>

<style scoped>
.buttons-group,
.inputs-group {
	display: flex;
	flex-wrap: wrap;
	gap: 8px;
	margin: 32px;
}

.inputs-group {
	flex-direction: column;
}

.inputs-group div {
	flex-direction: row;
	gap: 8px;
}

.editor {
	border: solid black 2px;
}
</style>