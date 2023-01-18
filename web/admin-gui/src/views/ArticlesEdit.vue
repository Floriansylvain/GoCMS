<script setup lang="ts">
import Modal from '@/components/ModalSuccessError.vue';
import { editArticle, getArticles, type Article } from '@/utils/database';
import Editor from '@tinymce/tinymce-vue';
import { onMounted, ref, type Ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';

const article: Ref<Article | void> = ref()
const editorData: Ref<string> = ref('')
const router = useRouter()

const successModalShow = ref(false)
const errorModalShow = ref(false)
const errorModalDescription: Ref<string> = ref("Quelque chose s'est mal passé...")

onMounted(async () => {
	await getArticles(useRoute().params.articleID as string)
		.then(articleFetch => {
			article.value = articleFetch[0]
			editorData.value = article.value.content.html
		})
		.catch(error => {
			errorModalDescription.value = "Impossible de récupérer l'article. Vérifiez l'URL. Tentez-vous d'accéder au mode édition directement depuis un lien ?"
			errorModalShow.value = true
			console.error(error)
		})
})

function abort() {
	router.push('/articles')
}

async function saveContent() {
	if (article.value == undefined) return;
	article.value.content.html = editorData.value
	await editArticle(article.value)
		.then(() => {
			successModalShow.value = true
		})
		.catch(error => {
			errorModalDescription.value = `Impossible de sauvegarder l'article. (${error.toString()})`
			errorModalShow.value = true
			console.error(error)
		})

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
			<button @click="abort()" class="button-secondary">Retour</button>
		</aside>
	</div>
	<Modal v-if="errorModalShow" :description="errorModalDescription" @close="errorModalShow = false" type="error">
	</Modal>
	<Modal v-if="successModalShow" description="Le contenu a bien été sauvegardé." @close="successModalShow = false"
		type="success">
	</Modal>
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