<script setup lang="ts">
import i18n from "@/assets/editor_translation.json";
import defaultData from '@/assets/editor_default_content.json'

import EditorJS, { type EditorConfig, type OutputData } from "@editorjs/editorjs";
import Header from "@editorjs/header";
import List from "@editorjs/list";
import NestedList from "@editorjs/nested-list";
import CheckList from "@editorjs/checklist";
import Quote from "@editorjs/quote";
import Warning from "@editorjs/warning";
import Marker from "@editorjs/marker";
import Code from "@editorjs/code";
import Table from "@editorjs/table";
import Underline from "@editorjs/underline";
import Delimiter from "@editorjs/delimiter";


const props = defineProps<{
	showDefaultData: boolean
}>()

function getEditorConfig(): EditorConfig {
	let editorConfig: EditorConfig = {
		tools: {
			Header,
			Quote,
			NestedList,
			Code,
			Table,
			Warning,
			Delimiter,
			Marker,
			Underline
		},
		i18n
	}
	if (!props.showDefaultData) return editorConfig
	editorConfig = {
		...editorConfig,
		data: defaultData
	}
	return editorConfig
}

const editor = new EditorJS(getEditorConfig())

async function getEditorData(): Promise<OutputData | undefined> {
	let data = undefined
	await editor.save()
		.then(outputData => {
			data = outputData
		})
		.catch((error) => {
			console.error(error)
		})
	return data
}
</script>

<template>
	<div class="container">
		<div id="editorjs"></div>
		<button @click="getEditorData()" class="button-primary">Enregistrer</button>
	</div>
</template>

<style scoped>
.container {
	display: flex;
	flex-direction: column;
	justify-content: center;
	width: 100%;
}

#editorjs {
	font-family: 'Hind', sans-serif;
}

button {
	margin: auto;
}
</style>