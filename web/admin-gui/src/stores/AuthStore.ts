import { getCookie } from "@/utils/cookies";
import { defineStore } from "pinia";
import { ref } from "vue";
import { useErrorsStore } from "./ErrorsStore";

export interface jwtFormat {
	code: number,
	expire: string,
	token: string
}

export const useAuthStore = defineStore("AuthStore", () => {
	const expire = ref('')
	const token = ref('')
	
	function clearAll(): void {
		expire.value = ''
		token.value = ''
	}
 
	function isSet(): boolean {
		return token.value !== undefined && token.value !== ''
	}
	
	function isExpired(): boolean {
		const tokenDate = new Date(expire.value)
		const currentDate = new Date()

		if (currentDate.getTime() > tokenDate.getTime()) {
			useErrorsStore().sessionExpired = true
			return true
		}
		return false
	}
	
	function isValid(): boolean {
		return isSet() && !isExpired()
	}
	
	function initStore(): void {
		const JWTtoken = getCookie('JWTtoken')
		const JWTexpire = getCookie('JWTexpire')
		if (JWTtoken !== "" && JWTexpire !== "") {
			token.value = JWTtoken
			expire.value = JWTexpire
		}
	}
	
	initStore()
	
	return { expire, token, isValid, clearAll }
})