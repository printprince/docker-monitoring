import {Container} from "react-dom/client";

const API_URL = 'http://localhost:8082/api';

export const api = {
    async getContainers(): Promise<Container[]> {
        const response = await fetch(`${API_URL}/containers`);
        if (!response.ok) {
            throw new Error('Failed to fetch containers');
        }
        return response.json();
    }
};

export default api;