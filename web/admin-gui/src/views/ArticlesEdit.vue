<script setup lang="ts">
import { getArticles, type Article } from '@/utils/database';
import Editor from '@tinymce/tinymce-vue';
import { onMounted, ref, type Ref } from 'vue';
import { useRoute } from 'vue-router';

const article: Ref<Article | void> = ref()
const editorData: Ref<string> = ref('')

onMounted(async () => {
	const articleFetch = await getArticles(useRoute().params.articleID as string)
	article.value = articleFetch[0]
	editorData.value = article.value.content.html
})

function abort() {
}

function saveContent() {
	console.log(editorData.value)
}
</script>
	
<template>
	<div class="container">
		<div id="editor">
			<Editor tinymce-script-src="/tinymce/tinymce.min.js"
				:init="{ promotion: false, language: 'fr_FR', resize: false, height: '100%', }"
				:plugins="['link', 'codesample']"
				toolbar="undo redo | styles | bold italic underline strikethrough | alignleft aligncenter alignright alignjustify | outdent indent | codesample link"
				v-model="editorData">
			</Editor>
		</div>
		<aside class="buttons">
			<button @click="saveContent()" class="button-primary">Enregistrer</button>
			<button @click="abort()" class="button-secondary">Annuler</button>
		</aside>
	</div>
</template>

<style scoped>
.container {
	display: flex;
	justify-content: center;
	width: 100%;
	height: 100%;
	overflow: hidden;
}

#editor {
	padding-top: 16px;
	width: 100%;
	height: 100%;
}

.tox-tinymce {
	border-radius: 0 !important;
}

.buttons {
	position: relative;
	z-index: 1;

	display: flex;
	flex-direction: column;
	justify-content: right;
	align-content: center;
	gap: 16px;

	width: fit-content;
	padding: 32px 16px;

	box-shadow: #0002 0 0 10px;
}
</style>