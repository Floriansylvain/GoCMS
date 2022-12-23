<script setup lang="ts">
import { useAuthStore, type jwtFormat } from '@/stores/AuthStore';
import { useErrorsStore } from '@/stores/ErrorsStore';
import { setCookie } from '@/utils/cookies';
import { ref, type Ref } from 'vue';
import { useRouter } from 'vue-router';

const isTokenOK: Ref<boolean|undefined> = ref(undefined)
const router = useRouter()

const authStore = useAuthStore()

function updateJWTcookies(JWTdata: jwtFormat): void {
	const cookieExpire = new Date(JWTdata.expire)
	cookieExpire.setDate(cookieExpire.getDate() + 1)

	setCookie({
		key: 'JWTtoken',
		value: JWTdata.token,
		expire: cookieExpire.toString()
	})
	setCookie({
		key: 'JWTexpire',
		value: JWTdata.expire,
		expire: cookieExpire.toString()
	})
}

function disableErrors(): void {
	useErrorsStore().sessionExpired = false
}

function jwtHandler(apiResponse: jwtFormat): void {
	if (apiResponse.code !== 200) {
		console.error(apiResponse)
		isTokenOK.value = false
		return
	}
	updateJWTcookies(apiResponse)
	disableErrors()
	authStore.token = apiResponse.token
	authStore.expire = apiResponse.expire
	isTokenOK.value = true
}

function login(email: string, password: string): void {
	fetch(`http://${__APP_ENV__.APP_HOST_ADDRESS}:${__APP_ENV__.APP_API_PORT}/login/`, {
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
					<div :class="`label-input${isTokenOK === false ? '-error' : ''}`">
						<label for="email">Adresse mail</label>
						<input id="email" name="email" placeholder="E-mail" type="text" v-model="email">
					</div>
					<div :class="`label-input${isTokenOK === false ? '-error' : ''}`">
						<label for="password">Mot de passe</label>
						<input id="password" name="password" placeholder="Mot de passe" type="password" v-model="password">
						<p v-if="isTokenOK === false">❌ E-mail et/ou mot de passe incorrect(s).</p>
						<p v-else-if="useErrorsStore().sessionExpired">⌛ Votre session a expiré.</p>
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
	width: fit-content;
	max-width: 100%;
	
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