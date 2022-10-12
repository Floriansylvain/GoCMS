export interface Pong {
    message: string;
}

export function pingApi(): Promise<Pong> {
    return fetch('http://localhost:8080/ping')
        .then(response => response.json())
}