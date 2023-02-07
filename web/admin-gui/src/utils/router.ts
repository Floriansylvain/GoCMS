import { useRouter } from 'vue-router'

const router = useRouter()

export function reloadPage() {
    router.go(0)
}
