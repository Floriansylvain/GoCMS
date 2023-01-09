<script setup lang="ts">
import { getArticles, type Article } from '@/utils/database';
import { onMounted, ref, type Ref } from 'vue';
import { RouterLink } from 'vue-router'

const articles: Ref<Array<Article>> = ref([])

onMounted(async () => {
	articles.value = await getArticles('')
})
</script>

<template>
	<main>
		<RouterLink class="button-primary" to="/articles/new">Créer un article</RouterLink>
		<table>
			<thead>
				<th>Titre</th>
				<th>Date création</th>
				<th>Tags</th>
				<th>Statut</th>
			</thead>
			<tbody>
				<tr v-for="article in articles">
					<td>{{ article.title }}</td>
					<td>{{ new Date(article.date).toLocaleDateString('fr-FR') }}</td>
					<td>{{ article.tags }}</td>
					<td>{{`${article.online ? 'En ligne' : 'Hors ligne'}`}}</td>
					<RouterLink :to="`articles/edit/${article.titleID}`">
						Éditer
					</RouterLink>
					<button>Supprimer</button>
				</tr>
			</tbody>
		</table>
	</main>
</template>

<style scoped>
main {
	padding: 32px;
}

main>a {
	width: fit-content;
}
</style>