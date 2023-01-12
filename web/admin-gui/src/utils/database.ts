import { useAuthStore } from "@/stores/AuthStore"
import { baseApiUrl } from "@/utils/api"

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

export async function getArticles(id: string) : Promise<Array<Article>> {
	return await fetch(`${baseApiUrl}/articles/${id}`, {
		credentials: 'include',
		method: 'GET',
	})
		.then(result => result.json())
		.catch(error => {
			console.error(error)
		})
}

export async function postArticle(article: Article) : Promise<object> {
	return await fetch(`${baseApiUrl}/articles/${article.titleID}`, {
		credentials: 'include',
		method: 'POST',
 		body: JSON.stringify(article)
	})
		.then(result => result.json())
		.catch(error => {
			console.error(error)
		})
}