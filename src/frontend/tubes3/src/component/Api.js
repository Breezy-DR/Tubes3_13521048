const API_URL = "http://localhost:3000";

export async function getUsers() {
    const response = await fetch(`${API_URL}/users`);
    const data = await response.json();
    return data;
}