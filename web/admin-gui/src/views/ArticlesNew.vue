<script setup lang="ts">
import { postArticle, type Article } from '@/utils/database';
import { ref, type Ref } from 'vue';
import { RouterLink, useRouter } from 'vue-router'

const router = useRouter()

const defaultData = `<h1>Bienvenue</h1><p>Vous &ecirc;tes en mode <em>&eacute;dition</em> d'article.</p><p>Celui-ci semble encore neuf ! Supprimez ces lignes et laissez libre cours &agrave; votre imagination :)</p><p>Pour plus d'infos, rendez-vous sur la <a title="Attention, rickroll incoming" href="https:/www.youtube.com/watch?v=dQw4w9WgXcQ" target="_blank" rel="noopener">page d'aide</a>.</p>`
const title = ref('')
const page = ref('')

function createArticle() {
	if (!isFormValid()) {
		alert("L'identifiant de l'article ne doit contenir que des caractères alphanumériques (aA-zZ, 1-9) !")
		return;
	}

	const article: Article = {
		idName: title.value,
		date: new Date().getTime(),
		content: {
			html: defaultData
		},
		online: false,
		pageId: page.value
	}
	postArticle(article)
	router.push(`/articles/edit/${article.idName}`)
}

const isFormEmpty: () => boolean = () => {
	return title.value === '' || page.value === ''
}

function isFormValid(): boolean {
	const regex = /^[a-zA-Z0-9]+$/
	return regex.test(title.value)
}
</script>

<template>
	<main>
		<h2>Créer un nouvel article</h2>
		<form @submit.prevent="createArticle()">
			<div class="inputs-group">
				<div class="label-input">
					<label for="title">Identifiant de l'article</label>
					<input id="title" name="title" placeholder="Identifiant de l'article" type="text" v-model="title">
					<p>⚠️ Doit être unique et composé de caractères alphanumériques !</p>
					<!-- TODO: Remplacer alerte pas ouf par passage du champ en erreur -->
				</div>
				<div class="label-input">
					<label for="page">Page</label>
					<input id="page" name="page" placeholder="Page" type="text" v-model="page">
				</div>
			</div>
			<div class="buttons-group">
				<RouterLink class="button-secondary" type="submit" to="/articles">Annuler</RouterLink>
				<button :class="`button-${isFormEmpty() ? 'disabled' : 'primary'}`" type="submit"
					:disabled="isFormEmpty()">Créer</button>
			</div>
		</form>
	</main>
</template>

<style scoped>
main {
	padding: 0 64px;
	max-width: 100%;
	width: 560px;
}

.buttons-group {
	display: flex;
	gap: 8px;
	width: 100%;
	margin: 32px 0;
}

.buttons-group a,
.buttons-group button {
	width: 100%;
}

.inputs-group {
	display: flex;
	flex-direction: column;
	gap: 8px;
}
</style>