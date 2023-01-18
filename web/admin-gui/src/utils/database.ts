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

export async function getArticles(id: string): Promise<Array<Article>> {
	return await fetch(`${baseApiUrl}/articles/${id}`, {
		credentials: 'include',
		method: 'GET',
	})
		.then(result => result.json())
}

async function sendArticle(article: Article, method: 'POST' | 'PUT') {
	return await fetch(`${baseApiUrl}/articles/${article.titleID}`, {
		credentials: 'include',
		method,
		body: JSON.stringify(article)
	})
		.then(result => result.json())
}

export async function postArticle(article: Article): Promise<object> {
	return await sendArticle(article, 'POST')
}

export async function editArticle(article: Article): Promise<object> {
	return await sendArticle(article, 'PUT')
}