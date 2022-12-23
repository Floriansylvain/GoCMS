import { defineStore } from "pinia";
import { ref, type Ref } from "vue";

interface storeErrors {
    sessionExpired: boolean
}

export const useErrorsStore = defineStore("ErrorsStore", () => {
    const errors: Ref<storeErrors> = ref({
        sessionExpired: false
    })

    return errors
})