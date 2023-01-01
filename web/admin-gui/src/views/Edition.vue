<script setup lang="ts">
import Editor from '@tinymce/tinymce-vue';
import { ref, type Ref } from 'vue';
import { useRoute } from 'vue-router';
import tinymceScriptSrc from '@/assets/tinymce/tinymce.min.js?url'

const articleID = useRoute().params.articleID
const defaultData = `<h1>Bienvenue</h1><p>Vous &ecirc;tes en mode <em>&eacute;dition</em> d'article.</p><p>Celui-ci semble encore neuf ! Supprimez ces lignes et laissez libre cours &agrave; votre imagination :)</p><p>Pour plus d'infos, rendez-vous sur la <a title="Attention, rickroll incoming" href="https:/www.youtube.com/watch?v=dQw4w9WgXcQ" target="_blank" rel="noopener">page d'aide</a>.</p>`
const editorData: Ref<string> = ref(defaultData)

console.log(articleID)

function abort() {
}

function saveContent() {
    console.log(editorData.value)
}
</script>

<template>
    <div class="container">
        <div id="editor">
            <Editor :tinymce-script-src="tinymceScriptSrc"
                :init="{ promotion: false, language: 'fr_FR', resize: false, height: '100%' }"
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