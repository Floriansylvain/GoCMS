<script setup lang="ts">
import { RouterLink } from 'vue-router'

const defaultData = `<h1>Bienvenue</h1><p>Vous &ecirc;tes en mode <em>&eacute;dition</em> d'article.</p><p>Celui-ci semble encore neuf ! Supprimez ces lignes et laissez libre cours &agrave; votre imagination :)</p><p>Pour plus d'infos, rendez-vous sur la <a title="Attention, rickroll incoming" href="https:/www.youtube.com/watch?v=dQw4w9WgXcQ" target="_blank" rel="noopener">page d'aide</a>.</p>`
</script>

<template>
	<main>
		<h2>Créer un nouvel article</h2>
		<form @submit.prevent="">
			<div class="inputs-group">
				<div class="label-input">
					<label for="title">Titre de l'article</label>
					<input id="title" name="title" placeholder="" type="text">
					<p>⚠️ Ce titre servira d'identifiant, il doit être unique !</p>
				</div>
				<div class="label-input">
					<label for="page">Page</label>
					<input id="page" name="page" placeholder="" type="text">
				</div>
			</div>
			<div class="buttons-group">
				<RouterLink class="button-secondary" type="submit" to="/articles">Annuler</RouterLink>
				<button class="button-primary" type="submit">Créer</button>
			</div>
		</form>
	</main>
</template>

<style scoped>
main {
	padding: 0 64px;
	max-width: 100%;
	width: 560px;
}

.buttons-group {
	display: flex;
	gap: 8px;
	width: 100%;
	margin: 32px 0;
}

.buttons-group a,
.buttons-group button {
	width: 100%;
}

.inputs-group {
	display: flex;
	flex-direction: column;
	gap: 8px;
}
</style>