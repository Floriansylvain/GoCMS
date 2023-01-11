import { useAuthStore } from "@/stores/AuthStore"

export interface Article {
	titleID: string,
	title: string,
	date: number,
	content: {
		html: string
	},
	tags: Array<string>,
	online: boolean
}

const baseURL = `http://${__APP_ENV__.APP_HOST_ADDRESS}:${__APP_ENV__.APP_API_PORT}`

export async function getArticles(id: string) : Promise<Array<Article>> {
	return await fetch(`${baseURL}/articles/${id}`, {
		credentials: 'include',
		method: 'GET',
	})
		.then(result => result.json())
		.catch(error => {
			console.error(error)
		})
}

export async function postArticle(article: Article) : Promise<object> {
	return await fetch(`${baseURL}/articles/${article.titleID}`, {
		credentials: 'include',
		method: 'POST',
 		body: JSON.stringify(article)
	})
		.then(result => result.json())
		.catch(error => {
			console.error(error)
		})
}