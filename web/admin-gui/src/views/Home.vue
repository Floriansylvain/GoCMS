<script setup lang="ts">
// import EditorJS from '@editorjs/editorjs';
import { pingApi } from '@/utils/ping';
import { onMounted } from 'vue';

onMounted(async function() {
    console.log('ping...')
    const pong = await pingApi()
    console.log(pong)
})

function addArticle() {
    fetch("http://localhost:8080/add-article", {
        method: "POST",
        body: JSON.stringify({
            id_name: `${Math.random() * 100}`,
            content: {},
            date: new Date().getTime()
        })
    })
        .then(response => response.json())
        .then(result => console.log(result))
}

function logArticles() {
    fetch("http://localhost:8080/get-all-articles")
        .then(response => response.json())
        .then(result => console.log(result))
}
</script>

<template>
    <p>Salut à tous</p>
    <button @click="addArticle">add rand article</button>
    <button @click="logArticles">display articles in console</button>
</template>

<style scoped>

</style>