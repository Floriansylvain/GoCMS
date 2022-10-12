<script setup lang="ts">
// import EditorJS from '@editorjs/editorjs';
import { pingApi } from '@/utils/ping';

console.log('ping...')
pingApi().then(response => {
    console.log(`${response.message}!`)
})
</script>

<template>
    <p>Salut à tous</p>
</template>

<style scoped>

</style>