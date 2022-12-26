<script setup lang="ts">
import { RouterView, useRoute } from 'vue-router'
import Navbar from './components/Navbar.vue';
</script>

<template>
	<Navbar v-if="useRoute().name !== 'login'"></Navbar>
	<RouterView />
</template>
