<script setup lang="ts">
// import EditorJS from '@editorjs/editorjs';
import { pingApi } from '@/utils/ping';
import { onMounted, ref } from 'vue';

const inputArticleID = ref('')

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

function deleteArticle(articleID: String) {
    fetch("http://localhost:8080/delete-article", {
        method: "DELETE",
        body: JSON.stringify({
            id_name: articleID
        })
    })
        .then(response => response.json())
        .then(result => console.log(result))
}
</script>

<template>
    <p>Salut à tous</p>
    <button @click="addArticle">add rand article</button>
    <button @click="logArticles">display articles in console</button> <br>
    <label for="articleIDinput"></label>
    <input id="articleIDinput" name="articleIDinput" type="text" v-model="inputArticleID">
    <button @click="deleteArticle(inputArticleID)">delete this article</button>
</template>

<style scoped>

</style>