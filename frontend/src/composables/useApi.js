import apiConfig from '../api/api.js';

export async function fetchRecommendation(inputData) {
    try {
        const response = await fetch(apiConfig.getRecommendation, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(inputData)
        });
        return response.json();
    } catch (error) {
        console.error('Error fetching recommendation:', error);
        throw error;
    }
}