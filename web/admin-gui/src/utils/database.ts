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

export interface GetArticle {
	content: Article[],
	total: number
	pagination: {
		skip: number,
		take: number
		links: {
			next: string,
			previous: string
		}
	}
}

export function fetchArticle(id: string): Promise<GetArticle> {
	return new Promise((resolve, reject) => {
		fetch(`${baseApiUrl}/articles/${id}`, {
			credentials: 'include',
			method: 'GET',
		})
			.then(response => response.json())
			.then(article => resolve(article))
			.catch(error => reject(error))
	})
}

export function deleteArticle(id: string): Promise<any> {
	return new Promise((resolve, reject) => {
		fetch(`${baseApiUrl}/articles/${id}`, {
			credentials: 'include',
			method: 'DELETE'
		})
			.then(result => result.json())
			.then(article => resolve(article))
			.catch(error => reject(error))
	})
}

export function sendArticleWithMethod(article: Article, method: 'POST' | 'PUT'): Promise<any> {
	return new Promise((resolve, reject) => {
		fetch(`${baseApiUrl}/articles/${article.titleID}`, {
			credentials: 'include',
			method,
			body: JSON.stringify(article)
		})
			.then(result => result.json())
			.then(article => resolve(article))
			.catch(error => reject(error))
	})
}
