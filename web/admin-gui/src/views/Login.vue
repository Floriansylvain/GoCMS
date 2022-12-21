<script setup lang="ts">
import { useAuthStore, type jwtFormat } from '@/stores/AuthStore';
import { ref, type Ref } from 'vue';
import { useRouter } from 'vue-router';

const isTokenOK: Ref<boolean|undefined> = ref(undefined)
const router = useRouter()

function jwtHandler(apiResponse: jwtFormat): void {
	if (apiResponse.code !== 200) {
		console.error(apiResponse)
		return
	}
	useAuthStore().token = apiResponse.token
	useAuthStore().expire = apiResponse.expire
	isTokenOK.value = true
}

//TODO rempalcer tous les :8080 par la var d'env du port de l'API (sinon c un peu dommage)
function login(email: string, password: string): void {
	fetch(`http://localhost:${__APP_ENV__.API_PORT}/login/`, {
		method: "POST",
		body: JSON.stringify({
			email: email,
			password: password
		})
	})
		.then(response => response.json())
		.then(result => {
			jwtHandler(result)
			if (isTokenOK.value === true)  {
				router.push('/home')
			}
		})
}

const email = ref('')
const password = ref('')

</script>

<template>
	<div class="login-page">
		<div class="login-form">
			<h2>Connexion à GohCMS</h2>
			
			<form @submit.prevent="login(email, password)">
				<div class="inputs-group">
					<div class="label-input">
						<label for="email">Adresse mail</label>
						<input id="email" name="email" placeholder="E-mail" type="text" v-model="email">
					</div>
					<div class="label-input">
						<label for="password">Mot de passe</label>
						<input id="password" name="password" placeholder="Mot de passe" type="password" v-model="password">
					</div>
				</div>
				<button class="button-primary" type="submit">Se connecter</button>
			</form>
		</div>
	</div>
</template>

<style scoped>
h2 { 
	margin: 0 0 16px 0;
}
.login-page {
	display: flex;
	justify-content: center;
	align-items: center;

	height: 100vh;
}

.login-form {
	width: 100%;
	max-width: 350px;
	
	padding: 32px;
}

.inputs-group,
.login-form form {
	display: flex;
	flex-direction: column;
	gap: 8px;
}

.login-form form {
	gap: 32px;
}
</style>