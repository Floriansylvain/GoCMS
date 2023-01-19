<script setup lang="ts">
import { postArticle, type Article } from '@/utils/database';
import { ref, type Ref } from 'vue';
import { RouterLink, useRouter } from 'vue-router'

const router = useRouter()

const defaultData = `<h1>Bienvenue</h1><p>Vous &ecirc;tes en mode <em>&eacute;dition</em> d'article.</p><p>Celui-ci semble encore neuf ! Supprimez ces lignes et laissez libre cours &agrave; votre imagination :)</p><p>Pour plus d'infos, rendez-vous sur la <a title="Attention, rickroll incoming" href="https:/www.youtube.com/watch?v=dQw4w9WgXcQ" target="_blank" rel="noopener">page d'aide</a>.</p>`
const title = ref('')
const rawTags = ref('')
const tags: Ref<string[]> = ref([])

const regexAccents = /[\u0300-\u036f]/g
const regexSymbols = /([-!$%^&*()_+|~=`{}\[\]:";'<>?,.\/])+/g

function updateUniqueTags(): void {
	tags.value = rawTags.value.split(/[,\s]+/g)
	tags.value = Array.from(new Set(tags.value.filter(x => x !== '')))
}

function generateTitleID(title: string): string {
	return title.toLowerCase().normalize('NFD')
		.replace(regexAccents, '')
		.replace(regexSymbols, '')
		.replace(/\s/g, '-')
}

function isFormEmpty(): boolean {
	return title.value === '' || rawTags.value === ''
}

async function createArticle() {
	const article: Article = {
		titleID: generateTitleID(title.value),
		title: title.value,
		date: new Date().getTime(),
		content: {
			html: defaultData
		},
		online: false,
		tags: tags.value
	}
	await postArticle(article)
	router.push(`/articles/edit/${article.titleID}`)
}
</script>

<template>
	<main>
		<h2>Créer un nouvel article</h2>
		<form @submit.prevent="createArticle()">
			<div class="inputs-group">
				<div class="label-input">
					<label for="title">Titre de l'article</label>
					<input id="title" name="title" placeholder="Titre de l'article" type="text" v-model="title">
					<p>Son URL d'accès ressemblera à: {{ generateTitleID(title) }}</p>
				</div>
				<div class="label-input">
					<label for="tags">Tags</label>
					<input id="tags" name="tags" placeholder="Tags" type="text" v-model="rawTags"
						@input="updateUniqueTags">
					<p>Spéparez les mots-clés par des virgules ou espaces.</p>
					<ul class="tags" v-if="tags.length !== 0">
						<li v-for="tag in tags">{{ tag }}</li>
					</ul>
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
	margin: 16px 0;
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