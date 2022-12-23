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

export function deleteCookie(cookieName: string): void {
	setCookie({
		key: cookieName,
		value: '',
		expire: new Date(0).toString()
	})
}