export interface cookie {
	key: string,
	value: string,
	expire: string
}

export function setCookie(cookieData: cookie): void {
	document.cookie = `${cookieData.key}=${cookieData.value}; SameSite=Strict; Expires=${cookieData.expire};Secure`
}

export function getCookie(cookieName: string): string {
	return document.cookie.match('(^|;)\\s*' + cookieName + '\\s*=\\s*([^;]+)')?.pop() || ''
}