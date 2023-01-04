<script setup lang="ts">
import { RouterLink } from 'vue-router'
</script>

<template>
	<main>
		<RouterLink class="button-primary" to="/articles/new">Créer un article</RouterLink>
	</main>
</template>

<style scoped>
main {
	padding: 32px;
}

main>a {
	width: fit-content;
}
</style>