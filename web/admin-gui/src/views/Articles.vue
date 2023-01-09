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
				<th>Actions</th>
			</thead>
			<tbody>
				<tr v-for="article in articles">
					<td>{{ article.title }}</td>
					<td>{{ new Date(article.date).toLocaleDateString('fr-FR') }}</td>
					<td>
						<ul class="tags">
							<li v-for="tag in article.tags">{{ tag }}</li>
						</ul>
					</td>
					<td class="status">{{`${article.online ? '🟢' : '🔴'}`}}</td>
					<td>
						<div class="action-buttons">
							<RouterLink class="button-secondary" :to="`articles/edit/${article.titleID}`">✏️
							</RouterLink>
							<button class="button-secondary">❌</button>
						</div>
					</td>
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

main table {
	width: 100%;

	border-collapse: collapse;
}

main table td {
	padding: 8px;
}

.action-buttons {
	margin: auto;
	width: fit-content;
}

.action-buttons>* {
	margin: 0 6px;
}

.status {
	text-align: center;
}
</style>