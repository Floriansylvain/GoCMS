import { getCookie } from "@/utils/cookies";
import { defineStore } from "pinia";
import { ref } from "vue";
import { useErrorsStore } from "./ErrorsStore";

export interface jwtFormat {
	code: number,
	expire: string,
	message: string
}

export const useAuthStore = defineStore("AuthStore", () => {
	const expire = ref('')
	
	function clearAll(): void {
		expire.value = ''
	}
 
	function isExpired(): boolean {
		if (expire.value === '') return true;
		const tokenDate = new Date(expire.value)
		const currentDate = new Date()

		if (currentDate.getTime() > tokenDate.getTime()) {
			useErrorsStore().sessionExpired = true
			return true
		}
		return false
	}
	
	function initStore(): void {
		const JWTexpire = getCookie('jwt_expire')
		if (JWTexpire !== "") {
			expire.value = JWTexpire
		}
	}
	
	initStore()
	
	return { expire, clearAll, isExpired}
})