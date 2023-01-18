<script setup lang="ts">
import { RouterView, useRoute } from 'vue-router'
import Navbar from './components/Navbar.vue';
</script>

<template>
	<div class="app-container">
		<Navbar v-if="useRoute().name !== 'login'"></Navbar>
		<RouterView />
	</div>
</template>

<style scoped>
.app-container {
	display: flex;
	flex-direction: column;
	height: 100vh;
}
</style>