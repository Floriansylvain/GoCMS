<script setup lang="ts">
// import EditorJS from '@editorjs/editorjs';
import { pingApi } from '@/utils/ping';
import { onMounted, ref } from 'vue';

const inputArticleID = ref('')
const username = ref('')
const password = ref('')

onMounted(async function() {
    console.log('ping...')
    const pong = await pingApi()
    console.log(pong)
})

function addArticle() {
    fetch("http://localhost:8080/add-article", {
        method: "POST",
        headers: {"Authorization" : `Basic ${btoa(`${username.value}:${password.value}`)}`},
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
    fetch("http://localhost:8080/get-all-articles", {headers: {"Authorization" : `Basic ${btoa(`${username.value}:${password.value}`)}`}})
        .then(response => response.json())
        .then(result => console.log(result))
}

function deleteArticle(articleID: String) {
    fetch("http://localhost:8080/delete-article", {
        method: "DELETE",
        headers: {"Authorization" : `Basic ${btoa(`${username.value}:${password.value}`)}`},
        body: JSON.stringify({
            id_name: articleID
        })
    })
        .then(response => response.json())
        .then(result => console.log(result))
}

function auth(type: string, username: string, password: string) {
    fetch(`http://localhost:8080/${type}`, {
        method: "POST",
        body: JSON.stringify({
            email: username,
            password: password
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
    <input id="articleIDinput" name="articleIDinput" type="text" placeholder="article ID to delete" v-model="inputArticleID">
    <button @click="deleteArticle(inputArticleID)">delete this article</button> <br>

    <input type="text" name="username" id="username" v-model="username" placeholder="username">
    <input type="text" name="password" id="password" v-model="password" placeholder="password">
    <button @click="auth('login', username, password)">login</button>
    <button @click="auth('logout', username, password)">logout</button>
</template>

<style scoped>

</style>