export interface Pong {
    message: string;
}

export async function pingApi(): Promise<String> {
    let result: String = ''
    await fetch('http://localhost:8080/ping')
        .then(response => response.json())
        .then(jsonResponse => result = jsonResponse.message)
    return result
}