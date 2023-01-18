<script setup lang="ts">
import { getArticles, type Article } from '@/utils/database';
import { onMounted, ref, type Ref } from 'vue';
import { RouterLink } from 'vue-router'
import { TabulatorFull as Tabulator } from 'tabulator-tables'

const articles: Ref<Array<Article>> = ref([])

const table = ref<HTMLInputElement | string>('')
const tabulator: Ref<Tabulator | undefined> = ref(undefined)

onMounted(async () => {
	articles.value = await (await getArticles('')).content
	tabulator.value = new Tabulator(table.value, {
		data: articles.value,
		reactiveData: true,
		layout: 'fitColumns',
		columns: [
			{
				title: 'Titre',
				field: 'title',
			},
			{
				title: 'Date création',
				field: 'date',
				formatter: function (cell) {
					return new Date(cell.getValue()).toLocaleDateString('fr-FR')
				}
			},
			{
				title: 'Tags',
				field: 'tags',
				formatter: function (cell) {
					return cell.getValue().join(' ')
				},
				sorter: 'alphanum'
			},
			{
				title: 'Statut',
				field: 'online',
				formatter: function (cell) {
					return cell.getValue() as boolean ? '🟢' : '🔴'
				},
				headerSort: false
			},
			{
				title: 'Actions',
				field: 'titleID',
				formatter: function (cell) {
					const container = document.createElement('div')

					const editButton = document.createElement('a')
					const deleteButton = document.createElement('a')
					editButton.classList.add('button-secondary')
					deleteButton.classList.add('button-secondary')
					editButton.textContent = '✏️'
					deleteButton.textContent = '🗑️'

					editButton.href = `/articles/edit/${cell.getValue()}`

					container.append(editButton, deleteButton)
					return container
				},
				headerSort: false
			}
		]
	})
})
</script>

<template>
	<main>
		<RouterLink class="button-primary" to="/articles/new">Créer un article</RouterLink>
		<div id="table" ref="table"></div>
	</main>
</template>

<style scoped>
main {
	display: flex;
	flex-direction: column;
	gap: 16px;

	padding: 32px;
}

main>a {
	width: fit-content;
}

#table {
	width: 100%;

	margin: auto;
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