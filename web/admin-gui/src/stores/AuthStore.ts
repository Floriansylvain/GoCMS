import { defineStore } from "pinia";
import { ref } from "vue";

export interface jwtFormat {
	code: number,
	expire: string,
	token: string
}

export const useAuthStore = defineStore("AuthStore", () => {
    const expire = ref("")
    const token = ref("")

    return { expire, token }
})