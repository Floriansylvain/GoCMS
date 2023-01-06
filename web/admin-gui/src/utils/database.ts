import { useAuthStore } from "@/stores/AuthStore"

export interface Article {
	idName: string,
	date: number,
	content: {
		html: string
	},
	pageId: string,
	online: boolean
}

const baseURL = `http://${__APP_ENV__.APP_HOST_ADDRESS}:${__APP_ENV__.APP_API_PORT}`

export async function getArticle(id: string) : Promise<Article> {
	return await fetch(`${baseURL}/articles/${id}`, {
		method: 'GET',
		headers: { "Authorization": `Bearer ${useAuthStore().token}` }
	})
		.then(result => result.json())
		.catch(error => {
			console.error(error)
		})
}