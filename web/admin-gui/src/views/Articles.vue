<script setup lang="ts">
import { onMounted, ref, type Ref } from 'vue'
import { RouterLink } from 'vue-router'
import { TabulatorFull as Tabulator } from 'tabulator-tables'
import { baseApiUrl } from '@/utils/api'

const table = ref<HTMLInputElement | string>('')
const tabulator: Ref<Tabulator | undefined> = ref(undefined)

onMounted(async () => {
	tabulator.value = new Tabulator(table.value, {
		layout: 'fitColumns',
		locale: 'fr-FR',
		reactiveData: true,
		pagination: true,
		paginationSizeSelector: true,
		paginationMode: 'remote',
		paginationSize: 10,
		ajaxURL: `${baseApiUrl}/articles`,
		ajaxURLGenerator: function (url, config, params) {
			config.credentials = 'include'
			return url + `?skip=${params.size * (params.page - 1)}&take=${params.size}`
		},
		dataReceiveParams: {
			"data": "content"
		},
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
		],
		langs: {
			"fr-FR": {
				"pagination": {
					"first": "Premier",
					"first_title": "Première Page",
					"last": "Dernier",
					"last_title": "Dernière Page",
					"prev": "Précédent",
					"prev_title": "Page Précédente",
					"next": "Suivant",
					"next_title": "Page Suivante",
					"all": "Toute",
					"page_size": "Nombre d'éléments"
				},
			},
		}
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

	height: 100%;

	padding: 32px;
}

main>a {
	width: fit-content;
}

#table {
	width: 100%;
	height: 100%;
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