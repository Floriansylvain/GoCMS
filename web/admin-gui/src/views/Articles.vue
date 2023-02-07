<script setup lang="ts">
import tabulator_langs from '@/assets/tabulator_langs.json'
import { baseApiUrl } from '@/utils/api'
import { deleteArticle } from '@/utils/database'
import { reloadPage } from '@/utils/router'
import { TabulatorFull as Tabulator, type CellComponent } from 'tabulator-tables'
import { onMounted, ref, type Ref } from 'vue'
import { RouterLink } from 'vue-router'

interface BasicTableButton {
	classes: string[],
	text: string
}

const table: Ref<HTMLInputElement | string> = ref('')
const tabulator: Ref<Tabulator | undefined> = ref(undefined)

const basicEditButton = getBasicTableButton({ classes: ['button-secondary', 'table-action-button'], text: '✏️' })
const basicDeleteButton = getBasicTableButton({ classes: ['button-secondary', 'table-action-button'], text: '❌' })

function getBasicTableButton(basicButton: BasicTableButton): HTMLAnchorElement {
	const button = document.createElement('a')
	button.classList.add(...basicButton.classes)
	button.textContent = basicButton.text
	return button
}

function updateTablePage(): void {
	const currentPage = tabulator.value?.getPage()

	if (isNaN(currentPage as number)) {
		reloadPage()
	} else {
		tabulator.value?.setPage(currentPage as number)
	}
}

function getRowEditButton(cell: CellComponent): Node {
	const editButton = basicEditButton.cloneNode(true) as HTMLAnchorElement
	editButton.href = `/articles/edit/${cell.getValue()}`
	return editButton
}

function getRowDeleteButton(cell: CellComponent): Node {
	const deleteButton = basicDeleteButton.cloneNode(true) as HTMLAnchorElement

	deleteButton.onclick = async () => {
		await deleteArticle(cell.getValue())
		updateTablePage()
	}

	return deleteButton
}

function getButtonsCell(cell: CellComponent): HTMLDivElement {
	const container = document.createElement('div')
	container.append(getRowEditButton(cell), getRowDeleteButton(cell))
	return container
}

function getTableAjaxUrlPage(url: string, config: any, params: any) {
	config.credentials = 'include'
	return url + `?skip=${params.size * (params.page - 1)}&take=${params.size}`
}

function initTabulatorTable(): void {
	tabulator.value = new Tabulator(table.value, {
		layout: 'fitColumns',
		locale: 'fr-FR',
		reactiveData: true,
		selectable: false,
		pagination: true,
		paginationSizeSelector: true,
		paginationMode: 'remote',
		paginationSize: 10,
		ajaxURL: `${baseApiUrl}/articles`,
		ajaxURLGenerator: getTableAjaxUrlPage,
		dataReceiveParams: { data: "content" },
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
				formatter: getButtonsCell,
				headerSort: false
			}
		],
		langs: tabulator_langs
	})
}

onMounted(async () => {
	initTabulatorTable()
})
</script>

<template>
	<main>
		<RouterLink class="button-primary" to="/articles/new">Créer un article</RouterLink>
		<div id="table" ref="table"></div>
	</main>
</template>

<style>
.table-action-button {
	margin: 0 4px;
}
</style>

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

.status {
	text-align: center;
}
</style>