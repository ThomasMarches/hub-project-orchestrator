import axios from 'axios';

const apiClient = axios.create({
    baseURL: process.env.API_URL
});

const fetchApiData = async (endpoint) => {
    const response = await fetch(`${apiClient.baseURL}/${endpoint}`);

    // Check if the request was successful
    if (!response.ok) {
        throw new Error(`API request failed: ${response.status}`);
    }

    // Parse the response as JSON
    const data = await response.json();
    return data;
};
